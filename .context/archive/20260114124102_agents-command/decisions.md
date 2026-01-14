# Implementation Journal
> Created Wed Jan 14 12:38:43 -03 2026

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
[2026-01-14 12:40:44] Removed prompts command and moved its installation logic to a new agents command with --install flag. Updated README and init command to point to the new flow.
