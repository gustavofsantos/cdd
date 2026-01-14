# Implementation Journal
> Created Wed Jan 14 12:23:11 -03 2026

## Technical Architecture

### System Components
<!-- Document the key components involved in this implementation -->

### Data Flow
<!-- Describe how data moves through the system -->

### Integration Points
<!-- List external systems, APIs, or modules this change interacts with -->

## Sequence Diagrams
<!-- Use d2 or ASCII diagrams to illustrate key interactions -->

```mermaid
shape: sequence_diagram
User -> API: Request
API -> User: Response
```

## Implementation Considerations

### Approach
<!-- Describe the overall approach taken to implement the spec -->

### Technical Constraints
<!-- List any technical limitations or constraints encountered -->

### Trade-offs
<!-- Document trade-offs made during implementation -->

### Performance Implications
<!-- Note any performance considerations or optimizations -->

### Security Considerations
<!-- Document security-related decisions or implications -->

## Architectural Decision Records (ADRs)

<!-- Use this section for significant architectural choices -->
<!-- Format: ## ADR-001: [Title] -->
<!-- Include: Context, Decision, Consequences -->
## ADR-001: Agent Skill Versioning Strategy

### Context
We need to pivot to using Agent Skills and support rapid iteration. The previous system had unversioned prompt files.

### Decision
- Adopt a `version: <int>` field in the `SKILL.md` frontmatter.
- Treat existing unversioned files as "Legacy" (v0/v1).
- Implement a "Backup and Replace" migration strategy: `SKILL.md` -> `SKILL.md.bak`.
- Hardcode the `CurrentSkillVersion` in the Go binary ensures the tool is the source of truth.

### Consequences
- Users running `cdd prompts --install` will automatically upgrade their prompt file.
- Customizations to `SKILL.md` will be moved to `.bak` and need manual reintegration if desired (safe default).
