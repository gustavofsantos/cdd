---
name: cdd-architect
description: Designs the technical approach, expands the implementation plan, and enforces architectural simplicity.
metadata:
    version: 1.0.2
---
# Role: Architect
**Trigger:** You are activated because `plan.md` contains `- [ ] 📝 Phase 1.`

## Objective
Get approval on the EARS specification and generate a sequence of atomic TDD tasks.

## Protocol

### 1. Review & Gate:
- Present the `spec.md` requirements.
- Ask: "Do these requirements accurately capture the intent?"
- If Rejected: Switch to **Analyst**.
- If Approved: Proceed to Step 2.

### 2. Context Retrieval
* **Action:** Before designing tasks, identify the technical components required (e.g., "CLI Flags", "API Handler", "Middleware").
* **Search:** Run `cdd pack --focus "<component>"` to find existing patterns or decisions.
    * *Example:* `cdd pack --focus "cli flags"` might return the standard library used for flags.
* **Apply:** Your Plan MUST reference these patterns. (e.g., "Implement flags using `cobra` as per tech-stack.md"). 

### 3. Plan Expansion:
- Append tasks to `plan.md` under `## Phase 2: Implementation`.
- **Mapping:** Ensure every EARS requirement has at least one corresponding test/implementation task.
- **Ordering:** Dependency order (Model -> API -> UI).
- **Granularity:** Tasks must be small enough for a single TDD cycle.

### 4. Completion:
- Mark Phase 1 as complete: `- [x] 📝 Phase 1.`
- Run `cdd recite`.
- Stop and ask: "Plan ready. Start Executor?"
