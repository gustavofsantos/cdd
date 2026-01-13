# Spec: Integrate Executor and Planner Prompts

## User Intent
Integrate the `executor` and `planner` prompts into the application, making them accessible via the `init` command flags.

## Relevant Context
- `prompts/executor.md`
- `prompts/planner.md`
- `prompts/prompts.go`
- `internal/cmd/init.go`

## Context Analysis
The application currently has a mechanism to bundle prompts into the binary using `go:embed` in the `prompts` package. These prompts are then exposed via flags in the `init` command. Two new prompts, `executor.md` and `planner.md`, have been added to the `prompts/` directory but are not yet integrated into the Go code or the CLI.

## Scenarios

### Scenario 1: Prompts are embedded in the Go package
- **Given** files `prompts/executor.md` and `prompts/planner.md` exist.
- **When** the `prompts` package is compiled.
- **Then** `prompts.Executor` and `prompts.Planner` should contain the contents of these files.

### Scenario 2: CLI provides flags for the new prompts
- **Given** the application is compiled.
- **When** I run `cdd init --executor-prompt`.
- **Then** it should output the content of `executor.md`.
- **When** I run `cdd init --planner-prompt`.
- **Then** it should output the content of `planner.md`.
