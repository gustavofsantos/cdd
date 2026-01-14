# Plan for multi-directory-skills

## Phase 0: Analysis
- [x] 🗣️ Phase 0: Alignment & Requirements (Fill `spec.md` using EARS)

## Phase 1: Architecture
- [x] 📝 Phase 1: Approval (User signs off)

## Phase 2: Implementation
- [x] Add `installTarget` flag to `agents` command
- [x] Refactor `installSkill` to take a base directory
- [x] Update `Run` logic to set base directory based on `--target`
- [x] Add tests for both `agent` and `claude` targets
- [x] Update documentation/help text to mention the new flag
