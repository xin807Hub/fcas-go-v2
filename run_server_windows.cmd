@echo off
setlocal

set "ROOT_DIR=%~dp0"
call "%ROOT_DIR%server\run_windows.cmd" %*
set "EXIT_CODE=%ERRORLEVEL%"

endlocal & exit /b %EXIT_CODE%
