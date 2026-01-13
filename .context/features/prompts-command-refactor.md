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

## Test Reference
- `internal/cmd/prompts_test.go`: Verifies the new `prompts` command and its flags (`--system`, `--executor`, etc.).
- `internal/cmd/init_test.go`: Verifies that the `init` command still functions correctly for project initialization after refactoring.
