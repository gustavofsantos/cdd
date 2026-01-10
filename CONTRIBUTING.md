# Contributing to CDD Tool Suite

Thank you for your interest in contributing to the CDD Tool Suite! We welcome contributions from everyone.

## Getting Started

1.  **Fork the repository** on GitHub.
2.  **Clone your fork** locally.
3.  **Create a branch** for your feature or bug fix.

## Development Environment

- Ensure you have **Go 1.24+** installed.
- The project uses standard Go modules.
- Run the setup script to configure your environment (installs `golangci-lint` and git hooks):

```bash
sh scripts/setup-dev.sh
```

## Building and Testing

To build the project:

```bash
go build -o cdd cmd/cdd/main.go
```

Currently, the project does not have an extensive test suite. We encourage you to add unit tests for any new logic you introduce.

To run tests (if you add any):

```bash
go test ./...
```

## Code Style

- Follow standard Go conventions (Effective Go).
- Ensure your code is formatted with `gofmt`.

## Submitting Changes

1.  Commit your changes with descriptive commit messages.
2.  Push your branch to your fork.
3.  Open a **Pull Request** against the main repository.

## Reporting Issues

If you find a bug or have a feature request, please open an issue on the GitHub repository.
