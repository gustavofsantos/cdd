# Workflow

## CLI Commands
The project is driven by the `cdd` CLI tool.

- **`go run cmd/cdd/main.go init`**: Initializes the `.context` directory and `setup` track.
- **`go run cmd/cdd/main.go start <track>`**: Creates a new track directory in `.context/tracks/`.
- **`go run cmd/cdd/main.go recite <track>`**: Displays the `spec.md` and `plan.md` for a track.
- **`go run cmd/cdd/main.go log <track>`**: Appends to `decisions.md`.
- **`go run cmd/cdd/main.go dump <track>`**: Pipes stdin to `scratchpad.md`.
- **`go run cmd/cdd/main.go archive <track>`**: Moves track to `.context/archive/` and promotes spec to `.context/features/`.
- **`go run cmd/cdd/main.go list`**: Lists active tracks.

## Development Flow (CDD)
1. **Recite**: Always check the plan.
2. **Spec**: Define *what* in `spec.md`.
3. **Plan**: Define *how* in `plan.md`.
4. **Implement**: Write code.
5. **Archive**: Close track.

## Dogfooding (Self-Host)
Since this project is the source code of the tool itself, always use the latest version to drive development:
- **Run**: `go run cmd/cdd/main.go <command> <args>`
- **Example**: `go run cmd/cdd/main.go recite setup`

## Testing
- **Current Status**: No automated tests found.
- **Goal**: Introduce testing framework (likely Go standard `testing` package).
