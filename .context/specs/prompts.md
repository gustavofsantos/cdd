# Prompts Specification

## 1. Overview
The `prompts` command provides access to the core CDD instruction sets. It allows users to view the system prompts and, more importantly, install the CDD protocol as an Agent Skill for automatic discovery by AI tools.

## 2. Requirements

### 2.1 Viewing Prompts
- Running `cdd prompts` without flags should display the list of available prompt categories or general help for the command.
- (Existing behavior should be documented here if known, but for now, I'll focus on the new/specified behavior).

### 2.2 Skill Installation
- The command `cdd prompts --install` must install the CDD System Prompt as an Agent Skill.
- **Location**: It creates a directory named `.agent/skills/cdd/`.
- **Artifact**: It creates a file named `SKILL.md` inside that directory.
- **Metadata**: The `SKILL.md` file must include the following YAML frontmatter:
  ```yaml
  ---
  name: cdd
  description: Protocol for implementing software features using the Context-Driven Development methodology.
  ---
  ```
- **Content**: The body of `SKILL.md` must contain the full content of the CDD System Prompt (sourced from `prompts/system.md`).

### 2.3 Prompt Guidelines
- All CDD prompts must be self-contained.
- They must not contain references to external files that might not exist in every environment (e.g., `AGENTS.local.md`, `GEMINI.md`).
- They should focus on Agent Skills as the primary mechanism for extending the engine's capabilities.

## 3. Relevant Context
- `internal/cmd/prompts.go`: Implementation of the `prompts` command.
- `prompts/system.md`: Source of the CDD protocol instructions.
- `.agent/skills/cdd/SKILL.md`: The installed skill definition.
