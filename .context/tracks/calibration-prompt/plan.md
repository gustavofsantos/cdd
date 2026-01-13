# Plan: Integrate Calibration Prompt

## Phase 0: Analysis
- [x] Analyze existing prompt integration mechanism. (Reusing knowledge from previous track)

## Phase 1: TDD Loop (Red-Green-Refactor)

### Task 1: Verify Calibration Prompt Embedding
- [x] Add test case to `prompts/integration_test.go` for `prompts.Calibration`. (1c106ec)
- [x] Run tests (should fail to compile). (1c106ec)
- [x] Update `prompts/prompts.go` to embed `calibration.md`. (1c106ec)
- [x] Run tests (should pass). (1c106ec)

### Task 2: Verify CLI Flag
- [x] Add test case to `internal/cmd/init_test.go` for `--calibration-prompt`. (1c106ec)
- [x] Run tests (should fail). (1c106ec)
- [x] Update `internal/cmd/init.go` to add `--calibration-prompt` flag and logic. (1c106ec)
- [x] Run tests (should pass). (1c106ec)

## Phase 2: Consolidation
- [x] Update `README.md` to include `--calibration-prompt`. (1c106ec)
- [x] Final check and Archive. (1c106ec)
