# Track: all-skills-support

## 1. User Intent
Add a `--all` flag to the `agents` command that installs skills for all supported platforms. Remove the default "agent" target—the user must explicitly specify a target.

## 2. Relevant Context
- The `agents` command exists to manage skill installation
- Supported platforms include Windows, macOS, and Linux
- Current behavior includes a default "agent" target that should be removed

## 3. Requirements (EARS)

Patterns:
    Ubiquitous: The <system> shall <response>
    Event-driven: When <trigger>, the <system> shall <response>
    State-driven: While <state>, the <system> shall <response>
    Unwanted: If <condition>, then the <system> shall <response>
    Optional: Where <feature>, the <system> shall <response>

### Functional Requirements

**Ubiquitous:**
- The `agents` command shall support a `--all` flag

**Event-driven:**
- When the user provides the `--all` flag, the system shall install skills for all supported platforms (Windows, macOS, Linux)
- When the user invokes the `agents` command without specifying a target, the system shall return an error and request an explicit target

**Unwanted Behavior:**
- If a target is not explicitly specified and no `--all` flag is provided, the command shall not default to "agent"

### Non-Functional Requirements
- The removal of the default "agent" target shall be backward-incompatible and documented
- The `--all` flag shall support installing on all platform configurations simultaneously