# Context-Driven Development (CDD) Tool Suite

The CDD Tool Suite is a CLI application designed to facilitate Context-Driven Development. It helps developers and AI agents manage context, track progress, and maintain a history of decisions and changes through a file-based state management system.

## The Rationale

In the rapidly evolving world of AI-assisted programming, noise is the enemy of efficiency. While modern AI models can write excellent code, they often get "lost" when navigating large, legacy, or high-entropy projects. Furthermore, the operational cost of using massive flagship models for every task can be prohibitive.

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
CDD uses **Agent Skills** to orchestrate the AI. Install the skill to your project:

```bash
cdd agents --install
```

This will create `.agent/skills/` with all CDD skills. If you are using **Antigravity**, you can install skills directly for Antigravity:

```bash
cdd agents --install --target antigravity
```

This creates `.agent/skills/` compatible with Antigravity's skill discovery. For other agents that support MCP/Skill protocol, use the default installation or specify the appropriate `--target` (claude, agents, etc.).

### 3. The Setup Track
Now that your environment is initialized, you are ready to start. The `setup` track is designed to help the AI map your project and create the initial context files.

**Start the AI mapping:**
Once the skill is installed, your agent will know how to proceed. Start by asking it to analyze the current state of the project.

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



### Agent Commands

Manage the AI agent integration:

- **Install CDD Skills**: `cdd agents --install [--target <target>]`
  
  Installs all five CDD skills (Orchestrator, Analyst, Architect, Executor, Integrator) to the specified target.
  
  **Available targets:**
  - `agent` (default): Creates `.agent/skills/` for local use
  - `antigravity`: Creates `.agent/skills/` compatible with Google Antigravity
  - `claude`: Creates `.claude/skills/` for Claude integration
  - `agents`: Creates `.agents/skills/` for generic agents
  - `cursor`: Creates `.cursorrules` for Cursor IDE

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

### Shell Completion (Tab Autocompletion)

The `cdd view` command supports tab autocompletion for track names. To enable this feature, you must set up shell completion:

```bash
# For bash
cdd completion bash | sudo tee /etc/bash_completion.d/cdd

# For zsh
cdd completion zsh | sudo tee /usr/share/zsh/site-functions/_cdd

# For fish
cdd completion fish | sudo tee /usr/share/fish/vendor_completions.d/cdd.fish
```

Or, to install for your user only:

```bash
# For bash (add to ~/.bashrc)
cdd completion bash | source

# For zsh (add to ~/.zshrc)
cdd completion zsh | source

# For fish (add to ~/.config/fish/config.fish)
cdd completion fish | source
```

After installation, reload your shell:

```bash
source ~/.bashrc   # or ~/.zshrc, depending on your shell
```

**Note:** Tab autocompletion only works after running the `cdd completion` command and sourcing the generated script in your shell. The completion function will suggest active task names when you type `cdd view <TAB>`.

## Project Structure

The tool manages a `.context` directory with the following structure:

- `tracks/`: Active workspaces (ephemeral).
- `archive/`: Completed and archived workspaces (history).
- `specs/`: Living documentation (promoted specifications, functional requirements).
- `product.md`: Product vision and global goals.
- `tech-stack.md`: Technology stack and constraints.
- `architecture.md`: High-level architecture and boundaries.
- `inbox.md`: Pending context updates waiting to be integrated.

