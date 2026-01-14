# Plan for all-skills-support

## Phase 0: Analysis
- [x] 🗣️ Phase 0: Alignment & Requirements (Fill `spec.md` using EARS)

## Phase 1: Architecture
- [x] 📝 Phase 1: Approval (User signs off)

## Phase 2: Implementation

### Requirement Mapping:
1. **Support `--all` flag** → Tasks 1-2
2. **Mandate explicit target** → Tasks 3-4
3. **Install across platforms** → Task 5

### Atomic TDD Tasks:

- [x] **Task 1:** Add `--all` flag parser to `agents` command
  - RED: Test that `--all` flag is recognized
  - GREEN: Parse flag from CLI arguments
  - REFACTOR: Extract flag parsing logic

- [x] **Task 2:** Implement multi-platform skill installation logic
  - RED: Test that `--all` installs for Windows, macOS, Linux
  - GREEN: Add platform-specific installation branches
  - REFACTOR: DRY up platform detection

- [x] **Task 3:** Remove default "agent" target fallback
  - RED: Test that missing target throws error
  - GREEN: Remove default target logic
  - REFACTOR: Improve error messaging

- [x] **Task 4:** Add target validation with helpful error message
  - RED: Test error when no target specified and no `--all` flag
  - GREEN: Validate target or `--all` flag is provided
  - REFACTOR: Enhance error message with examples

- [x] **Task 5:** Integration test for `--all` flag across all platforms
  - RED: Test `agents --all` installs all platform skills
  - GREEN: Verify installation completes successfully
  - REFACTOR: Add platform verification checks
