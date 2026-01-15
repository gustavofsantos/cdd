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

# Check for go
if ! command -v go &> /dev/null; then
    echo "Error: 'go' is not installed. Please install Go to build CDD from source."
    echo "Visit https://golang.org/doc/install for instructions."
    exit 1
fi

# Determine if we are in the repo or need to clone
if [ ! -f "cmd/cdd/main.go" ]; then
    if ! command -v git &> /dev/null; then
        echo "Error: 'git' is not installed. Please install git to download CDD source."
        exit 1
    fi
    echo "Not in the CDD repository. Cloning to temporary directory..."
    TEMP_DIR=$(mktemp -d)
    git clone --depth 1 "$REPO_URL" "$TEMP_DIR"
    cd "$TEMP_DIR"
    trap 'rm -rf "$TEMP_DIR"' EXIT
fi

echo "Building CDD..."
go build -o cdd cmd/cdd/main.go

echo "Installing CDD to $INSTALL_DIR..."
# Check if sudo is needed
if [ -w "$INSTALL_DIR" ]; then
    mv cdd "$INSTALL_DIR/cdd"
else
    echo "Need sudo permissions to install to $INSTALL_DIR"
    sudo mv cdd "$INSTALL_DIR/cdd"
fi

chmod +x "$INSTALL_DIR/cdd"

echo "CDD installed successfully to $INSTALL_DIR/cdd"

# Install Amp Toolbox wrappers (Optional)
if [ "$INSTALL_TOOLBOX" = true ]; then
    echo ""
    echo "Building Amp toolbox wrappers..."

    # Create toolbox directory
    if [ ! -d "$TOOLBOX_DIR" ]; then
        if [ -w "$INSTALL_DIR" ]; then
            mkdir -p "$TOOLBOX_DIR"
        else
            sudo mkdir -p "$TOOLBOX_DIR"
            sudo chown "$(whoami)" "$TOOLBOX_DIR"
        fi
    fi

    # Build and install each toolbox wrapper
    TOOLBOX_TOOLS=("init" "start" "recite" "log" "archive" "view" "agents" "delete" "version" "pack")

    for tool in "${TOOLBOX_TOOLS[@]}"; do
        echo "  Building cdd-$tool..."
        go build -o "cdd-$tool" "./cmd/toolbox/$tool/main.go"
        
        if [ -w "$TOOLBOX_DIR" ]; then
            mv "cdd-$tool" "$TOOLBOX_DIR/cdd-$tool"
        else
            sudo mv "cdd-$tool" "$TOOLBOX_DIR/cdd-$tool"
        fi
        
        chmod +x "$TOOLBOX_DIR/cdd-$tool"
    done

    echo "Amp toolbox installed successfully to $TOOLBOX_DIR"
else
    echo ""
    echo "Skipping Amp toolbox installation. Use --amp-toolbox to include it."
fi

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
