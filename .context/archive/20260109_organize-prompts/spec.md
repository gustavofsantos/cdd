# Spec: Organize Prompts

The goal is to move the prompts from a hidden internal directory to a more accessible root directory and rename them to Markdown files (`.md`). This will make it easier to manage and add new prompts in the future.

## User Intent
"I want to explore and include more prompts. But to ease the process, I need to move the prompts to a easier place, also define them as markdown files, they are .txt"

## Relevant Context
- `internal/cmd/prompts/bootstrap.txt`
- `internal/cmd/prompts/system.txt`
- `internal/cmd/init.go`

## Context Analysis
- Current location: `internal/cmd/prompts/`
- Current format: `.txt` (though content is Markdown)
- Usage: Embedded in `internal/cmd/init.go` for `cdd init --bootstrap-prompt` and `cdd init --system-prompt`.
- Refactoring needed: Move files to `prompts/` at root, rename to `.md`, and update Go embed directives.

## Scenarios

### Scenario 1: Relocate and Rename Prompts
- **Given** the current `.txt` prompts in `internal/cmd/prompts/`
- **When** I move them to `prompts/` and rename them to `.md`
- **Then** the files should exist at the new location with the new extension.

### Scenario 2: Update Go Embed
- **Given** the new location and extension of the prompt files
- **When** I update the `//go:embed` directives in `internal/cmd/init.go`
- **Then** the application should compile without errors.

### Scenario 3: Verify Functionality
- **Given** the updated application
- **When** I run `go run cmd/cdd/main.go init --bootstrap-prompt`
- **Then** it should output the content of `prompts/bootstrap.md`.
- **When** I run `go run cmd/cdd/main.go init --system-prompt`
- **Then** it should output the content of `prompts/system.md`.
