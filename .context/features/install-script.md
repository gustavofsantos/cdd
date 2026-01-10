# Specification: Installation Script

## User Intent
Create an `install.sh` script to automate the build and installation of the `cdd` tool.

## Relevant Context
*   `cmd/cdd/main.go`: Entry point for the application.
*   `INSTALLATION.md`: Current installation instructions (manual).
*   `scripts/`: Directory for scripts.

## Context Analysis
Currently, users have to manually run `go build` and move the binary. A script will simplify this. The script should:
1.  Accept an optional installation directory (defaulting to `/usr/local/bin` or similar if appropriate, but the user specifically said "takes a path").
2.  Build the binary using `go build`.
3.  Move the binary to the target directory.
4.  Optionally/Ideally check if the directory is in the `PATH` and advise or assist in adding it.

## Test Reference
*   **Manual Verification:** Run `./install.sh $(pwd)/test_install` and verified `test_install/cdd` exists and is executable.
