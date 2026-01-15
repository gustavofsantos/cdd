# Amp Toolbox Specification

## 1. Overview
The CDD CLI is fully compatible with Amp's toolbox system. All subcommands are exposed as individual executable wrappers that follow Amp's tool protocol (describe/execute actions), allowing Amp to discover and execute them via the `AMP_TOOLBOX` environment variable.

## 2. Requirements

### 2.1 Describe Action
- When a tool is invoked with `TOOLBOX_ACTION=describe`, the system shall output valid JSON describing the tool.
- The JSON output shall include `name` (string), `description` (string), and `args` (object).
- Each argument in the `args` object shall be defined as a tuple of `[type, help_text]`.
- The system shall not write any output to stderr during the describe action.

### 2.2 Execute Action
- When a tool is invoked with `TOOLBOX_ACTION=execute`, the system shall read JSON arguments from stdin.
- The system shall execute the corresponding `cdd` command with the provided arguments.
- The system shall inherit stdio from the parent process to pipe output and error directly to Amp.
- If an optional argument is missing, the system shall use sensible defaults (e.g., empty string, empty array, false).

### 2.3 Command Mapping
- The system shall expose each `cdd` subcommand as a separate executable tool (e.g., `cdd-init`, `cdd-start`, `cdd-recite`, `cdd-log`, `cdd-archive`, `cdd-view`, `cdd-agents`, `cdd-delete`, `cdd-version`).
- The system shall preserve the semantic behavior of each original command.

### 2.4 Toolbox Directory
- When a toolbox directory is specified during build or installation, the system shall place all executables in that directory.
- The tools shall be discoverable by Amp when the `AMP_TOOLBOX` environment variable points to that directory.

### 2.5 Error Handling
- If a required argument is missing, the system shall exit with a non-zero code.
- If the underlying `cdd` subcommand fails, the system shall propagate the exit code.
- The system shall write error messages to stderr for visibility in Amp.
