# Spec: Refactor Prompt Printing to Dedicated Command

## User Intent
Move the prompt printing logic from flags in the `init` command to a separate `prompts` command for better organization and clarity.

## Relevant Context
- `internal/cmd/init.go`
- `internal/cmd/prompts.go` (to be created)
- `internal/cmd/init_test.go`
- `README.md`

## Context Analysis
Currently, `cdd init` overloaded with flags to print various prompts. This is less intuitive than having a dedicated `cdd prompts` command. We need to create the new command, migrate the flags and logic, and then remove the old flags from `init`.

## Scenarios

### Scenario 1: New `prompts` command is available
- **Given** the application is compiled.
- **When** I run `cdd prompts --system`.
- **Then** it should output the CDD System Prompt.
- **When** I run `cdd prompts --executor`.
- **Then** it should output the Executor Prompt.
- (And similarly for other prompts: bootstrap, inbox, planner, calibration)

### Scenario 2: `init` command no longer has prompt flags
- **Given** the application is compiled.
- **When** I run `cdd init --system-prompt`.
- **Then** it should fail with an "unknown flag" error.
- **When** I run `cdd init`.
- **Then** it should still initialize the environment correctly.
