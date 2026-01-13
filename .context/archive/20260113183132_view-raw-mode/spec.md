# Track: view-raw-mode

## 1. User Intent
Enable `cdd view` to be used in shell pipelines (e.g., with `fzf`). 
The current output is formatted with Markdown/Glamour, which includes emojis and styling that make it hard to parse programmatically. 
The goal is to:
- Detect if the output is not a TTY or add a `--raw` flag.
- In "raw" mode, output a simple list of track names (one per line) without headers, emojis, or Markdown formatting.

## 2. Relevant Context
- `internal/cmd/view.go`: Command implementation and output rendering.
- `github.com/mattn/go-isatty`: (Already in `go.mod`) Can be used to detect TTY.

## 4. Scenarios
Feature: Raw output mode for view command

  Scenario: Pipe active tracks to another command
    When I run `cdd view | cat`
    Then it should output only the list of active track names, one per line
    And it should NOT include "Active Tracks" header or emojis

  Scenario: Pipe archived tracks to another command
    When I run `cdd view --archived | cat`
    Then it should output only the list of clean archived track names, one per line

  Scenario: Explicit raw mode with flag
    When I run `cdd view --raw`
    Then it should output the raw list even if in a TTY
