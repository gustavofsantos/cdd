# Plan: Integrate Calibration Prompt

## Phase 0: Analysis
- [x] Analyze existing prompt integration mechanism. (Reusing knowledge from previous track)

## Phase 1: TDD Loop (Red-Green-Refactor)

### Task 1: Verify Calibration Prompt Embedding
- [ ] Add test case to `prompts/integration_test.go` for `prompts.Calibration`.
- [ ] Run tests (should fail to compile).
- [ ] Update `prompts/prompts.go` to embed `calibration.md`.
- [ ] Run tests (should pass).

### Task 2: Verify CLI Flag
- [ ] Add test case to `internal/cmd/init_test.go` for `--calibration-prompt`.
- [ ] Run tests (should fail).
- [ ] Update `internal/cmd/init.go` to add `--calibration-prompt` flag and logic.
- [ ] Run tests (should pass).

## Phase 2: Consolidation
- [ ] Update `README.md` to include `--calibration-prompt`.
- [ ] Final check and Archive.
