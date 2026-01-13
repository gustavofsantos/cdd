# Plan: Integrate Executor and Planner Prompts

## Phase 0: Analysis
- [x] Analyze existing prompt integration mechanism.

## Phase 1: TDD Loop (Red-Green-Refactor)

### Task 1: Verify Prompts Embedding
- [x] Write test in `prompts/integration_test.go` to ensure `prompts.Executor` and `prompts.Planner` are not empty. (04971f8)
- [x] Run tests (should fail to compile because `Executor` and `Planner` don't exist in `prompts` package). (04971f8)
- [x] Update `prompts/prompts.go` to embed `executor.md` and `planner.md`. (04971f8)
- [x] Run tests (should pass). (04971f8)

### Task 2: Verify CLI Flags
- [x] Write a test for the `init` command flags in `internal/cmd/init_test.go`. (04971f8)
- [x] Update `internal/cmd/init.go` to add `--executor-prompt` and `--planner-prompt`. (04971f8)
- [x] Verify manually or with a script that `cdd init --executor-prompt` works. (04971f8)

## Phase 2: Consolidation
- [x] Update documentation if necessary (README.md). (04971f8)
- [x] Cleanup and Archive.
