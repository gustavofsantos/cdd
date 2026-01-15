# Installation

## Installation via Script (Recommended)

If you have `curl` and `tar` installed, you can use the one-line installation script. This script will automatically download the latest pre-compiled binary for your system (macOS or Linux) and install it to `/usr/local/bin`. If you run the script from within the CDD source repository, it will build from source instead.

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/gustavofsantos/cdd/main/install.sh)"
```

To also install the Amp toolbox wrappers, use the `--amp-toolbox` flag:

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/gustavofsantos/cdd/main/install.sh)" -- --amp-toolbox
```

To install CDD to a different directory (e.g., `~/bin`), use the `--install-dir` flag:

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/gustavofsantos/cdd/main/install.sh)" -- --install-dir ~/.bin
```

To install the Amp toolbox to a specific path, use the `--amp-toolbox-dir` flag:

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/gustavofsantos/cdd/main/install.sh)" -- --amp-toolbox --amp-toolbox-dir ~/.amp/toolbox
```

### Advanced Script Installation

You can also specify a custom installation directory:

```bash
git clone https://github.com/gustavofsantos/cdd.git
cd cdd
./install.sh /custom/bin/path
```

## Installation via Binary (Alternative)

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
