# Product Context

**Product Name:** CDD Tool Suite
**Description:** A CLI application designed to facilitate Context-Driven Development (CDD). It helps developers and AI agents manage context, track progress, and maintain a history of decisions and changes through a file-based state management system.

## Core Value Proposition
- **Context Management:** Organized structure for project context (.context/ directory).
- **Track Isolation:** Work on specific tasks in isolated "tracks" to prevent context pollution.
- **AI-Ready:** Designed to work seamlessly with AI agents by providing structured context and a specialized prompt system (Strategist/Tactician model).
- **Workflow Automation:** Commands to start, archive, and manage tracks.
- **Continuous Consolidation:** Built-in mechanisms to promote local track context to global project knowledge.

## Key Users
- Software Engineers
- AI Agents (Context Gardeners & Tacticians)

## Domain Logic
- **Tracks:** Active workspaces in `.context/tracks/`.
- **Archive:** Completed workspaces in `.context/archive/`.
- **Global Context:** Permanent records in `product.md`, `tech-stack.md`, `workflow.md`, and `patterns.md`.
- **Context Inbox:** Ephemeral queue for pending changes from closed tracks (`.context/inbox.md`).
- **CLI Commands:** `init`, `start`, `recite`, `log`, `dump`, `archive`, `list`, `version`.
- **Dogfooding:** The tool is used to develop itself, ensuring immediate feedback on UX and functionality.
