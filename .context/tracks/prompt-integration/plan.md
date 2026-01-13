# Plan: Integrate Executor and Planner Prompts

## Phase 0: Analysis
- [x] Analyze existing prompt integration mechanism.

## Phase 1: TDD Loop (Red-Green-Refactor)

### Task 1: Verify Prompts Embedding
- [ ] Write test in `prompts/integration_test.go` to ensure `prompts.Executor` and `prompts.Planner` are not empty.
- [ ] Run tests (should fail to compile because `Executor` and `Planner` don't exist in `prompts` package).
- [ ] Update `prompts/prompts.go` to embed `executor.md` and `planner.md`.
- [ ] Run tests (should pass).

### Task 2: Verify CLI Flags
- [ ] Write a test for the `init` command flags in `internal/cmd/init_test.go` (if it doesn't exist, create it or use a manual check if unit testing cobra commands is too complex for this project's current state).
- [ ] Update `internal/cmd/init.go` to add `--executor-prompt` and `--planner-prompt`.
- [ ] Verify manually or with a script that `cdd init --executor-prompt` works.

## Phase 2: Consolidation
- [ ] Update documentation if necessary (README.md).
- [ ] Cleanup and Archive.
