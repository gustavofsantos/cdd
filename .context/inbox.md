
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

---
###### Archived at: 2026-01-14 16:43:25 | Track: cleanup-unused-prompts

# Track: cleanup-unused-prompts

## 1. User Intent
Perform a clean up task to remove unused prompts and update the codebase to reflect these changes, ensuring all tests pass.

## 2. Relevant Context
- `prompts/bootstrap.md`: Unused prompt.
- `prompts/calibration.md`: Unused prompt.
- `prompts/migration.md`: Unused prompt.
- `prompts/prompts.go`: Contains embedded variables for prompts.
- `prompts/system_test.go`: Contains tests for the system prompt that are currently failing.
- `prompts/integration_test.go`: Uses some of the unused prompts.

## 3. Requirements (EARS)

- The system shall remove `prompts/bootstrap.md`, `prompts/calibration.md`, and `prompts/migration.md`.
- The system shall remove `Bootstrap` and `Calibration` variables from `prompts/prompts.go`.
- The system shall update `prompts/system_test.go` to match the current lean system prompt.
- The system shall update `prompts/integration_test.go` to remove references to deleted prompts.
- The system shall ensure `go test ./prompts/...` passes after cleanup.

---
###### Archived at: 2026-01-14 18:55:38 | Track: cursor-rules-support

# Track: cursor-rules-support

## 1. User Intent
Extend the `cdd agents --install` command to support Cursor editor by installing skills as `.cursorrules` files. Cursor does not support loading Agent Skills natively, so skills must be converted to Cursor rules format for compatibility.

## 2. Relevant Context
- `internal/cmd/agents.go`: Current implementation of agents command
- `prompts/system.md`, `prompts/analyst.md`, etc.: Skill content sources
- `.context/specs/agents.md`: Existing agents specification
- Current targets: `agent` (.agent/skills/), `agents` (.agents/skills/), `claude` (.claude/)
- New target: `cursor` (.cursorrules file format)

## 3. Requirements (EARS)

### Ubiquitous Requirements
- The system shall accept `--target cursor` flag in the agents install command.
- The system shall generate a single `.cursorrules` file in the project root directory.
- The system shall include all five CDD agent skills (cdd, cdd-analyst, cdd-architect, cdd-executor, cdd-integrator) in the `.cursorrules` content.
- The system shall include version metadata for the `.cursorrules` file for tracking and updates.

### Event-Driven Requirements
- When the user runs `cdd agents --install --target cursor`, the system shall create or update `.cursorrules` with concatenated skill content.
- When a newer version of skills is available, the system shall back up the existing `.cursorrules` file with a `.cursorrules.bak` extension before updating.

### State-Driven Requirements
- While the `.cursorrules` file is up-to-date, the system shall skip installation and report that it is current.
- While updating from an older version, the system shall preserve the backup of the previous version.

### Unwanted Behavior
- If `.cursorrules` already exists with the same version, the system shall not overwrite it.
- If version extraction fails, the system shall treat it as version "0.0.0" and proceed with update.

### Optional Requirements
- Where help documentation is displayed, the system shall list `cursor` as a valid `--target` option.
