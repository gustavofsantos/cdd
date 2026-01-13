# Product Vision & Strategy

## 1. Core Value Proposition
The CDD Tool Suite is a CLI application that facilitates Context-Driven Development through file-based state management. It enables developers and AI agents to work together efficiently by providing structured context, isolated workspaces ("tracks"), and a rigorous workflow protocol. The tool addresses the challenge of AI context pollution in large codebases by implementing extreme context isolation, allowing smaller AI models to perform at the level of flagship models through superior context engineering.

## 2. Target Audience / Personas
* **Experienced Engineers**: Developers who want to leverage AI without abdicating their role as the project architect.
* **Context-Engineers**: Developers who believe that well-structured context is the key to reliable AI outputs.
* **Cost-Conscious Teams**: Teams looking to maximize their AI ROI by getting the most out of efficient, low-latency models.
* **Legacy Navigators**: Anyone working in "noisy" brownfield projects where AI traditionally struggles to stay aligned.

## 3. Key Capabilities (The "What")
* **Track Management**: Create isolated workspaces for specific features/tasks to prevent context pollution
* **Workflow Automation**: Commands to start, archive, and manage development tracks
* **Prompt Orchestration**: Serve role-specific prompts (System, Bootstrap, Executor, Planner, Integrator) for AI agents
* **Context Engineering**: Maintain living documentation through structured specs and decision logs
* **Time Tracking**: Automatic capture of track lifecycle timestamps

## 4. Ubiquitous Language (Domain Glossary)
| Term | Definition |
| :--- | :--- |
| **Track** | An isolated workspace in `.context/tracks/<name>` for a specific feature or task. |
| **Spec** | The specification document (`spec.md`) defining requirements in EARS notation. |
| **Plan** | The task checklist and execution log (`plan.md`) showing workflow progress. |
| **Decisions** | The Architecture Decision Record journal (`decisions.md`) documenting key choices. |
| **Inbox** | The staging area (`.context/inbox.md`) for pending context updates awaiting integration. |
| **Bootstrap** | The initial project mapping phase where the AI scans and documents the codebase. |
| **Archive** | Completed tracks moved to `.context/archive/` with their specs promoted to global context. |

