# Prompt Specification

## Requirement: Specialized AI Roles
The system MUST provide dedicated prompts for different AI agent roles.

#### Scenario: Accessing the Strategist (System) Prompt
- **When** I run `cdd prompts --system`
- **Then** the system should output the Markdown content of the Strategist role prompt.

#### Scenario: Accessing the Planner Prompt
- **When** I run `cdd prompts --planner`
- **Then** the system should output the Markdown content of the Planner role prompt designed for creating track plans.

#### Scenario: Accessing the Executor Prompt
- **When** I run `cdd prompts --executor`
- **Then** the system should output the Markdown content of the Executor role prompt designed for task implementation.

#### Scenario: Accessing the Calibration Prompt
- **When** I run `cdd prompts --calibration`
- **Then** the system should output the Markdown content of the Calibration prompt used for setting up new environments.

## Requirement: Context Gardening (Inbox Handling)
The system MUST provide a prompt specifically for consolidating the Context Inbox.

#### Scenario: Accessing the Inbox Prompt
- **When** I run `cdd prompts --inbox` (or legacy `cdd init --inbox-prompt`)
- **Then** the system should output instructions for a "Context Gardener" to merge ephemeral updates into Global Context.

## Requirement: Promotion Logic
Prompts MUST be embedded into the binary for portability and consistency.
- **Location:** Files reside in `prompts/*.md` at the project root.
- **Mechanism:** Go `embed` package is used in `prompts/prompts.go`.
