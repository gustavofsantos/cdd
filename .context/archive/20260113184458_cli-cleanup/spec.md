# Command Simplification and Cleanup

## 1. Overview
The goal is to simplify the `cdd` CLI by removing obsolete or redundant commands. This improves maintainability and user experience (less surface area).

## 2. Relevant Context
- `internal/cmd/dump.go`: Currently writes to `scratchpad.md` (deprecated artifact).
- `internal/cmd/list.go`: Lists active/archived tracks. Heavily overlaps with `view` (Dashboard).
- `internal/cmd/view.go`: The central dashboard. Supersedes `list` functionality.

## 3. Options Analysis
- **`dump`**: The `scratchpad.md` artifact has been removed from the core workflow (System Prompt). Therefore, `dump` is dead code.
    - **Recommendation**: DELETE.
- **`list`**: The `view` command (without arguments) acts as a dashboard, listing tracks. `view --archived` lists archived. `view --raw` provides pipe-friendly lists. `list` is therefore completely redundant.
    - **Recommendation**: DELETE.
- **`recite`**: The user has identified this as a critical state-machine command.
    - **Recommendation**: KEEP.

## 4. Proposed Changes
1.  **Remove `dump`**: Delete `internal/cmd/dump.go` and `dump_test.go`. Remove from `root.go`.
2.  **Remove `list`**: Delete `internal/cmd/list.go` and `list_test.go`. Remove from `root.go`.
3.  **Update Documentation**:
    - Update `README.md` to remove references to `dump` and `list`.
    - Point users to `cdd view` for listing tracks.

## 5. Scenarios

### Scenario: List is Gone
When I run `cdd list`
Then I should see "unknown command" error

### Scenario: Dump is Gone
When I run `cdd dump`
Then I should see "unknown command" error

### Scenario: View Replaces List
Given I have active tracks
When I run `cdd view --raw`
Then I should see the list of tracks (matching old `cdd list`)

