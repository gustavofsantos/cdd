# Plan for view-raw-mode
- [x] 🗣️ Phase 0: Alignment & Analysis (Fill spec.md)
- [x] 📝 Phase 1: Approval (User signs off)

## Phase 2: Implementation
- [x] 🔴 Test: Add test case for raw output in `view_test.go`
- [x] 🟢 Impl: Add `--raw` flag and TTY detection to `view.go`
- [x] 🟢 Impl: Update `buildViewMarkdown` to return raw list when requested
- [x] 🔵 Refactor: Ensure consistent behavior between TTY detection and explicit flag
