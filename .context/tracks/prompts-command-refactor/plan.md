# Plan: Refactor Prompt Printing to Dedicated Command

## Phase 0: Analysis
- [x] Identify all prompt-related flags and variables in `internal/cmd/init.go`.

## Phase 1: TDD Loop (Red-Green-Refactor)

### Task 1: Create `prompts` command
- [x] Write a test in `internal/cmd/prompts_test.go` to verify the new `prompts` command and its flags. (3497c8d)
- [x] Run tests (should fail). (3497c8d)
- [x] Create `internal/cmd/prompts.go` with the `promptsCmd` and its flags. (3497c8d)
- [x] Run tests (should pass). (3497c8d)

### Task 2: Refactor `init` command
- [x] Update `internal/cmd/init_test.go` to remove checks for prompt flags on `init`. (3497c8d)
- [x] Update `internal/cmd/init.go` to remove prompt flags and related logic. (3497c8d)
- [x] Run tests (ensure `init` still works for basic initialization). (3497c8d)

## Phase 2: Consolidation
- [x] Update `README.md` to reflect the new command structure. (3497c8d)
- [x] Update `internal/cmd/init.go` success message to point to the new command. (3497c8d)
- [x] Final check and Archive.
