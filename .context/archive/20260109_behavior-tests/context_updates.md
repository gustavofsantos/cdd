# Proposed Global Context Updates
> Add notes here if product.md or tech-stack.md needs updating.

- Refactored CLI commands to use `FileSystem` interface (Dependency Injection) for better testability.
- Defined `FileSystem` interface and `MockFileSystem` in `internal/platform`.
- Added comprehensive unit tests for `start`, `recite`, `list`, `log`, `archive`, `dump` commands.
