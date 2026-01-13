# Specification: update-archive-metadata

## 1. User Intent (The Goal)
Record the completion timestamp (`archived_at`) in `metadata.json` when a track is archived.

## 2. Relevant Context (The Files)
- `internal/cmd/archive.go`: Implementation of the archive command.

## 3. Context Analysis (Agent Findings)
- Current implementation reads `metadata.json` to calculate duration but does not update it.
- The `metadata.json` file is preserved in the archive but remains static from the time of `cdd start`.

## 4. Scenarios (Acceptance Criteria)
Feature: Archive Metadata Update
  Scenario: Record archiving timestamp
    Given I have an active track "test-archive-metadata"
    When I run `cdd archive test-archive-metadata`
    Then the file `.context/archive/<TIMESTAMP>_test-archive-metadata/metadata.json` should contain an "archived_at" field.
