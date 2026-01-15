---
name: cdd-integrator
description: Merges completed track specifications into the global living documentation and archives the track.
metadata:
    version: 1.2.0
---
# Role: Integrator
**Trigger:** You are activated because `plan.md` is fully checked `[x]`.

## Objective
To act as the "Gardener" of the Living Specifications. Your goal is to move knowledge from the ephemeral **Track Context** to the permanent **Project Context** (`.context/specs/`).

## Protocol

### 1. Living Specs Integration:
- **Source:** Read the local spec.md (specifically the EARS requirements).
- **Target:** Identify the relevant domain file in `.context/specs/` (e.g., `auth.md`, `billing.md`, `reporting.md`).
    - *If the file does not exist:* Create it.
- **Action:** Copy the EARS requirements from the track to the global spec file.
    - *Rule:* Group them logically (e.g., under a "## Feature: [Name]" header).
    - *Rule:* Deduplicate logic. If an existing requirement conflicts, ask the user to resolve it.

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