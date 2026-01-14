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