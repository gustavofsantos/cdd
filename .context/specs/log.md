# Log Command Specification

## 1. Overview
The `log` command records permanent decisions or errors to the `decisions.md` file of a track. It supports standard arguments for quick logging and stdin for extensive multiline messages.

## 2. Requirements

### 2.1 Stdin Support
- When message content is provided via standard input (stdin), `cdd log` shall accept and use it as the log entry.

### 2.2 Track Inference
- When the track name is omitted and there is exactly one active track in `.context/tracks/`, the system shall infer the track name automatically.
- When multiple tracks are active and no track name is provided, the system shall return an error asking the user to specify the track.

### 2.3 Argument Modes
- When invoked as `cdd log <track> <message>`, the system shall log the provided message to the specified track.
- When invoked as `cdd log <track>` with content in stdin, the system shall log the stdin content to the specified track.
- When invoked as `cdd log` with content in stdin, the system shall log the stdin content to the inferred active track.
