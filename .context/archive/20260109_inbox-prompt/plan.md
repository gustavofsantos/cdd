# Plan: Inbox Prompt Support

## Phase 1: Analysis & Skeleton
- [x] Analyze `inbox.md` intent and `archive.go` implementation.
- [x] Create `prompts/inbox.md` with appropriate instructions.
- [x] Update `prompts/prompts.go` to embed the new prompt.

## Phase 2: CLI Implementation
- [x] Modify `internal/cmd/init.go` to add `--inbox-prompt` flag.
- [x] Implement the logic to output the prompt when the flag is present.

## Phase 3: Verification
- [x] Verify that `go run cmd/cdd/main.go init --inbox-prompt` works as expected.
- [x] Verify it doesn't break existing `init` behavior.
