# Track: agent-skill-migration

## 1. User Intent
> Implement a migration capability for Agent Skills to establish a clear versioning strategy for `SKILL.md` files. The system should recognize existing installations (including the current unversioned "initial" state), detecting when the tool's built-in prompt version is newer than the installed one. Upon detection, it should safe-guard the user's existing prompt (e.g., backup) and upgrade to the new version, enabling rapid iteration of the agent's capabilities.

## 2. Relevant Context
- `internal/cmd/prompts.go`: Core logic for `cdd prompts --install` and file generation.
- `.agent/skills/cdd/SKILL.md`: The installed artifact that needs version tracking.
- `prompts/prompts.go`: Source of the embedded system prompt string.

## 4. Scenarios
 Feature: Agent Skill Migration

   Scenario: Fresh Installation
     Given no `SKILL.md` exists in `.agent/skills/cdd/`
     When I run `cdd prompts --install`
     Then the latest `SKILL.md` is created with the current version in frontmatter
     And the output confirms "Skill 'cdd' installed (vX)"

   Scenario: Up-to-date Installation
     Given `SKILL.md` exists with `version: 2`
     And the internal tool version is `2`
     When I run `cdd prompts --install`
     Then no changes are made to the file
     And the output confirms "Skill 'cdd' is up to date (v2)"

   Scenario: Legacy Migration (v0/v1 to v2)
     Given `SKILL.md` exists without a `version` field (Legacy/Initial)
     And the internal tool version is `2`
     When I run `cdd prompts --install`
     Then the existing `SKILL.md` is renamed to `SKILL.md.bak` (or similar backup)
     And the new `SKILL.md` is created with `version: 2`
     And the output confirms "Migrated legacy skill to v2. Backup saved."
