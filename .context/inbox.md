

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


## Updates from Track: update-readme (Fri Jan 9 19:16:20 -03 2026)
- Refactored project documentation to improve "sellability" and organization.
- Created `INSTALLATION.md` to house platform-specific setup instructions, keeping the main README clean.
- Added a comprehensive "Why CDD?" (Rationale) section to `README.md`, detailing the Strategist/Tactician philosophy, cost-efficiency of small models, and the project's lineage from OpenSpec, Conductor, etc.
- Added a "Target Audience" section to `README.md` to clearly define who the tool is for.
- Cleaned up redundant instructions in `README.md` and added a proper "Getting Started" flow.
