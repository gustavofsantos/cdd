# Spec: Integrate Calibration Prompt

## User Intent
Integrate the new `calibration` prompt into the application, making it accessible via the `init` command flag `--calibration-prompt`.

## Relevant Context
- `prompts/calibration.md`
- `prompts/prompts.go`
- `internal/cmd/init.go`
- `README.md`

## Context Analysis
The application has a pattern for embedding prompts in the `prompts` package and exposing them through the `init` command. A new prompt `calibration.md` was added and needs to be wired into the system.

## Test Reference
- `prompts/integration_test.go`: Verifies embedding of `prompts.Calibration`.
- `internal/cmd/init_test.go`: Verifies CLI flag `--calibration-prompt`.
