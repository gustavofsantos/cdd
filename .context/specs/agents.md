# Agents Command Specification

## 1. Overview
The `agents` command manages the integration between the CDD CLI and AI agents. Its primary responsibility is to install and maintain the Agent Skill that defines the CDD protocol for AI orchestration.

## 2. Requirements

### 2.1 Skill Installation (Directory-Based Targets)
- When the command `cdd agents --install --target <target>` is executed, the system shall install the CDD System Prompt as an Agent Skill.
- When the `--target` flag is missing and `--all` is not specified, the system shall return an error message with examples of correct usage.
- When a valid target is provided, the system shall create a directory named `.agent/skills/cdd/` (or `.claude/skills/`, `.agents/skills/` depending on target).
- The system shall create a file named `SKILL.md` inside the target directory.
- The `SKILL.md` file shall include the required YAML frontmatter (name, version, description) and the full content of the CDD System Prompt.
- The system shall support `agent`, `agents`, and `claude` as directory-based targets.

### 2.2 Antigravity Target
- When the command `cdd agents --install --target antigravity` is executed, the system shall install all six CDD skills to `.agent/skills/` in Antigravity-compatible format.
- The system shall create `.agent/skills/{skill-id}/SKILL.md` for each of the six skills (`cdd`, `cdd-surveyor`, `cdd-analyst`, `cdd-architect`, `cdd-executor`, `cdd-integrator`).
- The system shall validate that each skill contains the required YAML frontmatter fields (`name`, `description`).
- If a skill with the same version already exists, the system shall skip the installation (idempotent behavior).

### 2.3 Cursor Rules Installation
- When the command `cdd agents --install --target cursor` is executed, the system shall generate a `.cursorrules` file in the project root.
- The `.cursorrules` file shall concatenate all six CDD agent skills with clear section separators.
- The file shall include YAML frontmatter containing version metadata at the top.
- If a newer version is being installed, the system shall backup the existing file to `.cursorrules.bak`.

### 2.4 Multi-Platform Installation
- When the command `cdd agents --install --all` is executed, the system shall install all six CDD skills for all supported platforms in a single invocation.
- The system shall install skills for all directory-based targets as well as special targets (`cursor`, `antigravity`).
- This flag shall override any specific `--target` provided.

### 2.5 Versioning and Migration
- The system shall track the version of the installed skill.
- When a newer version of the skill is available in the binary, the system shall offer to migrate/update the existing one.
- When an update occurs, the system shall backup the previous version with a `.bak` extension.
