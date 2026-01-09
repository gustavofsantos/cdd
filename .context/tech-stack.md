# Tech Stack

## Core
- **Language:** Go (1.24.3)
- **Framework:** Cobra (CLI Library)

## Libraries
- **UI/TUI:** Charmbracelet (lipgloss, glamour, x/term) - Used for rich terminal output and styling.
- **Markdown:** goldmark, glamour - For rendering Markdown in the terminal.

## Build & CI
- **Build System:** `go build`
- **CI/CD:** GitHub Actions (`.github/workflows/release.yml`)

## Storage
- **Data Store:** Local filesystem (Markdown and JSON files within `.context/`).
