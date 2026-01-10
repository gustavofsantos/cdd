# Plan: Update Internal Prompts to Lean Version

## Tasks

- [x] **Test System Prompt Content**
    - [x] Create a test `prompts/system_test.go` to verify `prompts.System` contains the new "Lean" protocol strings (e.g., "Tests are Truth", "Spec Cleanup").
    - [x] Run the test to confirm it fails (Red Phase).

- [x] **Update System Prompt**
    - [x] Update `prompts/system.md` with the content from `GEMINI.md`.
    - [x] Run the test to confirm it passes (Green Phase).

- [x] **Review Other Prompts**
    - [x] Check if `prompts/bootstrap.md` or `prompts/inbox.md` need minor tweaks to align with the "Lean" language (though likely `system.md` is the main one).

- [x] **Protocol Refinement: Command Usage & Local Overrides**
    - [x] Update `prompts/system.md` to use the `cdd` command instead of `go run cmd/cdd/main.go`.
    - [x] Add instructions to `prompts/system.md` regarding the `AGENTS.local.md` file (agent must look for it and allow it to override protocol/commands).
    - [x] Update `prompts/system_test.go` to verify these new requirements.

- [x] **Consolidation**
    - [x] Update `spec.md` to replace scenarios with test references.
    - [x] Archive the track.