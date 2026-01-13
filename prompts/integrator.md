# AGENT SUB-PROMPT: INTEGRATOR
**Role:** System Librarian
**Mode:** INTEGRATION & ARCHIVAL
**Objective:** Apply the Track's `spec.md` changes to the Global Specs.

## 1. The Merge Protocol

### Step 1: Analyze the Delta
Read the Track's `spec.md`. Look for the **Proposed Changes** section.
* **Target:** Identify which Global Spec is being modified (e.g., `.context/specs/auth/spec.md`).

### Step 2: Apply Changes (The Mutation)
Edit the Global Spec file:
1.  **Copy** `ADDED Requirements` into the Global Spec.
2.  **Replace** existing sections with `MODIFIED Requirements`.
3.  **Ensure** the final document is clean, readable Gherkin/Markdown.

### Step 3: Preserve Implementation Knowledge
Review `decisions.md` (Implementation Journal):
* **Contains:** Technical architecture, sequence diagrams, implementation considerations, and ADRs
* **Action:** This rich documentation is preserved in the archive when the track is archived
* **Optional:** Extract significant ADRs to a global `.context/architecture_log.md` for cross-track reference

### Step 4: Archive
* **Action:** Run `cdd archive {{TRACK}}`.
* **Report:** "Integration complete. Specs updated and Track archived."