# Product Context

**Product Name:** CDD Tool Suite
**Description:** A CLI application designed to facilitate Context-Driven Development (CDD). It helps developers and AI agents manage context, track progress, and maintain a history of decisions and changes through a file-based state management system.

## Core Value Proposition
- **Context Management:** Organized structure for project context (.context/ directory).
- **Track Isolation:** Work on specific tasks in isolated "tracks" to prevent context pollution.
- **AI-Ready:** Designed to work seamlessly with AI agents by providing structured context.
- **Workflow Automation:** Commands to start, archive, and manage tracks.

## Key Users
- Software Engineers
- AI Agents

## Domain Logic
- **Tracks:** Active workspaces in `.context/tracks/`.
- **Archive:** Completed workspaces in `.context/archive/`.
- **Global Context:** `product.md`, `tech-stack.md`, etc.
- **CLI Commands:** `init`, `start`, `recite`, `log`, `dump`, `archive`, `list`.
