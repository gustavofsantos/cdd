# Standards Specification

## Requirement: Technology Stack
The project MUST adhere to the following core technologies:
- **Language:** Go (Targeting version 1.24.3+)
- **CLI Framework:** Cobra
- **TUI/Styling:** Charmbracelet (lipgloss, glamour, x/term)
- **Markdown Rendering:** goldmark, glamour
- **Build & Release:** GoReleaser

#### Scenario: Building the application
- **Given** a Go environment
- **When** the developer runs `go build ./cmd/cdd`
- **Then** a binary named `cdd` should be generated.

## Requirement: Project Patterns
The codebase SHOULD follow established patterns:
- **Dependency Injection:** Abstract file system operations using the `FileSystem` interface (`internal/platform/fs.go`) to enable unit testing with `MockFileSystem`.
- **Cobra Commands:** Command definitions MUST be located in `internal/cmd/`.
- **Error Handling:** Standard Go error checking MUST be used consistently.

## Requirement: Development Workflow
The CDD workflow MUST follow these stages:
1. **Recite:** Load and review current plan/spec.
2. **Spec:** Define *what* is being built in a track's `spec.md`.
3. **Plan:** Define *how* it will be built in `plan.md`.
4. **Implement:** Write code, preferably following TDD.
5. **Archive:** Close the track, moving updates to the Global Context via the Inbox.

## Requirement: Testing Standard
- **Unit Testing:** Commands and logic MUST be tested using the standard Go `testing` package.
- **Mocking:** Use the `MockFileSystem` to verify file-based side effects without touching the real disk during tests.

## Requirement: Semantic Versioning
The project MUST follow Semantic Versioning (SemVer) for releases.
- **Major:** Incompatible API changes.
- **Minor:** Functionality in a backwards-compatible manner.
- **Patch:** Backwards-compatible bug fixes.
