# Lifecycle Specification

## Requirement: Project Initialization
The `cdd init` command MUST set up the `.context/` directory structure.

#### Scenario: Initializing a new project
- **Given** an empty directory
- **When** I run `cdd init`
- **Then** the directory `.context/` should be created
- **And** the global context files (`product.md`, `tech-stack.md`, `workflow.md`, `patterns.md`) should be initialized with templates.

## Requirement: Track Management
The system MUST support an isolated "Track" lifecycle.

#### Scenario: Starting a new track
- **Given** an initialized CDD environment
- **When** I run `cdd start feature-name`
- **Then** a directory `.context/tracks/feature-name/` should be created
- **And** it should contain `spec.md` and `plan.md` templates.

#### Scenario: Archiving a track
- **Given** an active track with updates
- **When** I run `cdd archive feature-name`
- **Then** the track directory should be moved to `.context/archive/`
- **And** any recorded context updates should be appended to `.context/inbox.md`.

## Requirement: Context Visibility
The system MUST provide commands to view and manage context.
- **`recite`**: Display `spec.md` and `plan.md` for a track.
- **`list`**: List all active tracks.
- **`log`**: Append decisions to a track's `decisions.md`.
- **`dump`**: Capture stdin into a track's `scratchpad.md`.
