# Agents Specification

## 1. Overview
The `agents` command manages the integration between the CDD CLI and AI agents. Its primary responsibility is to install and maintain the Agent Skill that defines the CDD protocol for AI orchestration.

## 2. Requirements

### 2.1 Skill Installation (Directory-Based Targets)
- The command `cdd agents --install --target <target>` must install the CDD System Prompt as an Agent Skill.
- **Requirement**: A `--target` flag must be explicitly specified. The command shall not default to any target.
- **Location**: It creates a directory named `.agent/skills/cdd/` (or `.claude/skills/`, `.agents/skills/` depending on target).
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
- **Error Handling**: If neither `--target` nor `--all` is specified, the command shall return an error message with examples of correct usage.

### 2.1.1 Antigravity Target
- The command `cdd agents --install --target antigravity` must install all six CDD skills to `.agent/skills/` in Antigravity-compatible format.
- **Location**: Creates `.agent/skills/{skill-id}/SKILL.md` for each of the six skills.
- **Skill IDs**: `cdd`, `cdd-surveyor`, `cdd-analyst`, `cdd-architect`, `cdd-executor`, `cdd-integrator`.
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
- **Content**: The `.cursorrules` file concatenates all six CDD agent skills with clear section separators.
- **Format**: Markdown with YAML frontmatter containing version metadata at the top.
- **Features**:
  - Idempotent: Second install with same version skips overwrite
  - Version tracking: Extracts version from skill metadata
  - Backup on update: Creates `.cursorrules.bak` when upgrading

### 2.3 Multi-Platform Installation (`--all` Flag)
- The command `cdd agents --install --all` shall install all six CDD skills for all supported platforms in a single invocation.
- **Platforms Covered**: The `--all` flag installs skills for all directory-based targets (`agent`, `agents`, `claude`) plus special targets (`cursor`, `antigravity`).
- **Idempotent**: Each platform installation follows its respective idempotent rules (version checking, backups on update).
- **Behavior**: When `--all` is used, it overrides any `--target` specification.

### 2.4 Versioning and Migration
- The command should track the version of the installed skill.
- If a newer version of the skill is available in the binary, it should offer to migrate/update the existing one.
- If an update occurs, the previous version should be backed up with a `.bak` extension.
- Version extraction works consistently across all target types (directory and cursor).

## 3. Relevant Context
- `internal/cmd/agents.go`: Implementation of the `agents` command.
- `prompts/system.md`, `prompts/surveyor.md`, `prompts/analyst.md`, `prompts/architect.md`, `prompts/executor.md`, `prompts/integrator.md`: Source of skill content.
- `.agents/skills/cdd-surveyor/SKILL.md` and other skill directories: Directory-based installations.
- `.cursorrules`: Cursor rules file generated for Cursor editor integration.
- Tests: `prompts/integration_test.go`, `prompts/surveyor_test.go`, `agents_install_test.go`, `agents_cursor_test.go`, `agents_target_cursor_test.go`, `agents_antigravity_e2e_test.go`, `agents_install_antigravity_test.go`, `agents_antigravity_discovery_test.go`, `agents_all_flag_test.go`, `agents_all_platforms_test.go`, `agents_no_default_target_test.go`, `agents_target_validation_test.go`, `agents_all_platforms_integration_test.go`
- Google Antigravity Documentation: https://antigravity.google/docs/skills

## 4. Breaking Changes
- **v2.0**: The default `--target agent` has been removed. Users must now explicitly specify a `--target` or use `--all`. This is a backward-incompatible change that requires migration of existing scripts and documentation.
