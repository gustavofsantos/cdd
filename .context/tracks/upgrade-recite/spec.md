# Track: upgrade-recite

## 1. User Intent
The user wants to improve the `recite` command to be more focused. Instead of showing the entire `plan.md`, it should only show the current active section (the first one containing an unchecked item) to reduce cognitive load and focus the agent's attention.

## 2. Relevant Context
- `internal/cmd/recite.go`: Contains the logic for the `recite` command.
- `internal/cmd/recite_test.go`: Contains tests for the `recite` command.
- `.context/tracks/*/plan.md`: The files being recited.

## 3. Requirements (EARS)
- **WHEN** the `recite` command is executed,
  **THEN** it MUST identify the first unchecked item (`- [ ]`) in the `plan.md` file.
- **WHEN** the `recite` command is executed,
  **THEN** it MUST display the section (header and its content) that contains the first unchecked item.
- **WHEN** the `recite` command is executed,
  **THEN** it MUST NOT display sections that only contain checked items (`- [x]`) or no items at all, unless they are the "next" section.
- **WHEN** multiple sections have unchecked items,
  **THEN** the `recite` command SHOULD ONLY display the first such section.
- **WHEN** NO unchecked items are found in the `plan.md` file,
  **THEN** the `recite` command SHOULD display a message indicating all tasks are completed.
- **WHEN** the `recite` command is executed,
  **THEN** it MUST still display the standard recitation header and instructions footer.

