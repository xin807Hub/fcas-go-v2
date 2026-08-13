param(
    [ValidateSet("auto", "ucrt64", "mingw64", "clang64")]
    [string]$Toolchain = "auto",
    [switch]$PersistUserPath,
    [switch]$PersistToolchainEnv,
    [switch]$RepairGoEnv,
    [switch]$Quiet
)

Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"

function Add-PathPrefix {
    param([string]$PathItem)

    if ([string]::IsNullOrWhiteSpace($PathItem)) {
        return
    }
    if (-not (Test-Path $PathItem)) {
        return
    }

    $target = [System.IO.Path]::GetFullPath($PathItem)
    $pathItems = @($env:PATH -split ";" | Where-Object { -not [string]::IsNullOrWhiteSpace($_) })
    if ($pathItems -contains $target) {
        return
    }
    $env:PATH = "$target;$env:PATH"
}

function Add-EnvListValue {
    param(
        [string]$Name,
        [string]$Value
    )

    if ([string]::IsNullOrWhiteSpace($Value)) {
        return
    }

    $currentValue = [Environment]::GetEnvironmentVariable($Name, "Process")
    $items = @()
    if ($currentValue) {
        $items = @($currentValue -split ";" | Where-Object { -not [string]::IsNullOrWhiteSpace($_) })
    }

    if ($items -notcontains $Value) {
        if ($items.Count -eq 0) {
            [Environment]::SetEnvironmentVariable($Name, $Value, "Process")
        } else {
            [Environment]::SetEnvironmentVariable($Name, "$Value;$($items -join ';')", "Process")
        }
    }
}

function Add-UserPathPrefix {
    param([string]$PathItem)

    if ([string]::IsNullOrWhiteSpace($PathItem)) {
        return
    }
    if (-not (Test-Path $PathItem)) {
        return
    }

    $target = [System.IO.Path]::GetFullPath($PathItem)
    $currentValue = [Environment]::GetEnvironmentVariable("Path", "User")
    $items = @()
    if ($currentValue) {
        $items = @($currentValue -split ";" | Where-Object { -not [string]::IsNullOrWhiteSpace($_) })
    }

    if ($items -contains $target) {
        return
    }

    if ($items.Count -eq 0) {
        [Environment]::SetEnvironmentVariable("Path", $target, "User")
    } else {
        [Environment]::SetEnvironmentVariable("Path", "$target;$($items -join ';')", "User")
    }
}

function Add-UserEnvListValue {
    param(
        [string]$Name,
        [string]$Value
    )

    if ([string]::IsNullOrWhiteSpace($Value)) {
        return
    }

    $currentValue = [Environment]::GetEnvironmentVariable($Name, "User")
    $items = @()
    if ($currentValue) {
        $items = @($currentValue -split ";" | Where-Object { -not [string]::IsNullOrWhiteSpace($_) })
    }

    if ($items -contains $Value) {
        return
    }

    if ($items.Count -eq 0) {
        [Environment]::SetEnvironmentVariable($Name, $Value, "User")
    } else {
        [Environment]::SetEnvironmentVariable($Name, "$Value;$($items -join ';')", "User")
    }
}

function Set-UserEnvValue {
    param(
        [string]$Name,
        [string]$Value
    )

    [Environment]::SetEnvironmentVariable($Name, $Value, "User")
}

function Clear-GoEnvOverrides {
    $goExe = (Get-Command go -ErrorAction Stop).Source
    $keys = @(
        "GOOS",
        "GOARCH",
        "CGO_ENABLED",
        "CC",
        "CXX",
        "PKG_CONFIG",
        "CGO_CFLAGS",
        "CGO_LDFLAGS"
    )

    foreach ($key in $keys) {
        & $goExe env -u $key | Out-Null
        if ($LASTEXITCODE -ne 0) {
            throw "Failed to clear go env override: $key"
        }
    }
}

function Resolve-MsysToolchain {
    param([string]$PreferredToolchain)

    $msysRoot = "C:\msys64"
    $candidateNames = if ($PreferredToolchain -eq "auto") {
        @("ucrt64", "mingw64", "clang64")
    } else {
        @($PreferredToolchain)
    }

    $pkgConfigCmd = Get-Command pkg-config -ErrorAction SilentlyContinue
    if ($pkgConfigCmd -and $pkgConfigCmd.Source -like "$msysRoot\*\bin\pkg-config.exe") {
        $currentToolchain = Split-Path -Leaf (Split-Path -Parent (Split-Path -Parent $pkgConfigCmd.Source))
        if ($candidateNames -notcontains $currentToolchain) {
            $candidateNames = @($currentToolchain) + $candidateNames
        }
    }

    foreach ($candidateName in $candidateNames) {
        $toolchainRoot = Join-Path $msysRoot $candidateName
        $binDir = Join-Path $toolchainRoot "bin"
        $libDir = Join-Path $toolchainRoot "lib"
        $pkgConfigDir = Join-Path $libDir "pkgconfig"
        $includeDir = Join-Path $toolchainRoot "include"
        $pkgConfigExe = Join-Path $binDir "pkg-config.exe"
        $libzmqPc = Join-Path $pkgConfigDir "libzmq.pc"
        $libzmqDll = Join-Path $binDir "libzmq.dll"

        if (-not (Test-Path $pkgConfigExe)) {
            continue
        }
        if (-not (Test-Path $libzmqPc)) {
            continue
        }
        if (-not (Test-Path $libzmqDll)) {
            continue
        }

        $ccExe = Join-Path $binDir "gcc.exe"
        $cxxExe = Join-Path $binDir "g++.exe"
        if (-not (Test-Path $ccExe)) {
            $ccExe = Join-Path $binDir "clang.exe"
            $cxxExe = Join-Path $binDir "clang++.exe"
        }
        if (-not (Test-Path $ccExe)) {
            continue
        }
        if (-not (Test-Path $cxxExe)) {
            continue
        }

        return [PSCustomObject]@{
            Name         = $candidateName
            MsysRoot     = $msysRoot
            Root         = $toolchainRoot
            BinDir       = $binDir
            LibDir       = $libDir
            IncludeDir   = $includeDir
            PkgConfigDir = $pkgConfigDir
            PkgConfigExe = $pkgConfigExe
            CcExe        = $ccExe
            CxxExe       = $cxxExe
            LibZmqPc     = $libzmqPc
            LibZmqDll    = $libzmqDll
            UsrBinDir    = Join-Path $msysRoot "usr\bin"
        }
    }

    throw "Unable to locate a usable MSYS2 ZeroMQ toolchain. Expected libzmq under C:\msys64\(ucrt64|mingw64|clang64)."
}

function Test-ZmqPkgConfig {
    param([object]$ToolchainInfo)

    $oldOutputEncoding = [Console]::OutputEncoding
    try {
        [Console]::OutputEncoding = [System.Text.Encoding]::UTF8
        $version = & $ToolchainInfo.PkgConfigExe --modversion libzmq 2>$null
        if ($LASTEXITCODE -ne 0 -or [string]::IsNullOrWhiteSpace($version)) {
            throw "pkg-config can not resolve libzmq. Check PKG_CONFIG_PATH and libzmq.pc."
        }
        return $version.Trim()
    } finally {
        [Console]::OutputEncoding = $oldOutputEncoding
    }
}

$toolchainInfo = Resolve-MsysToolchain -PreferredToolchain $Toolchain

Add-PathPrefix $toolchainInfo.UsrBinDir
Add-PathPrefix $toolchainInfo.BinDir
Add-EnvListValue -Name "PKG_CONFIG_PATH" -Value $toolchainInfo.PkgConfigDir

if ($PersistUserPath) {
    Add-UserPathPrefix $toolchainInfo.UsrBinDir
    Add-UserPathPrefix $toolchainInfo.BinDir
}

if ($RepairGoEnv) {
    Clear-GoEnvOverrides
}

if ($PersistToolchainEnv) {
    Set-UserEnvValue -Name "CGO_ENABLED" -Value "1"
    Set-UserEnvValue -Name "CC" -Value $toolchainInfo.CcExe
    Set-UserEnvValue -Name "CXX" -Value $toolchainInfo.CxxExe
    Set-UserEnvValue -Name "PKG_CONFIG" -Value $toolchainInfo.PkgConfigExe
    Add-UserEnvListValue -Name "PKG_CONFIG_PATH" -Value $toolchainInfo.PkgConfigDir
}

$env:CGO_ENABLED = "1"
$env:GOOS = (& go env GOHOSTOS).Trim()
$env:GOARCH = (& go env GOHOSTARCH).Trim()
$env:CC = $toolchainInfo.CcExe
$env:CXX = $toolchainInfo.CxxExe
$env:PKG_CONFIG = $toolchainInfo.PkgConfigExe
$env:ZMQ_ROOT = $toolchainInfo.Root
$env:ZMQ_BIN_DIR = $toolchainInfo.BinDir
$env:ZMQ_LIB_DIR = $toolchainInfo.LibDir
$env:ZMQ_INCLUDE_DIR = $toolchainInfo.IncludeDir
$env:ZMQ_DLL_PATH = $toolchainInfo.LibZmqDll

$libzmqVersion = Test-ZmqPkgConfig -ToolchainInfo $toolchainInfo

if (-not $Quiet) {
    Write-Host ("ZeroMQ toolchain : {0}" -f $toolchainInfo.Name)
    Write-Host ("libzmq version   : {0}" -f $libzmqVersion)
    Write-Host ("GOOS/GOARCH      : {0}/{1}" -f $env:GOOS, $env:GOARCH)
    Write-Host ("CGO_ENABLED      : {0}" -f $env:CGO_ENABLED)
    Write-Host ("CC               : {0}" -f $env:CC)
    Write-Host ("PKG_CONFIG       : {0}" -f $env:PKG_CONFIG)
    if ($PersistUserPath) {
        Write-Host "User PATH        : updated with MSYS2 runtime"
    }
    if ($PersistToolchainEnv) {
        Write-Host "User toolchain   : persisted for native Windows CGO builds"
    }
    if ($RepairGoEnv) {
        Write-Host "Go env overrides : cleared"
    }
}
