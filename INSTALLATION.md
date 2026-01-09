# Installation

## Installation via Binary

You can download the pre-compiled binary for your platform from the [Releases](https://github.com/gustavofsantos/cdd/releases) page.

**Linux:**

```bash
# Example for Linux x86_64
curl -L -O https://github.com/gustavofsantos/cdd/releases/latest/download/cdd_Linux_x86_64.tar.gz
tar -xzf cdd_Linux_x86_64.tar.gz
chmod +x cdd
sudo mv cdd /usr/local/bin/
rm cdd_Linux_x86_64.tar.gz
```

**macOS:**

```bash
# Example for macOS Apple Silicon (arm64)
curl -L -O https://github.com/gustavofsantos/cdd/releases/latest/download/cdd_Darwin_arm64.tar.gz
tar -xzf cdd_Darwin_arm64.tar.gz
chmod +x cdd
sudo mv cdd /usr/local/bin/
rm cdd_Darwin_arm64.tar.gz
```

**Windows (PowerShell):**

```powershell
# Example for Windows x86_64
Invoke-WebRequest -Uri "https://github.com/gustavofsantos/cdd/releases/latest/download/cdd_Windows_x86_64.zip" -OutFile "cdd.zip"
Expand-Archive -Path "cdd.zip" -DestinationPath "."
# Move cdd.exe to a directory in your PATH
Move-Item -Path "cdd.exe" -Destination "C:\Windows\System32\cdd.exe"
Remove-Item "cdd.zip"
```

## Build from Source

```bash
git clone https://github.com/gustavofsantos/cdd.git
cd cdd
go build -o cdd cmd/cdd/main.go
```

You can then move the `cdd` binary to a directory in your system's PATH.
