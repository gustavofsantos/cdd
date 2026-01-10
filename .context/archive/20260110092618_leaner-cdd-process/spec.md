# Specification: Leaner CDD Process Improvements

## User Intent
Simplify the CDD process by reducing markdown file clutter and improving track archiving.
1.  **Delete Scratchpad:** The `archive` command should delete `scratchpad.md` from the track directory before archiving.
2.  **Improved Timestamping:** Archived tracks should include hour, minute, and second in their timestamp (e.g., `20060102150405_track-name`) for better sorting.

## Relevant Context
*   `internal/cmd/archive.go`: Implementation of the `archive` command.
*   `internal/platform/fs.go`: File system abstraction.
*   `internal/cmd/archive_test.go`: Existing tests for the archive command.

## Context Analysis
The `archive` command currently uses `YYYYMMDD_track-name` for the destination directory. Adding time will resolve sorting issues when multiple tracks are archived on the same day. Deleting the scratchpad reduces storage of ephemeral data.

## Test Reference
*   `internal/cmd/archive_test.go`: `TestArchiveCmd_CleanupAndTimestamp` verifies both scratchpad deletion and full timestamping.
