# Workflow

## CLI Commands
The project is driven by the `cdd` CLI tool.

- **`go run cmd/cdd/main.go init`**: Initializes the `.context` directory.
    - `--inbox-prompt`: Retrieves the Context Gardener prompt for consolidation.
- **`go run cmd/cdd/main.go start <track>`**: Creates a new track directory in `.context/tracks/`.
- **`go run cmd/cdd/main.go recite <track>`**: Displays the `spec.md` and `plan.md` for a track.
- **`go run cmd/cdd/main.go log <track>`**: Appends to `decisions.md`.
- **`go run cmd/cdd/main.go dump <track>`**: Pipes stdin to `scratchpad.md`.
- **`go run cmd/cdd/main.go archive <track>`**: Moves track to `.context/archive/`, promotes spec, and appends `context_updates.md` to `.context/inbox.md`.
    - Proactively checks if `inbox.md` > 50 lines and suggests gardening.
- **`go run cmd/cdd/main.go list`**: Lists active tracks.
- **`go run cmd/cdd/main.go version`**: Displays version information.

## Development Flow (CDD)
1. **Recite**: Always check the plan.
2. **Spec**: Define *what* in `spec.md`.
3. **Plan**: Define *how* in `plan.md`.
4. **Implement**: Write code (TDD Recommended).
5. **Archive**: Close track and move updates to Inbox.
6. **Garden**: Periodically promote Inbox updates to Global Context.

## Testing
- **Framework:** Go `testing` package.
- **Approach:** Dependency Injection via `FileSystem` interface allows mocking file operations.
- **Execution:** Run `go test ./...` to verify behavior.

## Documentation
- **README.md**: High-level rationale and audience.
- **INSTALLATION.md**: Platform-specific setup instructions.

## Dogfooding (Self-Host)
Since this project is the source code of the tool itself, always use the latest version to drive development:
- **Run**: `go run cmd/cdd/main.go <command> <args>`
- **Example**: `go run cmd/cdd/main.go recite setup`
