# Specification: Update Internal Prompts to Lean Version

## User Intent
The user wants to update the internal prompts defined in the `prompts` directory to implement the `lean` version of the CDD protocol, use `cdd` as the default command, and support `AGENTS.local.md` for local overrides.

## Relevant Context
*   `prompts/system.md`: The file containing the system prompt that needs updating.
*   `prompts/prompts.go`: The Go file embedding the prompts.
*   `GEMINI.md`: The reference for the new "Lean" protocol.

## Context Analysis
The `prompts/system.md` file has been updated to:
1.  **Philosophy:** "Tests are Truth", `spec.md` is temporary.
2.  **Phase 1 (Red Phase):** Treat tests as documentation.
3.  **Phase 2 (Consolidation):** "Spec Cleanup" step.
4.  **Commands:** Use `cdd` directly.
5.  **Local Overrides:** Respect `AGENTS.local.md`.

## Test Reference
*   `prompts/system_test.go`