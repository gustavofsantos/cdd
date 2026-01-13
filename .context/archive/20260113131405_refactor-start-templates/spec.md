# Track: refactor-start-templates
**Target Spec:** .context/specs/lifecycle/spec.md

## Context
- `internal/cmd/start.go`: Contains hardcoded templates.
- `internal/cmd/templates/`: New home for extracted templates.

## Proposed Changes
### ADDED Requirements
* **Requirement: Template Externalization**
    * The CLI SHALL use external template files instead of hardcoded strings for track initialization.
    * Templates MUST be embedded into the binary using `go:embed`.
    * Templates SHOULD support variable substitution (e.g., using `text/template`).

### MODIFIED Requirements
* **Requirement: Track Management**
    * The `cdd start` command MUST load templates from embedded files.
