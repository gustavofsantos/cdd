# System Architecture

## 1. High-Level Design (System Context)
A standalone CLI application written in Go that manages a file-based state machine for context-driven development. The system interacts with:
- **Developer** (primary actor): Issues commands via terminal
- **AI Agent** (secondary actor): Consumes prompts and context files, produces code/decisions
- **File System**: Stores all state in `.context/` directory structure
- **Version Control**: Integrates with Git for change tracking

## 2. Architectural Pattern
**Command Pattern + File-Based State Machine**
- CLI commands implemented using Cobra framework
- Each command encapsulates a specific operation (start, archive, recite, log, etc.)
- State is persisted entirely through markdown files and JSON metadata
- No database or external services required

## 3. Core Components (Container Diagram level)
| Component | Responsibility | Dependencies |
| :--- | :--- | :--- |
| **CLI Commands** (`internal/cmd/`) | Handle user interactions, orchestrate workflows | Cobra, platform.FileSystem |
| **Prompt Manager** (`prompts/`) | Serve embedded AI agent prompts for different phases | Go embed, template rendering |
| **Template Engine** (`internal/cmd/templates/`) | Render track and context file templates | Go text/template |
| **Platform Abstraction** (`internal/platform/`) | File system operations with testable interface | OS filesystem |
| **Context Store** (`.context/`) | Persistent storage for tracks, specs, and global context | File system |

## 4. Cross-Cutting Concerns
* **Auth/Security:** None required (local CLI tool, no authentication)
* **Logging/Observability:** Stdout/stderr via Cobra command output, Glamour for markdown rendering
* **State Management:** File-based persistence in `.context/` directory with JSON metadata for timestamps

## 5. Architectural Decision Log (ADR) Summary
* ADRs are maintained per-track in `.context/tracks/<track-name>/decisions.md`
* Archived ADRs are preserved in `.context/archive/<track-name>/decisions.md`
