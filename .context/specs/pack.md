# Pack Command Specification

## 1. Overview
The `pack` command compresses global specifications by extracting only paragraphs relevant to a given topic. It helps manage cognitive load in large projects by delivering focused context without requiring knowledge of entire specification files.

## 2. Requirements

### 2.1 Command Invocation
- When invoked, the command shall require the `--focus <topic>` flag; otherwise, the system shall return an error.
- When the `--raw` flag is provided, the system shall output plain text without markdown rendering.
- When the `--limit <number>` flag is provided, the system shall constrain the maximum number of paragraphs returned (defaulting to -1 for no limit).

### 2.2 Specification Discovery
- The system shall discover all `.md` files in the `.context/specs/` directory.
- The system shall load and parse the content of each discovered file.
- The system shall handle missing or inaccessible spec directories gracefully.

### 2.3 Paragraph Extraction
- The system shall split each specification into paragraphs separated by blank lines.
- The system shall treat each paragraph as a distinct searchable unit.
- The system shall associate metadata (file source, relevance score) with each paragraph.

### 2.4 Topic-Based Filtering
- When `cdd pack --focus <topic>` is executed, the system shall search all paragraphs using fuzzy matching.
- The system shall score paragraphs on a scale of 0.0 to 1.0 (where 1.0 is a perfect match).
- The system shall return paragraphs with a score of 0.5 or higher by default.
- The system shall rank paragraphs by relevance score (highest first).

### 2.5 Output Formatting
- When the `--raw` flag is provided, the system shall output plain text markdown without ANSI rendering codes.
- When the `--raw` flag is not provided, the system shall render markdown with appropriate formatting for terminal display.
- The output shall include a header with the search topic and match count.
- The output shall display the score, source file, and paragraph content for each match.
- The system shall visually divide separate matches.

### 2.6 No-Match Handling
- If no paragraphs match the search topic, the system shall output a helpful message suggesting different topics or checking available specifications.

### 2.7 Output Limiting
- When `--limit N` is specified (where N >= 0), the system shall return at most N paragraphs.
- The system shall apply the limit after ranking by relevance.
- When `--limit 0` is specified, the system shall display only the match count header.
- When results are truncated, the system shall indicate "(showing X of Y)" in the header.
- When a negative limit is specified, the system shall return all matching paragraphs.

### 2.8 Shell Completion
- The `pack` command shall support shell completion for common topics (e.g., log, view, command, specification, requirement).
- The system shall perform case-insensitive completion.
- The system shall filter suggestions based on the user's partial input.
- The system shall support integer completion for the `--limit` flag.
