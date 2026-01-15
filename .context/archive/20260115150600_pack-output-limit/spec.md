# Track: pack-output-limit

## 1. User Intent
Add an output limit feature to the `cdd pack` command to prevent returning too many results that could overwhelm the context window or create cluttered output. The feature should allow users to specify a maximum number of results using a `--limit` flag, helping manage context efficiency when dealing with broad topic searches.

## 2. Relevant Context
- `.context/specs/pack.md`: Existing pack command specification
- `internal/cmd/pack.go`: Main pack command implementation
- `internal/cmd/pack_utils.go`: Filtering and result ranking logic
- `internal/cmd/pack_test.go`: Existing command tests
- `internal/cmd/pack_integration_test.go`: Integration test patterns
- The pack command currently returns all paragraphs matching score threshold (0.5+)
- Results are already ranked by relevance score (highest first)
- Output formatting already groups results by spec file

## 3. Requirements (EARS)
- Ubiquitous: The `pack` command shall accept an optional `--limit <number>` flag to constrain maximum results returned.
- Ubiquitous: When `--limit` is specified, the system shall return at most N paragraphs (where N is the limit value).
- Ubiquitous: The system shall apply the limit after ranking by relevance, returning the top N highest-scoring matches.
- Ubiquitous: The default behavior (when no `--limit` is specified) shall be to return all matching paragraphs (no limit).
- Event-driven: When the user specifies `--limit` that is lower than available matches, the output header shall indicate that results were truncated.
- Unwanted: If an invalid limit value is provided (non-numeric or negative), the system shall return an error with guidance.
- Optional: Where the `--limit` flag is provided with value 0, the system shall display only the match count without paragraph content.

Example:
- The system shall encrypt all data at rest.
- When the user clicks 'Submit', the system shall validate the payload.
- While the offline mode is active, the system shall queue all requests locally.
- If the API returns a 500 error, then the system shall retry up to 3 times.
- Where the 'Beta' flag is enabled, the system shall display the new dashboard.
