# Track: context-packer

## 1. User Intent
Implement a "just-in-time" context fetcher that compresses and filters the global specs in `.context/specs/` based on a user-provided topic focus. The command should help large legacy projects manage cognitive load by delivering only the relevant paragraphs from specs, reducing context window bloat when using AI agents.

## 2. Relevant Context
- `.context/specs/`: Directory containing all global specification files (log.md, view.md, help.md, agents.md, amp-toolbox.md)
- `internal/cmd/view.go`: Example command that reads and renders files from `.context/`
- `internal/platform/`: File system abstraction layer used by all commands
- `internal/cmd/log.go`: Example command pattern for argument parsing and execution
- The system uses Cobra for CLI commands with dependency injection pattern
- Specs follow markdown format with headers, requirements, and context sections

## 3. Requirements (EARS)
- Ubiquitous: The `pack` command shall read all markdown files from `.context/specs/` and output relevant content filtered by topic focus.
- Ubiquitous: The `pack` command shall accept a `--focus <topic>` flag to specify the search target.
- Event-driven: When the user invokes `cdd pack --focus <topic>`, the system shall search specs and output matching paragraphs with minimal overhead.
- Ubiquitous: The system shall compress output by extracting only paragraphs that contain semantic relevance to the topic (not full files).
- Ubiquitous: The system shall use either embedding-based search or fuzzy string matching as the search mechanism.
- Optional: Where `--raw` flag is provided, the system shall output plain text instead of formatted markdown.
- Unwanted: If no matching paragraphs are found, the system shall return a helpful message indicating no matches rather than empty output.
- Ubiquitous: The system shall support both directory-based targets and individual spec searches.
