# Plan: Leaner CDD Process Improvements

## Tasks

- [x] **Test Archive Timestamp and Cleanup**
    - [x] Update `internal/cmd/archive_test.go` to assert that `scratchpad.md` is removed and the destination directory uses the full timestamp.
    - [x] Run tests to confirm they fail (Red Phase).

- [x] **Implement Cleanup and Precise Timestamping**
    - [x] Modify `internal/cmd/archive.go` to delete `scratchpad.md` before archiving.
    - [x] Modify `internal/cmd/archive.go` to use `20060102150405` format for the archive directory.
    - [x] Run tests to confirm they pass (Green Phase).

- [x] **Consolidation**
    - [x] Update `spec.md` with test references.
    - [x] Archive the track.
