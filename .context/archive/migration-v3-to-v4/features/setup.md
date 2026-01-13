# Spec: Project Initialization

## User Intent
The user is initializing the CDD environment for the `cdd` project itself. This is a "Brownfield" initialization where the codebase already exists. The goal is to map the project structure and populate the Global Context files (`product.md`, `tech-stack.md`, `patterns.md`). The user has clarified that the focus is **the entire project**, as we are working on `cdd` to improve `cdd`.

## Relevant Context
- `go.mod`: Project dependencies (Cobra, Charmbracelet).
- `internal/cmd/`: Contains the Cobra command definitions for the CLI.
- `.context/`: The CDD context directory structure.
- `cmd/cdd/main.go`: Likely the entry point (implied by standard Go layout).

## Context Analysis
- The project allows recursive improvement (building the tool using the tool).
- Key architectural pattern: Cobra CLI application.
- State management: File-based in `.context/`.
- UI: TUI-based using Charmbracelet libraries.
