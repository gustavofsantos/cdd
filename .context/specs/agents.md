# Agents Specification

## 1. Overview
The `agents` command manages the integration between the CDD CLI and AI agents. Its primary responsibility is to install and maintain the Agent Skill that defines the CDD protocol for AI orchestration.

## 2. Requirements

### 2.1 Skill Installation (Directory-Based Targets)
- The command `cdd agents --install` must install the CDD System Prompt as an Agent Skill.
- **Location**: It creates a directory named `.agent/skills/cdd/`.
- **Artifact**: It creates a file named `SKILL.md` inside that directory.
- **Metadata**: The `SKILL.md` file must include the following YAML frontmatter:
  ```yaml
  ---
  name: cdd
  version: 2
  description: Protocol for implementing software features using the Context-Driven Development methodology.
  ---
  ```
- **Content**: The body of `SKILL.md` must contain the full content of the CDD System Prompt (sourced from `prompts/system.md`).
- **Supported Targets**: `agent`, `agents`, `claude` (directory-based)

### 2.1.1 Antigravity Target
- The command `cdd agents --install --target antigravity` must install all five CDD skills to `.agent/skills/` in Antigravity-compatible format.
- **Location**: Creates `.agent/skills/{skill-id}/SKILL.md` for each of the five skills.
- **Skill IDs**: `cdd`, `cdd-analyst`, `cdd-architect`, `cdd-executor`, `cdd-integrator`.
- **Format Validation**: Each skill must be validated to contain required YAML frontmatter fields:
  - `name`: The unique identifier for the skill (lowercase, hyphens for spaces)
  - `description`: Clear description of what the skill does and when to use it
  - Optional: `metadata.version` for versioning
- **Idempotent Installation**: If a skill with the same version already exists, skip installation.
- **Version Tracking**: Uses the same version extraction mechanism as other targets.
- **Compatibility**: Follows Google Antigravity's skill standard for workspace-scoped skills.

### 2.2 Cursor Rules Installation
- The command `cdd agents --install --target cursor` must generate a `.cursorrules` file in the project root.
- **Rationale**: Cursor does not support Agent Skills directly; it requires a flat rules file.
- **Content**: The `.cursorrules` file concatenates all five CDD agent skills with clear section separators.
- **Format**: Markdown with YAML frontmatter containing version metadata at the top.
- **Features**:
  - Idempotent: Second install with same version skips overwrite
  - Version tracking: Extracts version from skill metadata
  - Backup on update: Creates `.cursorrules.bak` when upgrading

### 2.3 Versioning and Migration
- The command should track the version of the installed skill.
- If a newer version of the skill is available in the binary, it should offer to migrate/update the existing one.
- If an update occurs, the previous version should be backed up with a `.bak` extension.
- Version extraction works consistently across all target types (directory and cursor).

## 3. Relevant Context
- `internal/cmd/agents.go`: Implementation of the `agents` command.
- `prompts/system.md`, `prompts/analyst.md`, `prompts/architect.md`, `prompts/executor.md`, `prompts/integrator.md`: Source of skill content.
- `.agent/skills/cdd/SKILL.md` and other skill directories: Directory-based installations.
- `.cursorrules`: Cursor rules file generated for Cursor editor integration.
- Tests: `agents_install_test.go`, `agents_cursor_test.go`, `agents_target_cursor_test.go`, `agents_antigravity_e2e_test.go`, `agents_install_antigravity_test.go`, `agents_antigravity_discovery_test.go`
- Google Antigravity Documentation: https://antigravity.google/docs/skills
