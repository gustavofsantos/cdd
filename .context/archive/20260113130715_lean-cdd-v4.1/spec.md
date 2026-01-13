# Track: lean-cdd-v4.1
**Target Spec:** `.context/specs/lifecycle/spec.md`, `.context/specs/prompts/spec.md`, `.context/specs/standards/spec.md`

## Context
Implementing Lean CDD v4.1 as requested by the user.

## Proposed Changes
### ADDED Requirements
* **Requirement: Pull Request Pattern**
    * The system SHALL use the `spec.md` file as a Delta Specification.
    * The Integrator SHALL apply changes defined in `spec.md` directly to Global Specs.
* **Requirement: 3-file Track Structure**
    * Tracks SHALL consist of exactly three files: `spec.md`, `plan.md`, and `decisions.md`.
    * #### Scenario: Starting a new track (v4.1)
        * When I run `cdd start <name>`
        * Then `spec.md`, `plan.md`, and `decisions.md` should be created.
        * And `scratchpad.md` or `context_updates.md` should NOT be created.

### MODIFIED Requirements
* **Requirement: Track Management** (in `specs/lifecycle/spec.md`)
    * (Updated to reflect the 3-file structure and removal of Inbox-based archival in favor of direct Spec mutation by the Integrator).
* **Requirement: Agent Roles** (in `specs/prompts/spec.md`)
    * Updated roles (Strategist, Planner, Executor, Integrator) to align with the Lean CDD v4.1 protocol.
* **Requirement: Development Workflow** (in `specs/standards/spec.md`)
    * 1. **Recite**: Load and review current plan/spec.
    * 2. **Spec**: Define the *Delta* (Additions/Modifications) in `spec.md`.
    * 3. **Plan**: Define TDD execution steps in `plan.md`.
    * 4. **Implement**: TDD execution (Red-Green-Refactor).
    * 5. **Integrate**: Apply the Delta to Global Specs and Archive.
