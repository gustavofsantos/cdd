# Prompt Specification

## Requirement: Specialized AI Roles
The system MUST provide dedicated prompts for different AI agent roles.

#### Scenario: Accessing the Strategist (System) Prompt
- **When** I run `cdd prompts --system`
- **Then** the system should output the Markdown content of the Strategist role prompt (The Orchestrator).

#### Scenario: Accessing the Planner Prompt
- **When** I run `cdd prompts --planner`
- **Then** the system should output the Markdown content of the Planner role prompt designed for creating Delta Specs.

#### Scenario: Accessing the Executor Prompt
- **When** I run `cdd prompts --executor`
- **Then** the system should output the Markdown content of the Executor role prompt designed for TDD implementation.

#### Scenario: Accessing the Integrator Prompt
- **When** I run `cdd prompts --integrator`
- **Then** the system should output the Markdown content of the Integrator role prompt designed for merging Deltas into Global Specs.

## Requirement: Context Gardening (Deprecated)
The manual consolidation of the Context Inbox is being replaced by the automated Integration of Delta Specs.

## Requirement: Promotion Logic
Prompts MUST be embedded into the binary for portability and consistency.
- **Location:** Files reside in `prompts/*.md` at the project root.
- **Mechanism:** Go `embed` package is used in `prompts/prompts.go`.
