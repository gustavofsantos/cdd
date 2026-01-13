# Technology Stack & Constraints

## 1. Core Foundations
| Category | Technology | Version / Note |
| :--- | :--- | :--- |
| **Language** | Go | 1.24.3 |
| **Framework** | Cobra | 1.10.2 (CLI) |
| **Database** | File System | N/A (Markdown/JSON) |

## 2. Libraries & Tools
* **Styling/UI:** `charmbracelet/glamour` (Markdown rendering), `lipgloss` (Styling)
* **CLI Interface:** `spf13/cobra` (Commands), `spf13/pflag` (Flags)
* **Testing:** Standard `testing` package
* **Utils:** `text/template` (Template rendering)

## 3. Coding Standards & Style
* **Idiomatic Go:** Follow standard Go conventions (Effective Go)
* **Error Handling:** Explicit error returns, no panics (except startup)
* **Simplicity:** Prefer standard library where possible
* **Testing:** Table-driven tests for logic, separate integration tests where needed

## 4. Project Structure (Tree)

```tree
/cmd
  /cdd              # Main entry point (main.go)
/internal
  /cmd              # Command implementations
    /templates      # Text templates for generation
  /platform         # Interfaces and abstractions (FS)
/prompts            # Embedded AI prompts
/.context           # Global context store
```
