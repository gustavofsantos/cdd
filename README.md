# Context-Driven Development (CDD) Tool Suite

The CDD Tool Suite is a CLI application designed to facilitate Context-Driven Development. It helps developers and AI agents manage context, track progress, and maintain a history of decisions and changes through a file-based state management system.

## Features

- **Context Management**: Organized structure for project context (`.context/` directory).
- **Track Isolation**: Work on specific tasks in isolated "tracks" to prevent context pollution.
- **Workflow Automation**: Commands to start, archive, and manage tracks.
- **Documentation Integration**: Promotes knowledge from tracks to living documentation.
- **AI-Ready**: Designed to work seamlessly with AI agents by providing structured context.

## Why CDD? (The Rationale)

In the rapidly evolving world of AI-assisted programming, noise is the enemy of efficiency. While modern AI models can write excellent code, they often get "lost" when navigating large, legacy, or high-entropy projects. Furthermore, the operational cost of using massive flagship models for every task can be prohibitive.

**Context-Driven Development (CDD)** is an attempt to solve these challenges through **Context Engineering**.

### The Strategist and the Tactician
CDD is built on a fundamental power dynamic: **The Engineer is the Strategist; the AI is the Tactician.**
- You remain in charge, commanding the changes and ensuring the strategy is sound.
- The AI executes the tactics, guided by a strict protocol that prevents it from drifting off-course.

### The CDD State Machine
The workflow is governed by a strict state machine with four distinct phases. The AI adopts a specific persona for each phase:

1.  **Analyst (Phase 0)**: Interviews the user to understand the intent and frames the problem in a `spec.md`.
2.  **Architect (Phase 1)**: Designs the technical solution and creates a step-by-step TDD `plan.md`.
3.  **Executor (Phase 2+)**: Implements the plan using a Red/Green/Refactor loop, checking off items one by one.
4.  **Integrator**: Merges the finished work into the global context and cleans up the inbox.

### Premium Value from Budget Models
By providing a "file-based extended memory" and a rigorous step-by-step workflow, CDD enables **small models** (like Gemini Flash or Claude Haiku) to perform at a level usually reserved for much larger models. It’s about process over brute force.

### Lineage & Evolution
This project carries the DNA of specifications-driven workflows like **OpenSpec**, **Conductor**, and **Manus**. It is heavily inspired by the planning and execution patterns of industry leaders like **Cursor**, **Windsurf**, **Claude Code**, and **Antigravity**, but takes a unique direction focused on extreme context isolation and cost-efficiency.

## Target Audience

- **Experienced Engineers**: Those who want to leverage AI without abdicating their role as the project architect.
- **Context-Engineers**: Developers who believe that well-structured context is the key to reliable AI outputs.
- **Cost-Conscious Teams**: Teams looking to maximize their AI ROI.
- **Legacy Navigators**: Anyone working in "noisy" brownfield projects where AI traditionally struggles to stay aligned.

## Quick Start: Your First Task

This guide will walk you through your first interaction with the CDD Tool Suite, from installation to completing your first task (Project Setup).

### 1. Installation
Follow the [Installation Instructions](INSTALLATION.md) to get the `cdd` binary installed on your system.

### 2. Initialize & Configure Agent
Go to your project root and initialize the environment:

```bash
cdd init
```

This command sets up the `.context` directory and automatically starts the `setup` track.

**Setup your AI Agent:**
You need to provide your AI (LLM) with the **System Prompt**. This defines its role and the rules of engagement.

```bash
cdd prompts --system
```

Copy the output of this command and set it as the "System Prompt" or "Custom Instructions" for your AI chat session.

### 3. The Setup Track
Now that your environment is initialized, you are ready to start. The `setup` track is designed to help the AI map your project and create the initial context files.

**Get the Bootstrap Prompt:**

```bash
cdd prompts --bootstrap
```

Copy the output and paste it as your first message to the AI. This will instruct the AI to begin the setup process.

### 4. The Workflow Loop
The AI will now guide you through the **CDD Loop**. Your job is to be the conduit between the AI and the terminal.

1.  **Recite**: The AI will ask to see the plan. Run `cdd recite setup` and paste the output back to the AI.
2.  **Work**: The AI will generate commands or code. Execute them.
3.  **Log**: If a decision is made, run `cdd log setup "Decision details"`.
4.  **Repeat**: The AI will update the plan and ask for `recite` again.

### 5. Archiving
Once the setup is complete, the AI will instruct you to archive the track.

```bash
cdd archive setup
```

This moves the track to the archive and promotes the findings to the global context files (`product.md`, `tech-stack.md`, etc.). You are now ready to start your next task!

## Usage

### Managing Tracks

- **Start a new track**:
  ```bash
  cdd start <track-name>
  ```
  Creates a new workspace in `.context/tracks/<track-name>`.

- **Archive a track**:
  ```bash
  cdd archive <track-name>
  ```
  Moves the track to `.context/archive` and promotes the specification to the Inbox.

### Working in a Track

- **View Plan (for Agent)**:
  ```bash
  cdd recite <track-name>
  ```
  Displays the current status of the plan with specific instructions for the agent.

- **Log Decision**:
  ```bash
  cdd log <track-name> "Decision message"
  ```
  Records a decision in `decisions.md`.



### Prompt Commands

You can retrieve the core prompts using the dedicated `prompts` command:

- **System Prompt**: `cdd prompts --system`
- **Bootstrap Prompt**: `cdd prompts --bootstrap`
- **Inbox Prompt** (Context Gardener): `cdd prompts --inbox`
- **Executor Prompt**: `cdd prompts --executor`
- **Planner Prompt**: `cdd prompts --planner`
- **Calibration Prompt**: `cdd prompts --calibration`
- **Integrator Prompt**: `cdd prompts --integrator`

### Viewing Status

### Viewing Status / Dashboard

- **Dashboard**: `cdd view` 
  Shows active tracks. Use `--archived` to see history. Use `--inbox` to see pending updates.

- **Track Details**: `cdd view <track-name> [flags]`
  **Flags:**
  - `--spec`: Show the track specification.
  - `--plan`: Show the track plan.
  - `--log`: Show the decision log.
  - `--raw`: Output raw text (pipe-friendly).

## Project Structure

The tool manages a `.context` directory with the following structure:

- `tracks/`: Active workspaces (ephemeral).
- `archive/`: Completed and archived workspaces (history).
- `specs/`: Living documentation (promoted specifications, functional requirements).
- `product.md`: Product vision and global goals.
- `tech-stack.md`: Technology stack and constraints.
- `architecture.md`: High-level architecture and boundaries.
- `inbox.md`: Pending context updates waiting to be integrated.

