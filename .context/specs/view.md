# View Command Specification

## 1. Overview
The `view` command provides a dashboard for monitoring active development tracks, viewing the staging inbox, and exploring archived work. It supports both human-readable formatted output and raw machine-readable output for shell integration.

## 2. Requirements

### 2.1 Dashboard Behavior
- When executed without arguments, `cdd view` shall display a list of current active tracks found in `.context/tracks/`.
- When the `--inbox` flag is provided, `cdd view` shall display the contents of `.context/inbox.md`.
- When the `--archived` flag is provided, `cdd view` shall display a list of archived tracks found in `.context/archive/`.
- When listing archived tracks, the system shall remove the timestamp prefix (e.g., `20260113183132_`) from the displayed names.

### 2.2 Detailed Track View
- When a `<track-name>` argument is provided, `cdd view` shall display the "Next Tasks" from that track's `plan.md`.
- When a `<track-name>` argument and the `--spec` flag are provided, `cdd view` shall display the content of that track's `spec.md`.
- Where the `--archived` flag is used with a `<track-name>`, `cdd view` shall lookup the track in `.context/archive/` instead of `.context/tracks/`.
- When a `<track-name>` argument and the `--plan` flag are provided, `cdd view` shall display the content of that track's `plan.md`.

### 2.3 Output Modes
- When the output is not a TTY (terminal), the system shall default to raw output mode.
- When the `--raw` flag is provided, the system shall use raw output mode regardless of TTY status.
- While in raw output mode, `cdd view` shall output only the requested data (e.g., track names) one per line, without headers, emojis, or Markdown formatting.

## 3. Relevant Files
- `internal/cmd/view.go`
- `.context/inbox.md`
- `.context/tracks/`
- `.context/archive/`
