# Spec: Support Log Stdin

## Context
The `cdd log` command currently requires two arguments: `<track-name>` and `<message>`.
Users (and AI agents) want to pipe content into the log command, especially for multiline messages or when scripting.
Also, if there is only one active track, the user shouldn't need to specify the track name.

## Requirements
1.  **Stdin Support**: `cdd log` should accept the message from stdin if provided.
2.  **Track Inference**: If the track name is omitted AND there is exactly one active track, `cdd` should infer the track name.
3.  **Argument Flexibility**:
    *   `cdd log <track>` (message from stdin)
    *   `cdd log` (track extracted from context if single, message from stdin)
    *   `cdd log <track> <message>` (existing behavior)

## PLAN
- [x] Modify `internal/cmd/log.go` to handle 0-2 args.
- [x] Implement Stdin reading if args < required.
- [x] Implement active track discovery for 0-arg case.
- [ ] Fix Tests (MockFileSystem requires dummy files to list directories).
