# Context Updates: inbox-cleanup-reminder

- Implemented `CheckInboxSize` utility in `internal/cmd/util.go`.
- Integrated the check into `cdd archive` command.
- When `.context/inbox.md` grows beyond 50 lines, the user is now proactively reminded to use `cdd init --inbox-prompt` for context consolidation.
- Added unit and integration tests for this behavior.
