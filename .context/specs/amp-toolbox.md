# Feature: Amp Toolbox Integration

## User Intent
Make the entire `cdd` CLI compatible with Amp's toolbox system by creating executable toolbox files that Amp can discover via the `AMP_TOOLBOX` environment variable.

## Overview
The CDD CLI is now fully compatible with Amp's toolbox system. All subcommands are exposed as individual executable wrappers that follow Amp's tool protocol (describe/execute actions).

## Requirements (EARS)

### Describe Action
- The system shall output valid JSON when invoked with `TOOLBOX_ACTION=describe`
- The JSON shall include required fields: `name` (string), `description` (string), `args` (object)
- Each arg in the object shall be a tuple of `[type, help_text]`
- The system shall not write any output to stderr during describe action

### Execute Action
- When the system is invoked with `TOOLBOX_ACTION=execute`, it shall read JSON arguments from stdin
- The system shall execute the corresponding cdd command with the provided arguments
- The system shall inherit stdio from the parent process (pipes output/error directly to Amp)
- If an argument is missing but optional, the system shall use sensible defaults (e.g., empty string, empty array, false)

### Command Mapping
- The system shall expose each cdd subcommand as a separate executable tool
- The system shall preserve the command's semantic behavior (no changes to what the command does)
- Commands exposed: `cdd-init`, `cdd-start`, `cdd-recite`, `cdd-log`, `cdd-archive`, `cdd-view`, `cdd-agents`, `cdd-delete`, `cdd-version`

### Toolbox Directory
- Where a toolbox directory is specified in the build or installation, the system shall place all executables there
- The system shall be discoverable by Amp when `AMP_TOOLBOX` is set to that directory

### Error Handling
- If a required argument is missing, the system shall exit with a non-zero code
- If the cdd subcommand fails, the system shall propagate the exit code
- Error messages shall be written to stderr for visibility in Amp

## Toolbox Tools

The following tools are available when `AMP_TOOLBOX` is configured:

| Tool | Description | Arguments |
|------|-------------|-----------|
| `cdd-init` | Bootstrap the CDD environment | None |
| `cdd-start` | Create an isolated workspace (Track) | `track_name` (required) |
| `cdd-recite` | Output the current Plan to the context window | `track_name` (optional) |
| `cdd-log` | Record a decision or event in the decision log | `track_name` (required), `message` (required) |
| `cdd-archive` | Archive a completed track and move it to history | `track_name` (required) |
| `cdd-view` | Display track details or dashboard | `track_name` (optional), flags: `spec`, `plan`, `log`, `raw`, `archived`, `inbox` |
| `cdd-agents` | Manage AI agent integration and skills | `install` (bool), `target` (string), `all` (bool) |
| `cdd-delete` | Destructively remove an active track | `track_name` (required) |
| `cdd-version` | Display the version of the cdd CLI | None |

## Setup Instructions

Users can enable Amp integration by:

1. Building or downloading the toolbox binaries
2. Setting `AMP_TOOLBOX` environment variable to the toolbox directory
3. Starting Amp (tools are auto-discovered)

See [AMP_TOOLBOX.md](../AMP_TOOLBOX.md) for detailed setup instructions.

## Implementation Details

- All toolbox wrappers are built from `cmd/toolbox/*/main.go` Go source files
- Each wrapper invokes the main `cdd` binary with appropriate arguments
- Build configuration in `.goreleaser.yaml` automatically compiles all wrappers
- No cobra integration needed in wrappers (they shell out to the main binary)

## Status

✅ Complete - Ready for production use with Amp
