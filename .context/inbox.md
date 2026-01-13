
---
###### Archived at: 2026-01-13 18:23:45 | Track: view-flags

# Track: view-flags

## 1. User Intent
The `cdd view` command currently displays too much information in dashboard mode. The goal is to:
- Simplify the default dashboard output.
- Add flags to control visibility of different contexts (Inbox, Archived Tracks).
- Add support for viewing specific files (Plan, Spec) for both active and archived tracks.

## 2. Relevant Context
- `internal/cmd/view.go`: Primary implementation of the `view` command.
- `.context/tracks/`: Directory for active tracks.
- `.context/archive/`: Directory for archived tracks.

## 4. Scenarios
Feature: Enhanced view command with flags

  Scenario: Default view (no track specified)
    Given there are active tracks
    When I run `cdd view`
    Then it should only show the list of active tracks

  Scenario: Show inbox with flag
    When I run `cdd view --inbox`
    Then it should show the Context Inbox

  Scenario: List archived tracks
    Given there are archived tracks
    When I run `cdd view --archived`
    Then it should show the list of archived tracks

  Scenario: View specific track plan (default)
    Given a track "my-feature" exists
    When I run `cdd view my-feature`
    Then it should show the Next Tasks for "my-feature"

  Scenario: View specific track spec
    Given a track "my-feature" exists
    When I run `cdd view my-feature --spec`
    Then it should show the content of "spec.md" for "my-feature"

  Scenario: View specific archived track plan
    Given an archived track "20260113161017_setup" exists
    When I run `cdd view setup --archived`
    Then it should show the Next Tasks for the archived track "setup"

  Scenario: View specific archived track spec
    Given an archived track "20260113161017_setup" exists
    When I run `cdd view setup --archived --spec`
    Then it should show the content of "spec.md" for the archived track "setup"


---
###### Archived at: 2026-01-13 18:28:57 | Track: view-archived-display

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

