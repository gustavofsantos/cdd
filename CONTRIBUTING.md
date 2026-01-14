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

The project includes a comprehensive test suite. Tests follow the TDD pattern (Red-Green-Refactor) and use the CDD workflow.

To run tests:

```bash
go test ./...
```

Run tests for a specific package:

```bash
go test ./internal/cmd
```

Run a specific test:

```bash
go test ./internal/cmd -run TestNamePattern
```

**Testing Agent Skills:**
When adding support for new agent integrations (like Antigravity), follow the pattern in `internal/cmd/agents.go`:

1. Create a dedicated installation function (e.g., `installAntigravitySkill`)
2. Add validation for the target format
3. Write comprehensive tests covering happy path, edge cases, and validation
4. Wire the handler into the command's switch statement

## Code Style

- Follow standard Go conventions (Effective Go).
- Ensure your code is formatted with `gofmt`.

## Submitting Changes

1.  Commit your changes with descriptive commit messages.
2.  Push your branch to your fork.
3.  Open a **Pull Request** against the main repository.

## Reporting Issues

If you find a bug or have a feature request, please open an issue on the GitHub repository.
