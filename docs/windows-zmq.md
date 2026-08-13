# Windows Native Build

## One-time fix

If `go build` on Windows keeps failing because of `zmq`, first repair the local Go/toolchain environment:

```powershell
Set-Location D:\code\fcas-go-v2\server
.\setup_zmq_env.ps1 -PersistUserPath -PersistToolchainEnv -RepairGoEnv
```

This does two things:

- adds `C:\msys64\usr\bin` and the detected MSYS2 toolchain `bin` directory to the user `PATH`
- persists `CGO_ENABLED=1`, `CC`, `CXX`, `PKG_CONFIG`, and `PKG_CONFIG_PATH` for the detected MSYS2 toolchain
- clears user-level Go overrides such as `GOOS`, `GOARCH`, `CGO_ENABLED`, `CC`, `CXX`, `PKG_CONFIG`, `CGO_CFLAGS`, and `CGO_LDFLAGS`

After that, open a new terminal window.

## Run on Windows

From the repository root:

```powershell
.\run_server_windows.cmd
```

## Build on Windows

From the repository root:

```powershell
.\build_server_windows.cmd
```

The build output goes to `build\server\` and includes:

- `fcas_server.exe`
- `config*.yaml`
- `libzmq.dll`

That output directory can be run directly on Windows without adding `libzmq.dll` to the system `PATH`.
