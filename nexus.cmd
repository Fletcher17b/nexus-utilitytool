@echo off
for /f "delims=" %%i in ('"%~dp0nexus-cli.exe" %*') do (
    cd /d "%%i"
)