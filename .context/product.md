# Product Context

**Product Name:** CDD Tool Suite
**Description:** A CLI application designed to facilitate Context-Driven Development (CDD).

## System Specifications (The Living Documentation)
The definitive truth about the system's capabilities and standards is now maintained in the `specs/` directory:

- [Product Vision & Core Value Prop](specs/product/spec.md)
- [Lifecycle & CLI Commands](specs/lifecycle/spec.md)
- [Standards & Patterns](specs/standards/spec.md)
- [Prompt Engineering & AI Roles](specs/prompts/spec.md)

## Domain Logic Summary
- **Tracks:** Active workspaces in `.context/tracks/`.
- **Archive:** Completed workspaces in `.context/archive/`.
- **Global Context:** Permanent records managed via Specifications.
- **Context Inbox:** Ephemeral queue for pending changes from closed tracks (`.context/inbox.md`).
- **Development Cycle:** Recite -> Spec -> Plan -> Implement -> Archive -> Garden.

## Dogfooding
The tool is used to develop itself, ensuring immediate feedback on UX and functionality.
