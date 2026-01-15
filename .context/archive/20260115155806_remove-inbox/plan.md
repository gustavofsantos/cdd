# Plan for remove-inbox

## Phase 0: Analysis
- [x] 🗣️ Phase 0: Alignment & Requirements (Fill `spec.md` using EARS)

## Phase 1: Architecture
- [x] 📝 Phase 1: Approval (User signs off)

## Phase 2: Implementation
- [x] 🔴 Refactor `internal/cmd/archive.go` to remove inbox appending logic
- [x] 🔴 Refactor `internal/cmd/archive_test.go` to remove inbox tests
- [x] 🔴 Refactor `internal/cmd/view.go` to remove inbox flag and logic
- [x] 🔴 Refactor `internal/cmd/view_test.go` to remove inbox tests
- [x] 🔴 Delete `internal/cmd/util.go` and `internal/cmd/util_test.go`
- [x] 🔴 Refactor `internal/cmd/root.go` to update description
- [x] 🔴 Update `README.md` to remove inbox references
