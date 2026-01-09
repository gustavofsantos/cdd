# Project Patterns

## Structural Patterns

### Cobra Commands
- **Description:** Definitions of CLI commands using `cobra.Command`.
- **Pattern (grep):** `grep -r "&cobra.Command" internal/`
- **Pattern (ast-grep):** `sg -p 'var $NAME = &cobra.Command{ $$$ }' -l go`

### Go Structs
- **Description:** Data structure definitions.
- **Pattern (grep):** `grep -r "type .* struct {" internal/`

### Error Handling
- **Description:** Standard error check block.
- **Pattern (grep):** `grep -r "if err != nil {" internal/`
