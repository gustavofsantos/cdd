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