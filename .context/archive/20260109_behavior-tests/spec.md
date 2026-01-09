# Specification: behavior-tests

## 1. User Intent (The Goal)
The user wants to implement tests to describe the behaviors of the `cdd` tool.
Since the current implementation likely relies on direct system calls (file system, stdout), the user anticipates a need for refactoring to make the code testable (Dependency Injection).

## 2. Relevant Context (The Files)
- `internal/cmd/start.go`
- `internal/cmd/recite.go`
- `internal/cmd/list.go`
- `internal/cmd/archive.go`
- `internal/cmd/dump.go`
- `internal/cmd/log.go`
- `internal/cmd/init.go` (if exists? checked `ls` and it exists)

## 3. Context Analysis (Agent Findings)
- The current commands in `internal/cmd/` use `os.MkdirAll`, `os.WriteFile`, `os.ReadFile`, `os.Stat` and `fmt.Printf` directly.
- VALIDATION: This makes unit testing impossible without side effects (creating real files) and capturing stdout.
- To test properly, we need to:
    1.  Abstract File System operations into an interface.
    2.  Use `cobra`'s `cmd.OutOrStdout()` instead of `fmt.Printf` for output capturing.
    3.  Inject the File System dependency into the commands (or a context struct).

## 4. Scenarios (Acceptance Criteria)
Feature: `cdd start`
  Scenario: Create a new track
    Given the track "unit-test-track" does not exist
    When I run "cdd start unit-test-track"
    Then the directory ".context/tracks/unit-test-track" should be created
    And the file "spec.md" should be created
    And the file "plan.md" should be created
    And the output should contain "Track 'unit-test-track' initialized."

  Scenario: Track already exists
    Given the track "existing-track" already exists
    When I run "cdd start existing-track"
    Then the output should contain "Error: Track 'existing-track' exists."
    And the exit code should be 1 (or return error)

Feature: `cdd recite`
  Scenario: Recite plan
    Given the track "active-track" exists with plan items
    When I run "cdd recite active-track"
    Then the output should display the plan content

