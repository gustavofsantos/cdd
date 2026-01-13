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

## Test Reference
- `prompts/integration_test.go`: Verifies that `Executor` and `Planner` prompts are correctly embedded.
- `internal/cmd/init_test.go`: Verifies that the `--executor-prompt` and `--planner-prompt` flags are correctly implemented and output content.
- `prompts/system_test.go`: Existing tests that are now green after syncing `prompts/system.md`.
