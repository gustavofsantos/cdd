# Plan for refactor-init-templates

## Phase 1: Extract Common Template Logic
[x] 🔴 Test: Verify current init.go behavior works correctly
[x] 🟢 Impl: Move `trackData` struct and `renderTrackTemplate` function to a shared location (e.g., `templates.go`)
[x] 🔵 Refactor: Update both `init.go` and `start.go` to import from the shared location

## Phase 2: Add Embed Support to init.go
[x] 🔴 Test: Verify templates are accessible via embed.FS
[x] 🟢 Impl: Add `embed` import to `init.go`
[x] 🟢 Impl: Add `//go:embed templates/*` directive to `init.go`
[x] 🟢 Impl: Declare `trackTemplates embed.FS` variable in `init.go`
[x] 🔵 Refactor: Remove any hardcoded template content if present

## Phase 3: Verification
[x] 🔴 Test: Run `cdd init` in a test directory to verify it creates the setup track correctly
[x] 🔴 Test: Verify all template files are rendered with correct data
[x] 🟢 Impl: Update any documentation if needed

