

## Updates from Track: setup (Fri Jan 9 18:32:19 -03 2026)
Initialized Global Context for the `cdd` project.
- Mapped project structure (Go/Cobra).
- Identified key directories: `cmd/cdd` (entry), `internal/cmd` (logic).
- Identified lack of tests.
- Populated `product.md`, `tech-stack.md`, `patterns.md`, `workflow.md`.
- **Dogfooding**: Established convention to run `go run cmd/cdd/main.go` to use the latest version of the tool during development.


## Updates from Track: behavior-tests (Fri Jan 9 18:47:05 -03 2026)
# Proposed Global Context Updates
> Add notes here if product.md or tech-stack.md needs updating.

- Refactored CLI commands to use `FileSystem` interface (Dependency Injection) for better testability.
- Defined `FileSystem` interface and `MockFileSystem` in `internal/platform`.
- Added comprehensive unit tests for `start`, `recite`, `list`, `log`, `archive`, `dump` commands.


## Updates from Track: semantic-versioning (Fri Jan 9 18:58:44 -03 2026)
# Context Updates: semantic-versioning

- Added `goreleaser` for automated builds and releases.
- Configuration in `.goreleaser.yaml`.
- Added `cdd version` command to the CLI.
- Versions are now injected at build time via ldflags into `internal/cmd.Version`.
- Updated GitHub Actions workflow `.github/workflows/release.yml` to use GoReleaser.
- Added `dist/` to `.gitignore`.
