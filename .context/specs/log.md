# Log Command Specification

## 1. Overview
The `log` command records permanent decisions or errors to the `decisions.md` file of a track. It supports standard arguments for quick logging and stdin for extensive multiline messages.

## 2. Requirements

### 2.1 Stdin Support
- `cdd log` shall accept the message content from standard input (stdin) if provided.
- This allows for multiline messages and piping from other commands.

### 2.2 Track Inference
- If the track name is omitted AND there is exactly one active track in `.context/tracks/`, `cdd` shall infer the track name automatically.
- If multiple tracks are active and no track name is provided, `cdd log` must return an error asking the user to specify the track.

### 2.3 Argument Modes
- **Explicit Track & Message**: `cdd log <track> <message>` (Traditional usage).
- **Explicit Track & Stdin**: `cdd log <track>` (Message read from stdin).
- **Implicit Track & Stdin**: `cdd log` (Track inferred from context, Message read from stdin).

## 3. Relevant Files
- `internal/cmd/log.go`
- `.context/tracks/<track>/decisions.md`
