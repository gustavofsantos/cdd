# Plan for update-archive-metadata

- [x] 🔴 Test: Create a reproduction script `test_archive_metadata.sh` that validates the presence of `archived_at` in the archived metadata.
- [x] 🟢 Impl: Modify `internal/cmd/archive.go` to update `metadata.json` before archiving.
- [x] 🔵 Verification: Run `test_archive_metadata.sh` and ensure it passes.
