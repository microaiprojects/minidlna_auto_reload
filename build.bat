@echo off

echo Building for Windows AMD64...
set GOOS=windows
set GOARCH=amd64
set CGO_ENABLED=0
go build -ldflags="-s -w" -o mini_dlna_updater_windows_amd64.exe

echo Build complete!
