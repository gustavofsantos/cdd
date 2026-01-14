# Plan for agent-skill-installation
- [x] 🗣️ Phase 0: Alignment & Analysis (Fill spec.md)
- [x] 📝 Phase 1: Approval (User signs off)

## Phase 2: Implementation
- [x] 🔴 Test: Verify `cdd prompts --install` creates `.agent/skills/cdd/SKILL.md` with correct content
- [x] 🟢 Impl: Refactor `prompts.go` to support `FileSystem` and implement `--install` logic
- [x] 🔵 Refactor: Ensure the frontmatter matches the updated spec exactly
- [x] 🔴 Test: Ensure `system_test.go` checks for Agent Skill mentions and NOT `AGENTS.local.md`
- [x] 🟢 Impl: Remove external references from `prompts/system.md` and pivot to Agent Skills

