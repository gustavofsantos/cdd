# Specification: semantic-versioning

## 1. User Intent (The Goal)
Automate semantic versioning and release process for the CDD Tool Suite using GoReleaser.

## 2. Relevant Context (The Files)
- `go.mod`: Project manifest.
- `.github/workflows/release.yml`: Current release pipeline.
- `.goreleaser.yaml`: (To be created) GoReleaser configuration.

## 3. Context Analysis (Agent Findings)
- The user prefers to avoid Node.js/JavaScript tools.
- GoReleaser is the standard for Go projects to handle builds and releases.
- GoReleaser typically derives versions from Git tags.
- We still want the tool to report its own version.

## 4. Scenarios (Acceptance Criteria)
### Scenario: Setup GoReleaser
Given I am in the root of the project
When I initialize goreleaser config
Then a `.goreleaser.yaml` should exist.

### Scenario: Running a Release
Given I have created a new git tag (e.g., `git tag -a v1.1.0 -m "Release v1.1.0"`)
When I run `goreleaser release --snapshot` (for local testing) or the CI runs
Then GoReleaser should:
1. Build binaries for multiple platforms.
2. Generate a changelog based on git commits.
3. Prepare a GitHub Release.

### Scenario: Version Sychronization
Given I have a `version.go` file
When I build using GoReleaser
Then the `Version` variable should be injected during build time using ldflags, ensuring the CLI reports the correct tag version.
