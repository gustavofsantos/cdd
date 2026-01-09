# Specification: update-installation-instructions

## 1. User Intent (The Goal)
> Update the installation instructions considering goreleaser and the way that the release is available.

## 2. Relevant Context (The Files)
- `INSTALLATION.md`: Contains the current installation instructions.
- `.context/tech-stack.md`: Mentions GoReleaser.
- `.goreleaser.yaml`: Configuration for the release process.
- User-provided image: Shows the actual release assets naming convention.

## 3. Context Analysis (Agent Findings)
- **Current Behavior:** `INSTALLATION.md` assumes the binary is available directly as a file (e.g., `cdd-linux-amd64`).
- **Findings:**
    - `.goreleaser.yaml` configures archives (`tar.gz` for Linux/macOS, `zip` for Windows).
    - The naming template is `{{ .ProjectName }}_{{ title .Os }}_{{ .Arch }}` (with mapping for `amd64` to `x86_64`).
    - The actual files are like `cdd_Linux_x86_64.tar.gz`.
- **Proposed Changes:**
    - Update `INSTALLATION.md` to guide the user through downloading the archive, extracting it, and installing the binary.
    - Provide updated code examples for Linux, macOS, and Windows.

## 4. Scenarios (Acceptance Criteria)
Feature: update-installation-instructions
  Scenario: Linux Installation Instructions Updated
    Given I am on a Linux system
    When I read `INSTALLATION.md`
    Then I should see instructions to download the `cdd_Linux_x86_64.tar.gz` archive (or arm64), extract it, and move it to a PATH directory.

  Scenario: macOS Installation Instructions Updated
    Given I am on a macOS system
    When I read `INSTALLATION.md`
    Then I should see instructions to download the `cdd_Darwin_x86_64.tar.gz` (or arm64), extract it, and move it to a PATH directory.

  Scenario: Windows Installation Instructions Updated
    Given I am on a Windows system
    When I read `INSTALLATION.md`
    Then I should see instructions to download the `cdd_Windows_arm64.zip` (based on the image), extract it, and move it to a PATH directory.

