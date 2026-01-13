# Proposed Global Context Updates
> Add notes here if product.md or tech-stack.md needs updating.

- Added Time Tracking to `cdd`:
    - `start` now creates `metadata.json` with a `started_at` timestamp.
    - `archive` now calculates and prints the duration of the track.
    - Uses `encoding/json` and `time` packages.
