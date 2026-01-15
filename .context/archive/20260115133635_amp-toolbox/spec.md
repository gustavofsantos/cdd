# Track: amp-toolbox

## 1. User Intent
Make the entire `cdd` CLI compatible with Amp's toolbox system by creating executable toolbox files that Amp can discover via the `AMP_TOOLBOX` environment variable.

## 2. Relevant Context
- **Amp Toolbox System**: Executables in `AMP_TOOLBOX` directory are discovered at Amp startup, invoked with `TOOLBOX_ACTION=describe` to return JSON tool metadata, then invoked with `TOOLBOX_ACTION=execute` with arguments on stdin
- **cdd CLI commands**: init, start, recite, log, archive, view, agents, delete, version
- **Tool format**: Each executable must output JSON on describe with fields: `name`, `description`, `args` (object with arg names and types)
- **Current cdd structure**: Go CLI using Cobra framework with commands in `internal/cmd/`
- **Goal**: Each cdd subcommand becomes an Amp toolbox executable

## 3. Requirements (EARS)

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