# Server configuration & start instructions

This document explains how to configure and start the ScribbleGuess backend server.

**Config**
- **SERVER_PORT**: the TCP port the server listens on. The code requires this environment variable to be set and parses it as an unsigned 16-bit integer (valid range 0–65535). See [internal/config/config.go](internal/config/config.go#L1-L20).

**Start (Windows — CMD)**
- Use the provided helper: [start_server.bat](start_server.bat#L1-L20) (sets `SERVER_PORT` to `8080` by default and runs `go run ./cmd/main.go`). Double-click the batch file or run it from CMD:

```
start_server.bat
```

- Or set the environment variable manually and run:

```
:: in CMD
set SERVER_PORT=8080
go run ./cmd/main.go

:: in PowerShell
$env:SERVER_PORT = "8080"
go run ./cmd/main.go

:: in Git Bash / WSL
export SERVER_PORT=8080
go run ./cmd/main.go
```

**Build and run (recommended for production/testing)**

```
go build -o bin/scribble ./cmd

:: set env and run (Windows)
set SERVER_PORT=8080
bin\scribble.exe

:: Unix-like
export SERVER_PORT=8080
./bin/scribble
```

**Behavior**
- The server prints the listening port at startup and supports graceful shutdown on `Ctrl+C` (it allows ~5 seconds for in-flight requests to finish). See [cmd/main.go](cmd/main.go#L1-L60).

**Troubleshooting & Notes**
- `SERVER_PORT` must be a valid integer within `0`–`65535`. If missing or invalid, the server will fail to start with an error.
- If you change the port in `start_server.bat`, ensure it remains within the valid range.
- If you intend to restrict the port to smaller ranges (for example `uint8`), update `internal/config/config.go` accordingly and verify all call sites.

If you'd like, I can also add a small `Makefile` or PowerShell script for common tasks (build, run, test). 
