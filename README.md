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
By providing a "file-based extended memory" and a rigorous step-by-step workflow, CDD enables **smaller, faster, and cheaper models** (like Gemini 2.0 Flash or Claude 3.5 Haiku) to perform at a level usually reserved for much larger models. It’s about process over brute force.

### Lineage & Evolution
This project carries the DNA of specifications-driven workflows like **OpenSpec**, **Conductor**, and **Manus**. It is heavily inspired by the planning and execution patterns of industry leaders like **Cursor**, **Windsurf**, **Claude Code**, and **Antigravity**, but takes a unique direction focused on extreme context isolation and cost-efficiency.

## Target Audience

- **Experienced Engineers**: Those who want to leverage AI without abdicated their role as the project architect.
- **Context-Engineers**: Developers who believe that well-structured context is the key to reliable AI outputs.
- **Cost-Conscious Teams**: Teams looking to maximize their AI ROI by getting the most out of efficient, low-latency models.
- **Legacy Navigators**: Anyone working in "noisy" brownfield projects where AI traditionally struggles to stay aligned.

## Getting Started

To get started with the CDD Tool Suite, please follow the [Installation Instructions](INSTALLATION.md).

After installation, initialize the CDD environment in your project root:

```bash
cdd init
```

This creates the `.context` directory structure and starts a `setup` track.

## Usage

### Managing Tracks

Start a new track for a task:

```bash
cdd start <track-name>
```

This creates a new workspace in `.context/tracks/<track-name>` with template files for specification, plan, decisions, and scratchpad.

### Working in a Track

- **View Plan**: `cdd recite <track-name>` - Displays the current plan.
- **Log Decision**: `cdd log <track-name> "Decision message"` - Records a decision.
- **Dump Output**: `command | cdd dump <track-name>` - Pipes output to the track's scratchpad.

### Completing a Track

Archive a completed track:

```bash
cdd archive <track-name>
```

This moves the track to `.context/archive`, promotes the specification to `.context/features`, and updates the inbox with context updates.

### Viewing Status

- **List Tracks**: `cdd list` - Lists active tracks.
- **List Archived**: `cdd list --archived` - Lists archived tracks.
- **View Dashboard**: `cdd view` - Shows a project dashboard (requires Glow).
- **View Track**: `cdd view <track-name>` - Shows details for a specific track (requires Glow).

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
