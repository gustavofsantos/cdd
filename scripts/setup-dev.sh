#!/bin/sh
set -e

# Define colors
GREEN='\033[0;32m'
NC='\033[0m' # No Color

echo "Setting up development environment..."

# 1. Install golangci-lint if not present
if ! command -v golangci-lint >/dev/null 2>&1; then
    echo "golangci-lint not found. Installing..."

    # Check if GOPATH is set, otherwise use default
    GOPATH=$(go env GOPATH)
    if [ -z "$GOPATH" ]; then
        GOPATH="$HOME/go"
    fi

    BIN_DIR="$GOPATH/bin"
    mkdir -p "$BIN_DIR"

    # Install specific version or latest
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "$BIN_DIR" v1.64.6

    echo "${GREEN}golangci-lint installed to $BIN_DIR${NC}"
    echo "Please ensure $BIN_DIR is in your PATH."
else
    echo "${GREEN}golangci-lint is already installed.${NC}"
fi

# 2. Setup git hooks
HOOK_FILE=".git/hooks/pre-push"
echo "Configuring pre-push git hook..."

mkdir -p .git/hooks

cat > "$HOOK_FILE" << 'EOF'
#!/bin/sh
set -e

# Define colors
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m'

echo "Running pre-push checks..."

# Ensure GOPATH/bin is in PATH for this hook
export PATH=$PATH:$(go env GOPATH)/bin

# Run linter
echo "Running golangci-lint..."
if ! golangci-lint run ./...; then
    echo "${RED}Linting failed!${NC}"
    exit 1
fi

# Run build
echo "Running go build..."
if ! go build ./...; then
    echo "${RED}Build failed!${NC}"
    exit 1
fi

# Run tests
echo "Running go test..."
if ! go test ./...; then
    echo "${RED}Tests failed!${NC}"
    exit 1
fi

echo "${GREEN}All checks passed!${NC}"
EOF

chmod +x "$HOOK_FILE"
echo "${GREEN}Git pre-push hook installed successfully.${NC}"
