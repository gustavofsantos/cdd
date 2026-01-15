# Track: surveyor-skill

## 1. User Intent
Add tests for the surveyor prompt and ensure the `agents --install` command produces the surveyor skill along with existing skills. The surveyor prompt and skill infrastructure must be properly integrated and validated.

## 2. Relevant Context
- `prompts/prompts.go` - Existing prompt registration mechanism
- `prompts/*_test.go` - Existing test patterns
- `prompts/analyst.md` - Template for prompt structure
- `.agents/skills/cdd-analyst/SKILL.md` - Existing skill structure
- Command: `agents --install` - Skill installation integration point

## 3. Requirements (EARS)

- The system shall embed the surveyor prompt from `prompts/surveyor.md` during build compilation.
- The system shall register the `Surveyor` variable in the `prompts` package.
- When `agents --install` is executed, the system shall discover and install the surveyor skill from `.agents/skills/cdd-surveyor/SKILL.md`.
- The system shall include the surveyor prompt in prompt integration tests.
- The system shall validate that the surveyor prompt has required YAML frontmatter (name, description, metadata).
- The system shall ensure the surveyor skill file exists and contains valid SKILL.md format.
- Where tests are run, the system shall confirm all prompts including surveyor are properly registered.

Patterns:
    Ubiquitous: The <system> shall <response>
    Event-driven: When <trigger>, the <system> shall <response>
    State-driven: While <state>, the <system> shall <response>
    Unwanted: If <condition>, then the <system> shall <response>
    Optional: Where <feature>, the <system> shall <response>

<!--
Example:
- The system shall encrypt all data at rest.
- When the user clicks 'Submit', the system shall validate the payload.
- While the offline mode is active, the system shall queue all requests locally.
- If the API returns a 500 error, then the system shall retry up to 3 times.
- Where the 'Beta' flag is enabled, the system shall display the new dashboard.
--->