# Plan for cleanup-unused-prompts

## Phase 0: Analysis
- [x] 🗣️ Phase 0: Alignment & Requirements (Fill `spec.md` using EARS)

## Phase 1: Architecture
- [x] 📝 Phase 1: Approval (User signs off)

## Phase 2: Implementation
- [x] Remove unused prompt files (`bootstrap.md`, `calibration.md`, `migration.md`)
- [x] Clean up `prompts/prompts.go` to remove unused embedded variables
- [x] Update `prompts/integration_test.go` to remove tests for deleted prompts
- [x] Fix `prompts/system_test.go` to align with the new lean system prompt
- [x] Verify all tests in `prompts/` pass
