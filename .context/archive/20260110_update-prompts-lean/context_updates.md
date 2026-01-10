Updated `prompts/system.md` to the new "Lean" CDD protocol (Tests are Truth, Spec Cleanup).
Changed default `cdd` commands in the system prompt to use `cdd` instead of `go run cmd/cdd/main.go`.
Added a requirement for the agent to check for and respect `AGENTS.local.md` for local configuration overrides.
Added/Updated `prompts/system_test.go` to validate these requirements.
