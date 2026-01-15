# Pack Command Specification

## 1. Overview
The `pack` command compresses global specifications by extracting only paragraphs relevant to a given topic. It helps manage cognitive load in large projects by delivering focused context without requiring knowledge of entire specification files.

## 2. Requirements

### 2.1 Command Invocation
- The command shall be invoked as `cdd pack --focus <topic>` where `<topic>` is the search term.
- The `--focus` flag is required; the command shall return an error if it is missing.
- The `--raw` flag is optional; when provided, output shall be plain text without markdown rendering.

### 2.2 Specification Discovery
- The command shall discover all `.md` files in the `.context/specs/` directory.
- Each discovered file shall be loaded and parsed for its content.
- The system shall handle missing or inaccessible spec directories gracefully.

### 2.3 Paragraph Extraction
- The command shall split each specification into paragraphs (separated by blank lines).
- Each paragraph shall be treated as a distinct searchable unit.
- Metadata (file source, relevance score) shall accompany each paragraph.

### 2.4 Topic-Based Filtering
- When the user invokes `cdd pack --focus <topic>`, the system shall search all paragraphs using fuzzy matching.
- Fuzzy matching shall score paragraphs on a scale of 0.0 to 1.0, where 1.0 is a perfect match.
- By default, the system shall return paragraphs with a score of 0.5 or higher.
- Paragraphs shall be ranked by relevance score (highest first).

### 2.5 Output Formatting
- If the `--raw` flag is provided, output shall be plain text markdown without ANSI rendering codes.
- Otherwise, the system shall render markdown with appropriate formatting for terminal display.
- Output shall include a header with the search topic and match count.
- Each match shall display the score, source file, and paragraph content.
- Matches shall be clearly separated by visual dividers.

### 2.6 No-Match Handling
- If no paragraphs match the search topic, the system shall output a helpful message indicating this.
- The message shall suggest exploring different topics or checking available specifications.

### 2.7 Shell Completion
- The `pack` command shall support shell completion for common topics.
- Completion suggestions shall include topics such as: log, view, command, specification, requirement, authentication, authorization, tracking, decision, architecture, testing, deployment, configuration, error, validation.
- Completion shall be case-insensitive.
- Completion shall filter suggestions based on the user's partial input.

## 3. Relevant Files
- `internal/cmd/pack.go`: Implementation of the pack command
- `internal/cmd/pack_utils.go`: Core utilities (paragraph extraction, fuzzy matching, spec discovery, filtering)
- `internal/platform/fs.go`: Filesystem abstraction layer
- Tests:
  - `internal/cmd/pack_test.go`: Basic command functionality
  - `internal/cmd/pack_paragraph_test.go`: Paragraph extraction tests
  - `internal/cmd/pack_fuzzy_test.go`: Fuzzy matching accuracy tests
  - `internal/cmd/pack_specs_test.go`: Spec discovery tests
  - `internal/cmd/pack_filter_test.go`: Filtering and matching tests
  - `internal/cmd/pack_discovery_test.go`: Multiple file discovery tests
  - `internal/cmd/pack_filter_accuracy_test.go`: Cross-topic filtering tests
  - `internal/cmd/pack_integration_test.go`: Integration tests
  - `internal/cmd/pack_output_test.go`: Output formatting tests
  - `internal/cmd/pack_e2e_test.go`: End-to-end workflow tests
  - `internal/cmd/pack_completion_test.go`: Shell completion tests
