# Help Documentation Specification

## 1. Overview
The CDD Tool Suite must provide comprehensive inline documentation to ensure discoverability and ease of use for both human developers and AI agents. This is primarily achieved through the `--help` flag on the root command and all subcommands.

## 2. Requirements

### 2.1 Standardized Command Help
Every command in the CDD suite (root and subcommands) must implement:
- **Detailed Long Description**: A multi-paragraph explanation of the command's purpose, behavior, and position in the CDD workflow.
- **Usage Examples**: At least one concrete example of how to invoke the command.
- **Flag Documentation**: Clear explanations for all available flags.

### 2.2 Root Command Overview
The root `cdd` help must provide:
- A high-level overview of the Context-Driven Development (CDD) methodology.
- A list of core principles (Ephemeral Tracks, Eternal Specs, Gateway Inbox).
- A summarized workflow guide.

## 3. Relevant Context
- `internal/cmd/*.go`: All command implementations.
- `internal/cmd/root.go`: Root command definition.
