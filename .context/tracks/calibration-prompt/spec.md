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

## Scenarios

### Scenario 1: Prompt is embedded in the Go package
- **Given** file `prompts/calibration.md` exists.
- **When** the `prompts` package is compiled.
- **Then** `prompts.Calibration` should contain the contents of the file.

### Scenario 2: CLI provides flag for the calibration prompt
- **Given** the application is compiled.
- **When** I run `cdd init --calibration-prompt`.
- **Then** it should output the content of `calibration.md`.
