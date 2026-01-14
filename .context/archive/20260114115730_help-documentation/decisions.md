# Implementation Journal
> Created Wed Jan 14 10:55:12 -03 2026

## Technical Architecture

### System Components
- **Cobra Commands**: Every command in `internal/cmd/` was updated to include detailed `Long` descriptions and `EXAMPLES`.
- **Unit Tests**: New help output verification tests were added to ensure documentation persists and meets requirements.

### Integration Points
- **Go Help System**: Leveraging Cobra's built-in help system (`--help` flag and `help` command).

## Implementation Considerations

### Approach
- Followed the TDD loop for every command.
- Standardized the help format across all commands: Description -> Behavior -> Flags -> Examples.
- Used `rootCmd` in tests for global commands established via `init()` functions to correctly capture subcommand help.

### Architectural Decision Records (ADRs)

## ADR-001: Standardized Help Format
**Context**: Commands had inconsistent and minimal documentation.
**Decision**: Every command must include a `Long` description and an `EXAMPLES` section.
**Consequences**: Improved user experience and discoverability without requiring external documentation.
