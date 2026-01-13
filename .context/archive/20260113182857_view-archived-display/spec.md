# Track: view-archived-display

## 1. User Intent
Improve the visibility of archived tracks in the dashboard. Specifically:
- When listing archived tracks (`cdd view --archived`), display the track name without the timestamp prefix.
- This ensures consistency between how tracks are listed and how they are queried (e.g., `cdd view setup --archived`).

## 2. Relevant Context
- `internal/cmd/view.go`: Contains the dashboard rendering logic for archived tracks.
- `internal/cmd/archive.go`: Defines the timestamp prefix format.

## 4. Scenarios
Feature: Clean display of archived track names

  Scenario: List archived tracks
    Given archived tracks "20260113161017_setup" and "20260113182345_view-flags" exist
    When I run `cdd view --archived`
    Then it should display the names "setup" and "view-flags"
    And it should NOT display the timestamp prefixes in the list
