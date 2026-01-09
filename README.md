# Context-Driven Development (CDD) Tool Suite

The CDD Tool Suite is a CLI application designed to facilitate Context-Driven Development. It helps developers and AI agents manage context, track progress, and maintain a history of decisions and changes through a file-based state management system.

## Features

- **Context Management**: Organized structure for project context (`.context/` directory).
- **Track Isolation**: Work on specific tasks in isolated "tracks" to prevent context pollution.
- **Workflow Automation**: Commands to start, archive, and manage tracks.
- **Documentation Integration**: Promotes knowledge from tracks to living documentation.
- **AI-Ready**: Designed to work seamlessly with AI agents by providing structured context.

## Installation

### Prerequisites

- Go 1.24 or higher
- [Glow](https://github.com/charmbracelet/glow) (optional, for `view` command)

### Installation via Binary

You can download the pre-compiled binary for your platform from the [Releases](https://github.com/Bitwise-Source/cdd/releases) page.

**Linux / macOS:**

```bash
# Example for Linux AMD64 (adjust URL for your platform/version)
curl -L -o cdd https://github.com/Bitwise-Source/cdd/releases/latest/download/cdd-linux-amd64
chmod +x cdd
sudo mv cdd /usr/local/bin/
```

### Build from Source

```bash
git clone https://github.com/Bitwise-Source/cdd.git
cd cdd
go build -o cdd cmd/cdd/main.go
```

You can then move the `cdd` binary to a directory in your system's PATH.

## Usage

### Initialization

Initialize the CDD environment in your project root:

```bash
cdd init
```

This creates the `.context` directory structure and starts a `setup` track.

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
