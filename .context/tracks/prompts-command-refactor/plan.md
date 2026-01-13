# Plan: Refactor Prompt Printing to Dedicated Command

## Phase 0: Analysis
- [x] Identify all prompt-related flags and variables in `internal/cmd/init.go`.

## Phase 1: TDD Loop (Red-Green-Refactor)

### Task 1: Create `prompts` command
- [ ] Write a test in `internal/cmd/prompts_test.go` to verify the new `prompts` command and its flags.
- [ ] Run tests (should fail).
- [ ] Create `internal/cmd/prompts.go` with the `promptsCmd` and its flags.
- [ ] Run tests (should pass).

### Task 2: Refactor `init` command
- [ ] Update `internal/cmd/init_test.go` to remove checks for prompt flags on `init`.
- [ ] Update `internal/cmd/init.go` to remove prompt flags and related logic.
- [ ] Run tests (ensure `init` still works for basic initialization).

## Phase 2: Consolidation
- [ ] Update `README.md` to reflect the new command structure.
- [ ] Update `internal/cmd/init.go` success message to point to the new command.
- [ ] Final check and Archive.
