# Track: help-documentation

## 1. User Intent
The goal is to improve the user experience and discoverability of `cdd` commands by providing extensive documentation when using the `--help` flag. This includes detailed long descriptions, usage examples, and clear explanations of arguments and flags for every command.

## 2. Relevant Context
- `internal/cmd/root.go`: Root command definition.
- `internal/cmd/start.go`: `start` command logic.
- `internal/cmd/archive.go`: `archive` command logic.
- `internal/cmd/view.go`: `view` command logic.
- `internal/cmd/log.go`: `log` command logic.
- `internal/cmd/delete.go`: `delete` command logic.
- `internal/cmd/init.go`: `init` command logic.
- `internal/cmd/prompts.go`: `prompts` command logic.
- `internal/cmd/recite.go`: `recite` command logic.
- `internal/cmd/version.go`: `version` command logic.

## 4. Scenarios
Feature: help-documentation
  Scenario: View detailed help for a command
    Given I have the `cdd` tool installed
    When I run `cdd <command> --help`
    Then I should see a detailed description of what the command does
    And I should see at least one usage example
    And I should see descriptions for all available flags

  Scenario: Root help shows clear overview
    When I run `cdd --help`
    Then I should see a brief overview of the CDD protocol and tool suite
    And I should see a list of all available commands with clear, concise summaries

