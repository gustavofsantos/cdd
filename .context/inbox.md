
---
###### Archived at: 2026-01-14 16:28:50 | Track: integrate-new-skills

# Track: integrate-new-skills

## 1. User Intent
Update the `cdd agents --install` command to install multiple Agent Skills (cdd, analyst, architect, executor, integrator) instead of just one. The `system.md` prompt will serve as the orchestrator.

## 2. Relevant Context
- `prompts/system.md`: Orchestrator prompt.
- `prompts/analyst.md`: Analyst prompt.
- `prompts/architect.md`: Architect prompt.
- `prompts/executor.md`: Executor prompt.
- `prompts/integrator.md`: Integrator prompt.
- `prompts/prompts.go`: Embedding of prompt files.
- `internal/cmd/agents.go`: Implementation of the `agents` command.

## 3. Requirements (EARS)
- The system shall embed all five new prompt files in `prompts/prompts.go`.
- When `cdd agents --install` is executed, the system shall install the orchestrator skill (cdd) from `system.md` into `.agent/skills/cdd/SKILL.md`.
- When `cdd agents --install` is executed, the system shall install the analyst skill (cdd-analyst) from `analyst.md` into `.agent/skills/cdd-analyst/SKILL.md`.
- When `cdd agents --install` is executed, the system shall install the architect skill (cdd-architect) from `architect.md` into `.agent/skills/cdd-architect/SKILL.md`.
- When `cdd agents --install` is executed, the system shall install the executor skill (cdd-executor) from `executor.md` into `.agent/skills/cdd-executor/SKILL.md`.
- When `cdd agents --install` is executed, the system shall install the integrator skill (cdd-integrator) from `integrator.md` into `.agent/skills/cdd-integrator/SKILL.md`.
- The system shall maintain versioning and migration logic for all installed skills.
- If a skill already exists with the same or higher version, then the system shall skip installation for that specific skill.
- If a legacy skill exists, then the system shall back it up before installing the new version.

---
###### Archived at: 2026-01-14 16:36:03 | Track: multi-directory-skills

# Track: multi-directory-skills

## 1. User Intent
Enable the `cdd agents --install` command to target either the `.agent` (or `.agents`) directory or the `.claude` directory for skill installation, allowing users to choose based on their AI IDE/orchestrator preferences.

## 2. Relevant Context
- `internal/cmd/agents.go`: Contains the `installSkill` and `NewAgentsCmd` logic.
- `.agent/skills`: Current default installation path.
- The user specifically mentioned `.agents` (plural) and `.claude`.

## 3. Requirements (EARS)
- The `cdd agents --install` command shall allow the user to specify the target provider (e.g., `agent` or `claude`).
- When the target is `agent`, the system shall install skills into `.agent/skills/` (defaulting to the existing singular form, but supporting the plural `.agents` if explicitly requested or as the new standard if clarified). *Decision: I will use a flag to select the target.*
- When the target is `claude`, the system shall install skills into `.claude/skills/`.
- If no target is specified, then the system shall default to `.agent/skills/`.
- The system shall maintain the same versioning and backup logic regardless of the target directory.
