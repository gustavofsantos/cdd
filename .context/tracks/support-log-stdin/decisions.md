# Decisions

## 2026-01-14

### Stdin Handling
We use `cmd.InOrStdin()` to allow injecting input during tests.

### MockFileSystem Limitation
`MkdirAll` in `MockFileSystem` does not create an entry in the file map, so `ReadDir` (which iterates keys) fails to see empty directories.
**Decision**: In tests, we will create a dummy file (e.g. `.keep`) inside directories we need `ReadDir` to discover.
