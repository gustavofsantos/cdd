# Track: antigravity-skills-support

## 1. User Intent
Add Antigravity skill installation support to the `agents` command. The CDD tool should be able to generate and install Antigravity-compatible skills alongside existing Claude and Cursor integrations, using the `--install --target antigravity` flags.

## 2. Relevant Context
- `internal/cmd/agents.go` - Current implementation of `--install` for Claude and Cursor
- `prompts/` - Source markdown files for CDD skills (system, analyst, architect, executor, integrator)
- https://antigravity.google/docs/skills - Antigravity skills standard documentation
- `.agent/skills/` - Target directory for workspace-scoped Antigravity skills (per Antigravity spec)

## 3. Requirements (EARS)

### Ubiquitous
- The `agents` command shall support `--target antigravity` as a valid installation target.
- The system shall install CDD skills into the `.agent/skills/` directory structure when `--target antigravity` is specified.
- Each skill shall be installed in its own folder with a `SKILL.md` file following Antigravity's format (YAML frontmatter + instructions).
- The system shall use the skill name as the folder identifier (e.g., `.agent/skills/cdd-system/`, `.agent/skills/cdd-analyst/`, etc.).

### Event-driven
- When `cdd agents --install --target antigravity` is executed, the system shall generate and install all five CDD skills (Orchestrator, Analyst, Architect, Executor, Integrator) into `.agent/skills/`.
- When a skill already exists at the target location, the system shall prompt the user to confirm overwrite (or silently overwrite with a flag option).

### State-driven
- While the installation is in progress, the system shall provide feedback on which skills are being installed.
- While skills are written to disk, the system shall validate that each `SKILL.md` file contains required YAML frontmatter (`name` and `description` fields).

### Unwanted Behavior
- If the `.agent/` directory does not exist, the system shall create it before writing skills.
- If a skill's `SKILL.md` is invalid or missing required fields, the system shall fail with a clear error message.

### Optional
- Where verbose output is enabled (e.g., with a `--verbose` flag), the system shall print detailed installation progress and file paths for each skill.