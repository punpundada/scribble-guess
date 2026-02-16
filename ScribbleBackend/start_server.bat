@echo off

set WINDOW_NAME=Go Server
set SERVER_PORT=7412

title %WINDOW_NAME%

set SERVER_PORT=%SERVER_PORT%

echo Starting server on %SERVER_PORT%...

go run ./cmd/main.go

pause
