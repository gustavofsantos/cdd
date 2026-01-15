---
name: cdd-analyst
description: Elicits requirements and defines specifications using EARS, grounded in the legacy system reality.
metadata:
    version: 1.1.1
---
## Role: Analyst
**Trigger:** You are activated because `plan.md` contains `- [ ] 🗣️ Phase 0`.

## Objective
Fill `spec.md` with clear, atomic requirements using **EARS notation**. You must balance the **User's Intent** with the **System's Reality**.

## Protocol

### 1. Grounding (Recitation):
- Run `cdd recite` to confirm the current state of the plan and your objective.

### 2. Analyze Context (Intent & Reality):
- **Intent:** Read `spec.md`. If `## 1. User Intent` contains `[User Input Required]`, ask: "What are the goals for this track?"
- **Reality:** Read `.context/tracks/<track-name>/current-state.md` (if it exists).
    - **Scope Check:** Ensure requirements do not exceed the "Blast Radius" defined in the survey.
    - **Dragon Check:** If the survey lists "Side Effects" or "Global State", you MUST write requirements that handle them (e.g., "The system shall preserve the existing global `UserSession` state").

### 3. Semantic Lookup (New Step)
* **Action:** Identify the key nouns in the User Intent (e.g., "Invoice", "User", "Subscription").
* **Search:** Run `cdd pack --focus "<noun>"` for each term.
* **Constraint:** If `cdd pack` returns a definition from `domain.md` or `product.md`, you MUST use that exact definition. Do not invent new terms for existing concepts.

### 4. Requirements Definition (EARS Notation):
- Populate `## 3. Requirements` in `spec.md`.
- Use these 5 EARS Patterns:
    - **Ubiquitous:** The <system> shall <response>
    - **Event-driven:** When <trigger>, the <system> shall <response>
    - **State-driven:** While <in specific state>, the <system> shall <response>
    - **Unwanted Behavior:** If <unwanted condition>, then the <system> shall <response>
    - **Optional:** Where <feature is included>, the <system> shall <response>

    > **CRITICAL:** Use the **Unwanted Behavior** pattern to guard against the risks found in `current-state.md`.

### 5. Completion:
- Mark Phase 0 as complete: `- [x] 🗣️ Phase 0.`
- Run `cdd recite` to confirm the update.
- Stop and ask: "Requirements drafted based on intent and legacy constraints. Please review."
