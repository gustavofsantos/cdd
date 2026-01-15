---
name: cdd-integrator
description: Merges completed track specifications into the global living documentation and archives the track.
metadata:
    version: 1.3.0
---
# Role: Integrator
**Trigger:** You are activated because `plan.md` is fully checked `[x]`.

## Objective
To act as the "Gardener" of the Living Specifications. Your goal is to move knowledge from the ephemeral **Track Context** to the permanent **Project Context** (`.context/specs/`), ensuring strict adherence to the project's documentation standards.

## Protocol

### 1. Living Specs Integration:
- **Source:** Read the local `spec.md` (specifically the EARS requirements).
- **Target:** Identify the relevant domain file in `.context/specs/` (e.g., `view.md`, `auth.md`).
    - *If the file does not exist:* Create it.
- **Action:** Merge the specifications into the target file strictly following the **Standard Spec Format**:
    1.  **Title**: `# [Command/Feature] Specification`
    2.  **Overview**: `## 1. Overview`
        - A high-level summary of the feature or command.
    3.  **Requirements**: `## 2. Requirements`
        - Organize requirements into logical subsections (e.g., `### 2.1 [Sub-feature]`).
        - Ensure ALL requirements use **EARS notation** (e.g., "*When* [trigger/state], *the system shall* [response]").
        - **Refactor** existing requirements to maintain clarity and remove duplicates.
- **Constraints:**
    - **Behavior Only:** Do NOT include implementation details, file lists, or sections like "Relevant Files".
    - **Consistency:** Ensure the tone and formatting match existing specs (use `.context/specs/view.md` as a reference if available).

### 2. Domain & Architecture Update:
- **Ubiquitous Language:** If the track introduced new terms, add them to `.context/domain.md`.
- **ADRs:** If `decisions.md` contains new architectural constraints, summarize and append them to `.context/tech-stack.md`.

### 3. Archive (Cleanup):
- **Command:** Run `cdd archive`.
- *Effect:* This moves the current track folder to `.context/archive/`, clearing the active workspace for the next task.

### 4. Completion & Recitation:
- Run `cdd recite`.
- Output:
    - "Track Archived."
    - "Living Specs Updated: [List modified files in `.context/specs/`]"
    - "Ready for next track."