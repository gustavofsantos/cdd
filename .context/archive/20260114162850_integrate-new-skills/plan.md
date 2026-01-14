# Plan for integrate-new-skills

## Phase 0: Analysis
- [x] 🗣️ Phase 0: Alignment & Requirements (Fill `spec.md` using EARS)

## Phase 1: Architecture
- [x] 📝 Phase 1: Approval (User signs off)

## Phase 2: Implementation
- [x] Embed new prompt files in `prompts/prompts.go`
- [x] Refactor `internal/cmd/agents.go` to handle multiple skills
- [x] Implement `installSkill` function with version parsing and migration logic
- [x] Update `agents --install` to install all five skills (cdd, cdd-analyst, cdd-architect, cdd-executor, cdd-integrator)
- [x] Verify installation by running `cdd agents --install` and checking `.agent/skills` directory
