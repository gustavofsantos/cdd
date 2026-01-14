
---
###### Archived at: 2026-01-14 12:16:09 | Track: agent-skill-installation

# Track: agent-skill-installation

## 1. User Intent
Install the CDD System Prompt as an Agent Skill in the user repository. This allows AI agents to discover and follow the Context-Driven Development protocol automatically.

## 2. Relevant Context
- `internal/cmd/prompts.go`: Logic for `cdd prompts` command.
- `prompts/system.md`: Source content for the CDD System Prompt.
- `prompts/prompts.go`: Go source where prompts are embedded.
- `.agent/skills/cdd/SKILL.md`: The output file following the [Agent Skill specification](https://agentskills.io/specification).

## 4. Scenarios

### Scenario: Successful Skill Installation
- **Given** the user has a repository initialized with CDD.
- **When** the user runs `cdd prompts --install`.
- **Then** the command should create the directory `.agent/skills/cdd/`.
- **And** it should create the file `.agent/skills/cdd/SKILL.md`.
- **And** `SKILL.md` should have the following frontmatter:
  ```yaml
  ---
  name: cdd
  description: Protocol for implementing software features using the Context-Driven Development methodology.
  ---
  ```
- **And** the rest of `SKILL.md` should contain the content of `prompts/system.md`.
- **And** the command should print "Skill 'cdd' installed at .agent/skills/cdd/SKILL.md".

### Scenario: Installation without arguments shows help
- **Given** the user runs `cdd prompts`.
- **When** no flags are provided.
- **Then** the command should output the help message (existing behavior).

### Scenario: Prompts are self-contained
- **Given** the user views the CDD System Prompt.
- **When** the content is rendered.
- **Then** it must not contain any reference to external files like `AGENTS.local.md` or `GEMINI.md`.
- **And** it should emphasize Agent Skills as the standard for extensions.

