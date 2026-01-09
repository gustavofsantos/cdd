# Plan: Organize Prompts

Move existing prompts to a root `prompts/` directory as `.md` files and update the codebase to use them.

## Phase 1: Relocation
- [x] Create a root `prompts/` directory.
- [x] Move `internal/cmd/prompts/bootstrap.txt` to `prompts/bootstrap.md`.
- [x] Move `internal/cmd/prompts/system.txt` to `prompts/system.md`.
- [x] Remove the empty `internal/cmd/prompts/` directory.

## Phase 2: Code Update
- [x] Update `internal/cmd/init.go` to embed prompts from the new location.
    - Used a separate `prompts` package at root to allow embedding from a root directory.

## Phase 3: Verification
- [x] Run `go run cmd/cdd/main.go init --bootstrap-prompt` and verify output.
- [x] Run `go run cmd/cdd/main.go init --system-prompt` and verify output.
