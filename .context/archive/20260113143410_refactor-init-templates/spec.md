# Track: refactor-init-templates

## Proposed Changes

### MODIFIED Requirements
* WHEN the `cdd init` command is executed, the system SHALL use embedded template files from `internal/cmd/templates/` to render track files (previously: uses `renderTrackTemplate` function but lacks the `//go:embed` directive)
* WHERE `internal/cmd/init.go` file, the system SHALL include the `//go:embed templates/*` directive to embed template files (previously: missing this directive)
* WHERE `internal/cmd/init.go` file, the system SHALL declare the `trackTemplates embed.FS` variable (previously: missing this variable)
* WHERE `internal/cmd/init.go` file, the system SHALL import the `embed` package (previously: missing this import)

### ADDED Requirements
* WHEN rendering templates in `init.go`, the system SHALL reuse the existing `renderTrackTemplate` function from `start.go`
* WHERE both `init.go` and `start.go`, the system SHALL share common template rendering logic to avoid code duplication

## Relevant Files
* `internal/cmd/init.go` - Main file to be refactored, needs embed directive and variable
* `internal/cmd/start.go` - Reference implementation showing the correct pattern
* `internal/cmd/templates/setup_spec.md` - Template used by init command
* `internal/cmd/templates/setup_plan.md` - Template used by init command
* `internal/cmd/templates/decisions.md` - Template used by both init and start commands
