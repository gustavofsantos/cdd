# Agents Specification

## 1. Overview
The `agents` command manages the integration between the CDD CLI and AI agents. Its primary responsibility is to install and maintain the Agent Skill that defines the CDD protocol for AI orchestration.

## 2. Requirements

### 2.1 Skill Installation
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

### 2.2 Versioning and Migration
- The command should track the version of the installed skill.
- If a newer version of the skill is available in the binary, it should offer to migrate/update the existing one.
- If an update occurs, the previous version should be backed up with a `.bak` extension.

## 3. Relevant Context
- `internal/cmd/agents.go`: Implementation of the `agents` command.
- `prompts/system.md`: Source of the CDD protocol instructions.
- `.agent/skills/cdd/SKILL.md`: The installed skill definition.
