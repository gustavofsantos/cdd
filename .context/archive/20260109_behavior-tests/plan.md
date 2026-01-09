# Plan for behavior-tests
- [x] 🗣️ Phase 0: Alignment & Analysis (Fill spec.md)
- [x] 📝 Phase 1: Approval (User signs off)

## Refactoring Core
- [x] Define `FileSystem` interface in `internal/platform/fs.go` (and `MockFileSystem`).
- [x] Refactor `start.go` to use `FileSystem` and `cmd.Printf`.
- [x] Implement `start_test.go`.
- [x] Refactor `recite.go` to use `FileSystem` and `cmd.Printf`.
- [x] Implement `recite_test.go`.
- [x] Refactor `list.go` to use `FileSystem` and `cmd.Printf`.
- [x] Implement `list_test.go`.
- [x] Refactor `log.go` to use `FileSystem` and `cmd.Printf`.
- [x] Implement `log_test.go`.
- [x] Refactor `archive.go` to use `FileSystem` and `cmd.Printf`.
- [x] Implement `archive_test.go`.
- [x] Refactor `dump.go` to use `FileSystem` and `cmd.Printf`.
- [x] Implement `dump_test.go`.

