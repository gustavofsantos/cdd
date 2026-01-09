# Installation

## Installation via Binary

You can download the pre-compiled binary for your platform from the [Releases](https://github.com/gustavofsantos/cdd/releases) page.

**Linux:**

```bash
# Example for Linux AMD64
curl -L -o cdd https://github.com/gustavofsantos/cdd/releases/latest/download/cdd-linux-amd64
chmod +x cdd
sudo mv cdd /usr/local/bin/
```

**macOS:**

```bash
# Example for macOS Apple Silicon
curl -L -o cdd https://github.com/gustavofsantos/cdd/releases/latest/download/cdd-darwin-arm64
chmod +x cdd
sudo mv cdd /usr/local/bin/
```

**Windows (PowerShell):**

```powershell
# Example for Windows AMD64
Invoke-WebRequest -Uri "https://github.com/gustavofsantos/cdd/releases/latest/download/cdd-windows-amd64.exe" -OutFile "cdd.exe"
# Move to a directory in your PATH, e.g., C:\Windows\System32 or a custom tools folder
Move-Item -Path "cdd.exe" -Destination "C:\Windows\System32\cdd.exe"
```

## Build from Source

```bash
git clone https://github.com/gustavofsantos/cdd.git
cd cdd
go build -o cdd cmd/cdd/main.go
```

You can then move the `cdd` binary to a directory in your system's PATH.
