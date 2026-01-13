# Context Inbox

This file contains ephemeral updates from closed tracks. 
The Context Gardener promotes these updates to global files periodically.

---


## Updates from Track: update-prompts-lean (Sat Jan 10 09:16:49 -03 2026)
Updated `prompts/system.md` to the new "Lean" CDD protocol (Tests are Truth, Spec Cleanup).
Changed default `cdd` commands in the system prompt to use `cdd` instead of `go run cmd/cdd/main.go`.
Added a requirement for the agent to check for and respect `AGENTS.local.md` for local configuration overrides.
Added/Updated `prompts/system_test.go` to validate these requirements.


## Updates from Track: time-tracking (Tue Jan 13 10:44:43 -03 2026)
# Proposed Global Context Updates
> Add notes here if product.md or tech-stack.md needs updating.

- Added Time Tracking to `cdd`:
    - `start` now creates `metadata.json` with a `started_at` timestamp.
    - `archive` now calculates and prints the duration of the track.
    - Uses `encoding/json` and `time` packages.


## Updates from Track: prompt-integration (Tue Jan 13 10:55:00 -03 2026)
- Integrated `executor.md` and `planner.md` prompts into the `prompts` package using `go:embed`.
- Added `--executor-prompt` and `--planner-prompt` flags to the `cdd init` command.
- Updated `internal/cmd/init.go` to use `cmd.Println` instead of `fmt.Println` to support testing output.
- Synchronized `prompts/system.md` with `GEMINI.md` to fix failing tests in the `prompts` package.
- Updated `README.md` to document the new prompt flags.


## Updates from Track: calibration-prompt (Tue Jan 13 11:07:16 -03 2026)
- Integrated the new `calibration.md` prompt into the `prompts` package.
- Added `--calibration-prompt` flag to the `cdd init` command.
- Updated `README.md` to document the new flag.


## Updates from Track: prompts-command-refactor (Tue Jan 13 11:12:18 -03 2026)
- Refactored prompt printing: moved logic from flags in the `init` command to a new dedicated `prompts` command.
- The `prompts` command supports flags: `--system`, `--bootstrap`, `--inbox`, `--executor`, `--planner`, and `--calibration`.
- Validated that `init` still correctly initializes the `.context` directory structure without prompt flags.
- Updated `README.md` to reflect the new command usage.
