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
