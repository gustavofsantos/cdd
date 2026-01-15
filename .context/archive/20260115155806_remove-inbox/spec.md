# Track: remove-inbox

## 1. User Intent
Remove the need for the `inbox.md` file and the interaction with it when archiving tracks. This includes removing the `inbox` functionality from the `view` command.

## 2. Relevant Context
- `internal/cmd/archive.go`
- `internal/cmd/archive_test.go`
- `internal/cmd/view.go`
- `internal/cmd/view_test.go`
- `internal/cmd/util.go`
- `internal/cmd/util_test.go`

## 3. Requirements (EARS)
- **Ubiquitous:** The `archive` command shall **not** write to or create `.context/inbox.md` when archiving a track.
- **Ubiquitous:** The `view` command shall **not** support the `--inbox` (or `-i`) flag.
- **Ubiquitous:** The `view` command shall **not** read or display content from `.context/inbox.md`.
- **Ubiquitous:** The system shall **not** fail if `.context/inbox.md` is missing.
