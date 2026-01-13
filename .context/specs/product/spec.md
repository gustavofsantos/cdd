# Product Specification: CDD Tool Suite

## Goal
A CLI application designed to facilitate Context-Driven Development (CDD). It helps developers and AI agents manage context, track progress, and maintain a history of decisions and changes through a file-based state management system.

## Core Value Proposition
- **Context Management:** Organized structure for project context (`.context/` directory).
- **Track Isolation:** Work on specific tasks in isolated "tracks" to prevent context pollution.
- **AI-Ready:** Designed to work seamlessly with AI agents by providing structured context and a specialized prompt system (Strategist/Tactician model).
- **Workflow Automation:** Commands to start, archive, and manage tracks.
- **Continuous Consolidation:** Built-in mechanisms to promote local track context to global project knowledge.

## Key Users
- **Software Engineers:** Human developers orchestrating the project.
- **AI Agents:** Context Gardeners (Consolidation) and Tacticians (Implementation).

## Requirement: Global Context Architecture
The system MUST maintain a set of permanent records in the `.context/` directory:
- `product.md`: Product vision and domain logic.
- `tech-stack.md`: Technology choices and versions.
- `workflow.md`: Development processes and commands.
- `patterns.md`: Architectural and coding patterns.

## Requirement: Dogfooding
The tool MUST be used to develop itself. Every feature of the CDD Tool Suite should be planned and implemented using the CDD process.
