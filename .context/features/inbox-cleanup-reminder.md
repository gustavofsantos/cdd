# Spec - Inbox Cleanup Reminder

Implement a proactive check that monitors the size of `.context/inbox.md` and suggests consolidation when it grows too large.

## User Intent
"create a new track to implement a feature that after every change in the inbox.md file done through the cdd command, it should check if the file has more than 50 lines and then suggest the user to get the cleaner prompt and run it with its agent"

## Relevant Context
- `internal/cmd/archive.go`: Currently the main command that modifies `inbox.md`.
- `internal/cmd/init.go`: Contains the `--inbox-prompt` logic.
- `.context/inbox.md`: The file to monitor.

## Context Analysis
- The `inbox.md` file accumulates updates from archived tracks.
- As it grows, it becomes harder for AI agents to process efficiently.
- The "cleaner prompt" is the Context Gardener prompt available via `cdd init --inbox-prompt`.

## Scenarios

### Scenario 1: Inbox exceeds 50 lines after archive
**GIVEN** an `inbox.md` file with 48 lines
**WHEN** I run `cdd archive <track>` and it appends 5 lines to `inbox.md`
**THEN** the command should succeed
**AND** it should display a suggestion: "⚠️ Your .context/inbox.md has 53 lines. It's getting large! Run 'cdd init --inbox-prompt' to get the Context Gardener prompt and consolidate your context."

### Scenario 2: Inbox remains under 50 lines after archive
**GIVEN** an `inbox.md` file with 10 lines
**WHEN** I run `cdd archive <track>` and it appends 5 lines to `inbox.md`
**THEN** the command should succeed
**AND** it should NOT display the suggestion.
