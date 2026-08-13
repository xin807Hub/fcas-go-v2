param(
    [string]$Config,
    [string]$ServiceConfig,
    [string]$GoRoot,
    [ValidateSet("auto", "ucrt64", "mingw64", "clang64")]
    [string]$Toolchain = "auto",
    [ValidateSet("run", "build", "check")]
    [string]$Action = "run",
    [string]$OutputDir,
    [switch]$PersistUserPath,
    [switch]$PersistToolchainEnv,
    [switch]$RepairGoEnv,
    [switch]$CheckOnly,
    [switch]$PrintEnv,
    [Parameter(ValueFromRemainingArguments = $true)]
    [string[]]$AppArgs
)

# Run the server with the Windows ZMQ toolchain loaded in the current session.
# Examples:
#   .\run_windows.ps1 -CheckOnly
#   .\run_windows.ps1 -Config .\config_dev.yaml
#   .\run_windows.ps1 -Action build
#   .\run_windows.ps1 -Action check -PersistUserPath -PersistToolchainEnv -RepairGoEnv
#   .\run_windows.ps1 -Toolchain ucrt64 -- -v
$scriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
$repoRoot = Split-Path -Parent $scriptDir

if ($CheckOnly) {
    $Action = "check"
}

if ($GoRoot) {
    $env:GOROOT = $GoRoot
    $goBin = Join-Path $env:GOROOT "bin"
    if (Test-Path $goBin) {
        $env:PATH = "$goBin;$env:PATH"
    }
}

if ($Config) {
    $env:GVA_CONFIG = $Config
    $env:CONFIG = $Config
}

if ($ServiceConfig) {
    $env:SERVICE_CONFIG = $ServiceConfig
}

. (Join-Path $scriptDir "setup_zmq_env.ps1") `
    -Toolchain $Toolchain `
    -PersistUserPath:$PersistUserPath `
    -PersistToolchainEnv:$PersistToolchainEnv `
    -RepairGoEnv:$RepairGoEnv `
    -Quiet:(-not $PrintEnv)

if ($Action -eq "check") {
    Write-Host "ZMQ Windows environment is ready."
    return
}

Push-Location $scriptDir
try {
    $goExe = (Get-Command go -ErrorAction Stop).Source

    if ($Action -eq "build") {
        if (-not $OutputDir) {
            $OutputDir = Join-Path $repoRoot "build\server"
        }

        $resolvedOutputDir = [System.IO.Path]::GetFullPath($OutputDir)
        New-Item -ItemType Directory -Force -Path $resolvedOutputDir | Out-Null

        $buildVersionFile = Join-Path $repoRoot "BUILD_VERSION"
        $buildVersion = ""
        if (Test-Path $buildVersionFile) {
            $buildVersion = (Get-Content $buildVersionFile -Raw).Trim()
        }
        $buildTime = Get-Date -Format "yyyyMMddHHmmss"
        $ldflags = "-X main.buildTime=$buildTime -X main.buildVer=$buildVersion"
        $outputExe = Join-Path $resolvedOutputDir "fcas_server.exe"

        & $goExe build -ldflags $ldflags -o $outputExe .
        if ($LASTEXITCODE -ne 0) {
            exit $LASTEXITCODE
        }

        Get-ChildItem -LiteralPath $scriptDir -Filter "config*.yaml" | ForEach-Object {
            Copy-Item -LiteralPath $_.FullName -Destination $resolvedOutputDir -Force
        }

        if ($env:ZMQ_DLL_PATH -and (Test-Path $env:ZMQ_DLL_PATH)) {
            Copy-Item -LiteralPath $env:ZMQ_DLL_PATH -Destination (Join-Path $resolvedOutputDir "libzmq.dll") -Force
        }

        Write-Host ("Build output      : {0}" -f $outputExe)
        Write-Host ("Runtime config(s) : {0}" -f $resolvedOutputDir)
        if ($env:ZMQ_DLL_PATH -and (Test-Path $env:ZMQ_DLL_PATH)) {
            Write-Host ("Bundled DLL       : {0}" -f (Join-Path $resolvedOutputDir "libzmq.dll"))
        }
        return
    }

    & $goExe run . @AppArgs
} finally {
    Pop-Location
}
