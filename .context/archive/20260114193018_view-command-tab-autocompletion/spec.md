# Track: view-command-tab-autocompletion

## 1. User Intent
Implement tab autocompletion for the `view` command that autocompletes with the current task if only one task is active, or presents a selection menu if multiple tasks are active.

## 2. Relevant Context
- `cmd/view.go` - The view command implementation
- Task state and active task tracking in codebase

## 3. Requirements (EARS)

- **Ubiquitous**: The `view` command shall support tab autocompletion for task arguments.
- **Event-driven**: When the user presses TAB while typing a task argument for the `view` command, the system shall trigger autocompletion logic.
- **State-driven**: While exactly one task is active, the system shall automatically complete the task argument with that task's ID or name.
- **State-driven**: While multiple tasks are active, the system shall present a selection menu listing all active tasks for the user to choose from.
- **Unwanted**: If no tasks are active, then the system shall not autocomplete or display a selection menu.