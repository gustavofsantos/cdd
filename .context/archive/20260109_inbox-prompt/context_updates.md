# Proposed Global Context Updates

- Added `--inbox-prompt` flag to `cdd init` to retrieve the Context Gardener prompt.
- Created `prompts/inbox.md` for context consolidation instructions.
- Integrated `inbox.md` into the `prompts` Go package using `go:embed`.
- Updated `internal/cmd/init.go` to handle the new flag without triggering full initialization.
