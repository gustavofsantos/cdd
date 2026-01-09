# Plan - Inbox Cleanup Reminder

## Phase 1: Preparation & Testing Infrastucture
- [x] Create a helper function `CheckInboxSize(fs platform.FileSystem, cmd *cobra.Command)` in a new file `internal/cmd/util.go` or similar.
- [x] Add a test case for `CheckInboxSize` (or integration test in `archive_test.go`).

## Phase 2: Implementation (Red-Green-Refactor)
- [x] **Red**: Add a test in `internal/cmd/archive_test.go` that verifies the suggestion is printed when `inbox.md` > 50 lines.
- [x] **Green**: Implement `CheckInboxSize` and call it in `archive.go` after writing to `inbox.md`.
- [x] **Refactor**: Ensure the message is clear and follows the user's request.

## Phase 3: Verification
- [x] Manually verify with a dummy track and a large `inbox.md`.
