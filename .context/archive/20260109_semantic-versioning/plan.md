# Plan for semantic-versioning

- [x] 🗣️ Phase 0: Alignment & Analysis (Fill spec.md)
- [x] 📝 Phase 1: Approval (User signs off)
- [x] 🛠️ Phase 2: Implementation (TDD Loop)
    - [x] Initialize GoReleaser configuration (`.goreleaser.yaml`).
    - [x] Update `internal/cmd/version.go` to support ldflags injection.
    - [x] Refactor `.github/workflows/release.yml` to use `goreleaser-action`.
    - [x] Verify build and version injection locally (`goreleaser build --snapshot`).
- [x] ✅ Phase 3: Finalization
    - [x] Update documentation.
    - [x] Archive track.
