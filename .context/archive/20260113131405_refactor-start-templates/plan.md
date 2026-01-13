# Plan for refactor-start-templates
[x] 🔴 Test: Verify templates are loaded from files (Manual/Unit)
[x] 🟢 Impl: Create `internal/cmd/templates/` directory
[x] 🟢 Impl: Create `spec.md`, `plan.md`, `decisions.md` templates in `internal/cmd/templates/`
[x] 🟢 Impl: Update `internal/cmd/start.go` to use `go:embed` and `text/template`
[x] 🔵 Refactor: Ensure clean error handling if templates are missing
