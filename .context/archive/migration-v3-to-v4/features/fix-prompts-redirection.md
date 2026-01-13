# Specification: fix-prompts-redirection

## 1. User Intent (The Goal)
Fix the bug where `cdd prompts --system > example.txt` prints to the screen (stderr) instead of being redirected to the file (stdout).

## 2. Relevant Context (The Files)
- `internal/cmd/prompts.go`: Implementation of the prompts command.
- `internal/cmd/root.go`: Root command configuration.
- `internal/cmd/list.go`: Also affected by same issue.

## 3. Context Analysis (Agent Findings)
- `cmd.Println` and `cmd.Printf` in this project seem to be writing to `stderr` by default.
- Redirection `> file.txt` results in an empty file and terminal output.
- Redirection `2> file.txt` captures the output, confirming it's on `stderr`.

## 4. Scenarios (Acceptance Criteria)
Feature: CLI Output Redirection
  Scenario: Redirect prompts output to a file
    Given I have the cdd CLI
    When I run `cdd prompts --system > output.txt`
    Then the file `output.txt` should contain the prompt content
    And the terminal should show no output.
