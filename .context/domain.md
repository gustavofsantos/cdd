# Domain Model

## 1. Ubiquitous Language
Definition: The specific vocabulary used by the team. These terms must be used precisely in code and specs.
Format: Term: Definition. Constraints or Invariants.

* **Track**: A temporary workspace for a specific unit of work. Contains `spec.md`, `plan.md`, `decisions.md` (and optional `current_state.md`). It is ephemeral and meant to be archived upon completion.
* **Context**: The collective state of the project managed by the tool, residing in the `.context` directory. It includes global files (`product.md`, `tech-stack.md`, `domain.md`) and the repository of `specs`.
* **Agent Skill**: A specialized persona or capability extension for the AI agent (e.g., Analyst, Architect). Skills are installed into `.agent/skills` (or similar targets).
* **CDD Loop**: The iterative process of Recite -> Work -> Log -> Repeat.
* **Recite**: The action of reading the current plan to determine the next immediate step.
* **Inbox**: A holding area (`.context/inbox.md`) for specifications from archived tracks that need to be processed into the permanent living documentation.
* **Living Documentation**: The set of specifications in `.context/specs/` that represent the current source of truth for system behavior.
* **Pack**: The process of searching and aggregating relevant context from the global state for AI consumption.

## 2. Domain Events (Event Storming)
**Definition:** Things that happen in the business that we care about.
Format: [Past Tense Verb], e.g., `OrderPlaced`, `PaymentFailed`.

* `TrackInitialized`: A new track has been created via `cdd start`.
* `StepRecited`: The agent or user has requested the next step in the plan.
* `DecisionLogged`: A technical or product decision has been recorded in the track's log.
* `SkillInstalled`: A new agent skill has been installed into the environment.
* `TrackArchived`: A track has been completed, its spec moved to the inbox, and the workspace cleaned up.
* `ContextPacked`: Parts of the global context have been searched and aggregated for the agent.

## 3. Bounded Contexts
**Definition:** The logical boundaries within the system.

* **Core Domain:** [Context Orchestration]
    * The CLI logic that manages the lifecycle of tracks (`start`, `recite`, `log`, `archive`).
    * The state machine driven by `plan.md`.
* **Supporting Subdomains:** [Agent Integration]
    * The mechanism for installing and managing Agent Skills (`agents` command).
    * Integration with external agent runtimes (Antigravity, Claude, etc.).
* **Generic Subdomains:** [File System & Templating]
    * Handling file operations, reading/writing markdown files.
    * Managing templates for tracks and skills.