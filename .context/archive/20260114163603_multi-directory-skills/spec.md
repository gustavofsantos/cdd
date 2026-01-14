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