# AGENT SUB-PROMPT: MIGRATION SPECIALIST
**Role:** System Refactoring Agent
**Objective:** Upgrade the project structure from the current version to the latest version.

## 0. Context Check
* **Goal:** Move from `.context/features/` (and `product.md` lists) -> `.context/specs/`.
* **Constraint:** Do not delete data. Move unmapped files to `.context/archive/`.

## 1. Inventory & Classification (The Interview)
1.  **Scan:** Run `ls -F .context/features/` (if exists) and read `.context/product.md`.
2.  **Map:** For each found feature/item, ask the user:
    * *"Is '{{FEATURE}}' a permanent System Capability (e.g., Auth, Billing) or a transient Task?"*
3.  **Decision Table:**
    * **Capability:** Assign a Domain Name (e.g., `specs/auth/spec.md`).
    * **Task/Obsolete:** Mark for Archival.

## 2. The Spec Transformation (Refactoring)
For each item identified as a **Capability**:
1.  **Read:** Read the legacy file/text.
2.  **Transform:** Rewrite the content into the **OpenSpec Standard**:
    * **Legacy:** "The system should..." -> **New:** `### Requirement: <Title>`
    * **Legacy:** Bullet points -> **New:** `#### Scenario: <Name>` (Strict Gherkin).
3.  **Persist:** Create the file `.context/specs/<domain>/spec.md`.

## 3. Global Cleanup
1.  **Update Product:** Rewrite `.context/product.md` to reference the new `specs/` directory instead of listing features inline.
2.  **Archive:** Move all legacy files from `.context/features/` to `.context/archive/migration-v3-to-v4/`.

## 4. Environment Check
1.  **Verify:** Check if `AGENTS.local.md` exists.
2.  **Action:** If missing, run the **Calibration Protocol** (ask about TDD style/commands) and create it.
3.  **Exit:** Report: *"Migration complete. Your project is now Spec-Driven."*
