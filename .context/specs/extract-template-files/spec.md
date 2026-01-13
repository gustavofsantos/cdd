# Spec: Template Rendering System

## Status
✅ **IMPLEMENTED** (2026-01-13)

## Overview
The CDD CLI uses an embedded template system to generate track files (spec.md, plan.md, decisions.md) when initializing the environment or starting new tracks. This spec defines the shared template rendering architecture.

## Requirements

### Template Embedding
* WHEN the CDD binary is built, the system SHALL embed all template files from `internal/cmd/templates/` into the binary using Go's `embed` package
* WHERE `internal/cmd/templates.go`, the system SHALL declare a package-level `trackTemplates embed.FS` variable with the `//go:embed templates/*` directive
* WHEN the binary is executed, the system SHALL access embedded templates without requiring external files

### Template Rendering
* WHEN a command needs to render a template, the system SHALL use the shared `renderTrackTemplate()` function
* WHERE `renderTrackTemplate()` is called, the system SHALL accept three parameters: template name, template filename in FS, and template data
* WHEN rendering a template, the system SHALL use `template.ParseFS()` to parse the embedded template
* WHEN template parsing succeeds, the system SHALL execute the template with the provided `trackData` struct
* WHERE template execution completes, the system SHALL return the rendered content as a byte slice

### Code Reusability
* WHERE both `init.go` and `start.go` need template rendering, the system SHALL use the shared code from `templates.go`
* WHEN adding new commands that need templates, the system SHALL reuse the existing `renderTrackTemplate()` function
* WHERE template rendering logic exists, the system SHALL maintain it in a single location to avoid duplication

### Template Data Structure
* WHEN rendering templates, the system SHALL use the `trackData` struct containing:
  - `TrackName`: The name of the track being created
  - `CreatedAt`: Timestamp of track creation in "Mon Jan 2 15:04:05 MST 2006" format

## Implementation Details

### File Structure
```
internal/cmd/
├── templates.go          # Shared template rendering logic
├── templates/            # Template files (embedded)
│   ├── spec.md          # Standard track spec template
│   ├── plan.md          # Standard track plan template
│   ├── decisions.md     # Standard track decisions template
│   ├── setup_spec.md    # Setup track spec template
│   └── setup_plan.md    # Setup track plan template
├── init.go              # Uses shared template logic
└── start.go             # Uses shared template logic
```

### Components
- **templates.go**: Contains `trackTemplates`, `trackData`, and `renderTrackTemplate()`
- **init.go**: Creates setup track using `setup_spec.md` and `setup_plan.md`
- **start.go**: Creates regular tracks using `spec.md` and `plan.md`

## Relevant Files
* `internal/cmd/templates.go` - Shared template rendering module
* `internal/cmd/init.go` - Initialization command using templates
* `internal/cmd/start.go` - Track creation command using templates
* `internal/cmd/templates/*.md` - Template files embedded in binary
* `internal/cmd/init_test.go` - Tests for init command behavior

## Related Tracks
* `20260113143410_refactor-init-templates` - Implementation track (archived)
* `20260113131405_refactor-start-templates` - Previous related work (archived)

## Notes
- Templates are embedded at compile time, requiring recompilation for template changes
- This is acceptable for a CLI tool where templates are core functionality
- The shared module reduces code duplication by ~30 lines
- All tests pass, confirming backward compatibility
