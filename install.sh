#!/bin/bash

# CDD Installation Script
set -e

# Default values
INSTALL_DIR="/usr/local/bin"
TOOLBOX_DIR=""
INSTALL_TOOLBOX=false
REPO_URL="https://github.com/gustavofsantos/cdd.git"

# Parse arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        --amp-toolbox)
            INSTALL_TOOLBOX=true
            shift
            ;;
        --install-dir)
            INSTALL_DIR="$2"
            shift 2
            ;;
        --amp-toolbox-dir)
            TOOLBOX_DIR="$2"
            shift 2
            ;;
        -*)
            echo "Unknown option $1"
            exit 1
            ;;
        *)
            # For backward compatibility, the first non-flag argument is the INSTALL_DIR
            INSTALL_DIR="$1"
            shift
            ;;
    esac
done

# Set default TOOLBOX_DIR if not provided
if [ -z "$TOOLBOX_DIR" ]; then
    TOOLBOX_DIR="$INSTALL_DIR/toolbox"
fi

echo "═══════════════════════════════════════════════════════════════"
echo "        Context-Driven Development (CDD) Installer"
echo "═══════════════════════════════════════════════════════════════"

# Determine if we are in the repo or need to download
if [ -f "cmd/cdd/main.go" ]; then
    echo "Running from CDD repository. Building from source..."
    
    # Check for go
    if ! command -v go &> /dev/null; then
        echo "Error: 'go' is not installed. Please install Go to build CDD from source."
        echo "Visit https://golang.org/doc/install for instructions."
        exit 1
    fi

    echo "Building CDD..."
    go build -o cdd cmd/cdd/main.go

    echo "Installing CDD to $INSTALL_DIR..."
    if [ -w "$INSTALL_DIR" ]; then
        mv cdd "$INSTALL_DIR/cdd"
    else
        echo "Need sudo permissions to install to $INSTALL_DIR"
        sudo mv cdd "$INSTALL_DIR/cdd"
    fi
    chmod +x "$INSTALL_DIR/cdd"

    if [ "$INSTALL_TOOLBOX" = true ]; then
        echo ""
        echo "Building Amp toolbox wrappers..."
        if [ ! -d "$TOOLBOX_DIR" ]; then
            if [ -w "$INSTALL_DIR" ]; then mkdir -p "$TOOLBOX_DIR"; else sudo mkdir -p "$TOOLBOX_DIR"; sudo chown "$(whoami)" "$TOOLBOX_DIR"; fi
        fi
        TOOLBOX_TOOLS=("init" "start" "recite" "log" "archive" "view" "agents" "delete" "version" "pack")
        for tool in "${TOOLBOX_TOOLS[@]}"; do
            echo "  Building cdd-$tool..."
            go build -o "cdd-$tool" "./cmd/toolbox/$tool/main.go"
            if [ -w "$TOOLBOX_DIR" ]; then mv "cdd-$tool" "$TOOLBOX_DIR/cdd-$tool"; else sudo mv "cdd-$tool" "$TOOLBOX_DIR/cdd-$tool"; fi
            chmod +x "$TOOLBOX_DIR/cdd-$tool"
        done
    fi
else
    echo "Not in the CDD repository. Downloading pre-built binary..."
    
    OS=$(uname -s)
    ARCH=$(uname -m)
    
    case "$OS" in
        Darwin)  OS_NAME="Darwin" ;;
        Linux)   OS_NAME="Linux" ;;
        *)       echo "Unsupported OS: $OS"; exit 1 ;;
    esac

    case "$ARCH" in
        x86_64) ARCH_NAME="x86_64" ;;
        arm64|aarch64) ARCH_NAME="arm64" ;;
        *)      echo "Unsupported architecture: $ARCH"; exit 1 ;;
    esac

    BINARY_NAME="cdd_${OS_NAME}_${ARCH_NAME}.tar.gz"
    DOWNLOAD_URL="https://github.com/gustavofsantos/cdd/releases/latest/download/${BINARY_NAME}"

    echo "Downloading CDD for ${OS_NAME} ${ARCH_NAME}..."
    TEMP_DIR=$(mktemp -d)
    trap 'rm -rf "$TEMP_DIR"' EXIT

    if ! curl -L -f -o "$TEMP_DIR/$BINARY_NAME" "$DOWNLOAD_URL"; then
        echo "Error: Failed to download binary from $DOWNLOAD_URL"
        echo "Please check if a release exists for your platform."
        exit 1
    fi

    echo "Extracting..."
    tar -xzf "$TEMP_DIR/$BINARY_NAME" -C "$TEMP_DIR"

    echo "Installing CDD to $INSTALL_DIR..."
    if [ -w "$INSTALL_DIR" ]; then
        mv "$TEMP_DIR/cdd" "$INSTALL_DIR/cdd"
    else
        echo "Need sudo permissions to install to $INSTALL_DIR"
        sudo mv "$TEMP_DIR/cdd" "$INSTALL_DIR/cdd"
    fi
    chmod +x "$INSTALL_DIR/cdd"

    if [ "$INSTALL_TOOLBOX" = true ]; then
        echo "Installing Amp toolbox wrappers..."
        if [ ! -d "$TOOLBOX_DIR" ]; then
            if [ -w "$INSTALL_DIR" ]; then mkdir -p "$TOOLBOX_DIR"; else sudo mkdir -p "$TOOLBOX_DIR"; sudo chown "$(whoami)" "$TOOLBOX_DIR"; fi
        fi
        
        # Move all cdd-* files found in the archive
        find "$TEMP_DIR" -maxdepth 1 -name "cdd-*" -type f -exec bash -c '
            FILE=$1
            DEST=$2
            if [ -w $(dirname "$DEST") ]; then
                mv "$FILE" "$DEST/$(basename "$FILE")"
            else
                sudo mv "$FILE" "$DEST/$(basename "$FILE")"
            fi
            chmod +x "$DEST/$(basename "$FILE")"
        ' -- {} "$TOOLBOX_DIR" \;
    fi
fi

echo "CDD installed successfully!"

# Shell configuration check
if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
    echo ""
    echo "WARNING: $INSTALL_DIR is not in your PATH."
    echo "You may need to add it to your shell configuration (e.g., .bashrc, .zshrc):"
    echo "  export PATH=\$PATH:$INSTALL_DIR"
fi

# Amp Toolbox setup instructions
if [ "$INSTALL_TOOLBOX" = true ]; then
    echo ""
    echo "═══════════════════════════════════════════════════════════════"
    echo "Amp Toolbox Setup (Optional)"
    echo "═══════════════════════════════════════════════════════════════"
    echo ""
    echo "To enable CDD tools in Amp, set the AMP_TOOLBOX environment variable:"
    echo ""
    echo "  export AMP_TOOLBOX=$TOOLBOX_DIR"
    echo ""
    echo "Add this line to your shell configuration (.bashrc, .zshrc, etc.):"
    echo "  echo 'export AMP_TOOLBOX=$TOOLBOX_DIR' >> ~/.bashrc"
    echo ""
    echo "Then restart your terminal or source your config:"
    echo "  source ~/.bashrc"
    echo ""
    echo "Amp will automatically discover all CDD tools on next startup."
    echo "See AMP_TOOLBOX.md for more details."
    echo ""
fi

echo "Installation complete! Run 'cdd' to get started!"
