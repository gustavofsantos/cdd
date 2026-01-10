#!/bin/bash

# CDD Installation Script
set -e

INSTALL_DIR=${1:-"/usr/local/bin"}

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

# Shell configuration check
if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
    echo "WARNING: $INSTALL_DIR is not in your PATH."
    echo "You may need to add it to your shell configuration (e.g., .bashrc, .zshrc):"
    echo "  export PATH=\$PATH:$INSTALL_DIR"
fi

echo "Run 'cdd' to get started!"
