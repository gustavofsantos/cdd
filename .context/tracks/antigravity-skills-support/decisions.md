# Implementation Journal
> Created Wed Jan 14 18:58:13 -03 2026

## Architecture Overview

### Approach
Extend the existing `agents` command to support Antigravity as a new installation target, following the established pattern used for Claude and Cursor. This keeps implementation minimal and leverages existing infrastructure.

### Key Decisions

1. **Reuse Existing Installation Pattern**
   - Use the directory-based `installSkill` approach (like Claude/Agents) rather than monolithic file approach (like Cursor).
   - Target directory: `.agent/skills/` per Antigravity specification.
   - Rationale: Antigravity expects individual skill folders; directory approach aligns with spec.

2. **YAML Frontmatter Handling**
   - CDD prompt files already contain YAML frontmatter with `name`, `description`, and `metadata.version`.
   - Reuse extraction logic (`extractVersion`) and frontmatter parsing without modification.
   - Rationale: Reduces code complexity; existing prompts already conform to Antigravity requirements.

3. **No Additional Formatter Needed**
   - Skills are written as-is; they already comply with Antigravity's SKILL.md format.
   - Validation ensures required fields exist before writing.
   - Rationale: Avoids unnecessary abstraction; prompts are already well-structured.

4. **Version-Based Update Detection**
   - Same versioning strategy as Claude/Cursor: compare installed vs. new version.
   - Only update if versions differ; skip if already installed.
   - Rationale: Prevents unnecessary rewrites; aligns with existing upgrade behavior.

5. **Minimal Code Changes**
   - Add "antigravity" case to target switch statement.
   - Create `installAntigravitySkill` function (similar to `installSkill` but targeting `.agent/skills/`).
   - Rationale: Follows established patterns; minimal diff; easy to maintain.