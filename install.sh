#!/bin/bash

# CDD Installation Script
set -e

INSTALL_DIR=${1:-"/usr/local/bin"}
TOOLBOX_DIR="$INSTALL_DIR/toolbox"

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

# Install Amp Toolbox wrappers
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
TOOLBOX_TOOLS=("init" "start" "recite" "log" "archive" "view" "agents" "delete" "version")

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

# Shell configuration check
if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
    echo ""
    echo "WARNING: $INSTALL_DIR is not in your PATH."
    echo "You may need to add it to your shell configuration (e.g., .bashrc, .zshrc):"
    echo "  export PATH=\$PATH:$INSTALL_DIR"
fi

# Amp Toolbox setup instructions
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

echo "Installation complete! Run 'cdd' to get started!"
