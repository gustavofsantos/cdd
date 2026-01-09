# Context Updates: semantic-versioning

- Added `goreleaser` for automated builds and releases.
- Configuration in `.goreleaser.yaml`.
- Added `cdd version` command to the CLI.
- Versions are now injected at build time via ldflags into `internal/cmd.Version`.
- Updated GitHub Actions workflow `.github/workflows/release.yml` to use GoReleaser.
- Added `dist/` to `.gitignore`.
