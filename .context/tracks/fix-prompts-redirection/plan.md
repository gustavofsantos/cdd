# Plan for fix-prompts-redirection

- [x] 🔴 Test: Create a reproduction script `reproduce_bug.sh` that validates output streams.
- [x] 🟢 Impl: Update `internal/cmd/root.go` or individual commands to ensure they write to `stdout`.
- [x] 🔵 Refactor: Ensure consistency across all `cmd.Println` calls.
- [x] 🔵 Verification: Run `reproduce_bug.sh` and ensure it passes.
