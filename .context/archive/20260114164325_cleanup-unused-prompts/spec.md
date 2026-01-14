# Track: cleanup-unused-prompts

## 1. User Intent
Perform a clean up task to remove unused prompts and update the codebase to reflect these changes, ensuring all tests pass.

## 2. Relevant Context
- `prompts/bootstrap.md`: Unused prompt.
- `prompts/calibration.md`: Unused prompt.
- `prompts/migration.md`: Unused prompt.
- `prompts/prompts.go`: Contains embedded variables for prompts.
- `prompts/system_test.go`: Contains tests for the system prompt that are currently failing.
- `prompts/integration_test.go`: Uses some of the unused prompts.

## 3. Requirements (EARS)

- The system shall remove `prompts/bootstrap.md`, `prompts/calibration.md`, and `prompts/migration.md`.
- The system shall remove `Bootstrap` and `Calibration` variables from `prompts/prompts.go`.
- The system shall update `prompts/system_test.go` to match the current lean system prompt.
- The system shall update `prompts/integration_test.go` to remove references to deleted prompts.
- The system shall ensure `go test ./prompts/...` passes after cleanup.