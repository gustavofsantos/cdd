# Specification: time-tracking

## 1. User Intent (The Goal)
> Implement time tracking for tracks. Capture start time, archive time, and calculate the difference.

## 2. Relevant Context (The Files)
- `internal/cmd/start.go`
- `internal/cmd/archive.go`
- `internal/cmd/start_test.go`
- `internal/cmd/archive_test.go`

## 3. Context Analysis (Agent Findings)
> - Current Behavior: `start` creates files but stores no timestamp (except in `decisions.md` as text). `archive` moves files but calculates no duration.
> - Proposed Changes:
    - Modify `start` to create `metadata.json` with `started_at` timestamp.
    - Modify `archive` to read `metadata.json`, calculate duration, and display it.

## 4. Test Reference
- `internal/cmd/start_test.go`
- `internal/cmd/archive_test.go`
