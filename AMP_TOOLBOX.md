# CDD CLI - Amp Toolbox Integration

The `cdd` CLI is compatible with [Amp's toolbox system](https://ampcode.com/manual#toolboxes), allowing Amp to discover and use all CDD commands as native tools.

## Overview

Amp can discover executables in a directory specified by the `AMP_TOOLBOX` environment variable. The CDD project provides individual executable wrappers for each subcommand that follow Amp's toolbox protocol:

- `TOOLBOX_ACTION=describe`: Output JSON metadata about the tool
- `TOOLBOX_ACTION=execute`: Execute the tool with arguments passed via stdin as JSON

## Available Tools

The following tools are available:

| Tool | Description | Arguments |
|------|-------------|-----------|
| `cdd-init` | Bootstrap the CDD environment | None |
| `cdd-start` | Create an isolated workspace (Track) | `track_name` (required) |
| `cdd-recite` | Output the current Plan to the context window | `track_name` (optional) |
| `cdd-log` | Record a decision or event in the decision log | `track_name` (required), `message` (required) |
| `cdd-archive` | Archive a completed track and move it to history | `track_name` (required) |
| `cdd-view` | Display track details or dashboard | `track_name` (optional), `spec`, `plan`, `log`, `raw`, `archived`, `inbox` (all boolean flags) |
| `cdd-agents` | Manage AI agent integration and skills | `install` (bool), `target` (string), `all` (bool) |
| `cdd-delete` | Destructively remove an active track | `track_name` (required) |
| `cdd-version` | Display the version of the cdd CLI | None |

## Setup

### Option 1: Download Pre-built Binaries (Recommended)

1. Download the latest release from [GitHub Releases](https://github.com/gustavofsantos/cdd/releases)
2. Extract the toolbox binaries to a directory (e.g., `~/.cdd/toolbox`)
3. Set the environment variable:
   ```bash
   export AMP_TOOLBOX=~/.cdd/toolbox
   ```

### Option 2: Build from Source

1. Clone the repository:
   ```bash
   git clone https://github.com/gustavofsantos/cdd.git
   cd cdd
   ```

2. Build the toolbox wrappers:
   ```bash
   go build -o ~/.cdd/toolbox/cdd-init ./cmd/toolbox/init/main.go
   go build -o ~/.cdd/toolbox/cdd-start ./cmd/toolbox/start/main.go
   go build -o ~/.cdd/toolbox/cdd-recite ./cmd/toolbox/recite/main.go
   go build -o ~/.cdd/toolbox/cdd-log ./cmd/toolbox/log/main.go
   go build -o ~/.cdd/toolbox/cdd-archive ./cmd/toolbox/archive/main.go
   go build -o ~/.cdd/toolbox/cdd-view ./cmd/toolbox/view/main.go
   go build -o ~/.cdd/toolbox/cdd-agents ./cmd/toolbox/agents/main.go
   go build -o ~/.cdd/toolbox/cdd-delete ./cmd/toolbox/delete/main.go
   go build -o ~/.cdd/toolbox/cdd-version ./cmd/toolbox/version/main.go
   ```

   Or use goreleaser to build all binaries:
   ```bash
   goreleaser build --snapshot --clean
   ```

3. Set the environment variable:
   ```bash
   export AMP_TOOLBOX=~/.cdd/toolbox
   ```

### Step 3: Persist the Environment Variable

Add the export to your shell configuration file (`~/.bashrc`, `~/.zshrc`, or `~/.config/fish/config.fish`):

```bash
export AMP_TOOLBOX=$HOME/.cdd/toolbox
```

## Using with Amp

Once set up, start Amp normally:

```bash
amp
```

Amp will automatically discover all CDD tools and make them available. You can then ask Amp to use CDD commands:

- "Create a new track called 'feature-x'" → Uses `cdd-start`
- "What's the current plan?" → Uses `cdd-recite`
- "Record this decision: we're using PostgreSQL" → Uses `cdd-log`
- "Archive the current track" → Uses `cdd-archive`

## How It Works

Each tool wrapper is a small executable that:

1. **On Startup (Describe Phase)**: When Amp starts, it invokes each executable with `TOOLBOX_ACTION=describe`, which outputs JSON metadata about the tool's name, description, and arguments.

2. **On Use (Execute Phase)**: When Amp decides to use the tool, it invokes the executable again with `TOOLBOX_ACTION=execute` and passes arguments via stdin as JSON.

3. **Command Dispatch**: The wrapper parses the arguments and invokes the main `cdd` CLI with the appropriate subcommand and flags.

## Requirements

- The `cdd` binary must be installed and available in your `$PATH`
- All toolbox wrapper binaries must be executable (`chmod +x`)
- The `cdd` CLI must be version 0.1.0 or later

## Troubleshooting

### Tools not appearing in Amp

1. Check that `AMP_TOOLBOX` is set:
   ```bash
   echo $AMP_TOOLBOX
   ```

2. Verify the binaries exist and are executable:
   ```bash
   ls -la $AMP_TOOLBOX/cdd-*
   ```

3. Test a tool manually:
   ```bash
   TOOLBOX_ACTION=describe $AMP_TOOLBOX/cdd-init
   ```

### Tools fail with "cdd binary not found"

Ensure the main `cdd` CLI is installed and in your `$PATH`:

```bash
which cdd
cdd version
```

## Development

To add a new tool:

1. Create a new wrapper in `cmd/toolbox/<command>/main.go`
2. Add the build configuration to `.goreleaser.yaml`
3. Update this documentation

For more details, see the [Contributing Guide](CONTRIBUTING.md).
