# Context-Driven Development (CDD) Tool Suite

The CDD Tool Suite is a CLI application designed to facilitate Context-Driven Development. It helps developers and AI agents manage context, track progress, and maintain a history of decisions and changes through a file-based state management system.

## Features

- **Context Management**: Organized structure for project context (`.context/` directory).
- **Track Isolation**: Work on specific tasks in isolated "tracks" to prevent context pollution.
- **Workflow Automation**: Commands to start, archive, and manage tracks.
- **Documentation Integration**: Promotes knowledge from tracks to living documentation.
- **AI-Ready**: Designed to work seamlessly with AI agents by providing structured context.

## Why CDD? (The Rationale)

In the rapidly evolving world of AI-assisted programming, noise is the enemy of efficiency. While modern AI models can write excellent code, they often get "lost" when navigating large, legacy, or high-entropy projects. Furthermore, the operational cost of using massive flagship models for every task can be prohibitive—especially in regions where API costs are a significant factor.

**Context-Driven Development (CDD)** is an attempt to solve these challenges through **Context Engineering**.

### The Strategist and the Tactician
CDD is built on a fundamental power dynamic: **The Engineer is the Strategist; the AI is the Tactician.** 
- You remain in charge, commanding the changes and ensuring the strategy is sound. 
- The AI executes the tactics, guided by a strict protocol that prevents it from drifting off-course.

### Premium Value from Budget Models
By providing a "file-based extended memory" and a rigorous step-by-step workflow, CDD enables **small models** (like Gemini Flash or Claude Haiku) to perform at a level usually reserved for much larger models. It’s about process over brute force.

### Lineage & Evolution
This project carries the DNA of specifications-driven workflows like **OpenSpec**, **Conductor**, and **Manus**. It is heavily inspired by the planning and execution patterns of industry leaders like **Cursor**, **Windsurf**, **Claude Code**, and **Antigravity**, but takes a unique direction focused on extreme context isolation and cost-efficiency.

## Target Audience

- **Experienced Engineers**: Those who want to leverage AI without abdicated their role as the project architect.
- **Context-Engineers**: Developers who believe that well-structured context is the key to reliable AI outputs.
- **Cost-Conscious Teams**: Teams looking to maximize their AI ROI by getting the most out of efficient, low-latency models.
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
cdd init --system-prompt
```

Copy the output of this command and set it as the "System Prompt" or "Custom Instructions" for your AI chat session.

### 3. The Setup Track
Now that your environment is initialized, you are ready to start. The `setup` track is designed to help the AI map your project and create the initial context files.

**Get the Bootstrap Prompt:**

```bash
cdd init --bootstrap-prompt
```

Copy the output and paste it as your first message to the AI. This will instruct the AI (Role: Principal Archaeologist) to begin the setup process.

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
  Moves the track to `.context/archive` and promotes the specification.

### Working in a Track

- **View Plan**:
  ```bash
  cdd recite <track-name>
  ```
  Displays the current status of the plan. Always run this before asking the AI for the next step.

- **Log Decision**:
  ```bash
  cdd log <track-name> "Decision message"
  ```
  Records a decision in `decisions.md`.

- **Dump Output**:
  ```bash
  command | cdd dump <track-name>
  ```
  Pipes the output of a command directly to `scratchpad.md` so the AI can read it.

### Prompt Commands

You can retrieve the core prompts using the `init` command flags:

- **System Prompt** (for the AI Agent):
  ```bash
  cdd init --system-prompt
  ```
- **Bootstrap Prompt** (to start the `setup` track):
  ```bash
  cdd init --bootstrap-prompt
  ```
- **Inbox Prompt** (to process context updates):
  ```bash
  cdd init --inbox-prompt
  ```
- **Executor Prompt** (for task execution):
  ```bash
  cdd init --executor-prompt
  ```
- **Planner Prompt** (for high-level planning):
  ```bash
  cdd init --planner-prompt
  ```

### Viewing Status

- **List Tracks**: `cdd list` - Lists active tracks.
- **List Archived**: `cdd list --archived` - Lists archived tracks.
- **View Dashboard**: `cdd view` - Shows a project dashboard.
- **View Track**: `cdd view <track-name>` - Shows details for a specific track.

## Project Structure

The tool manages a `.context` directory with the following structure:

- `tracks/`: Active workspaces.
- `archive/`: Completed and archived workspaces.
- `features/`: Living documentation (promoted specifications).
- `product.md`: Product context.
- `tech-stack.md`: Technology stack context.
- `workflow.md`: Workflow context.
- `patterns.md`: Project patterns.
- `inbox.md`: Pending context updates.
