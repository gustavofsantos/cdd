# Lifecycle Specification

## Requirement: Project Initialization
The `cdd init` command MUST set up the `.context/` directory structure.

#### Scenario: Initializing a new project
- **Given** an empty directory
- **When** I run `cdd init`
- **Then** the directory `.context/` should be created
- **And** the global context files (`product.md`, `tech-stack.md`, `workflow.md`, `patterns.md`) should be initialized with templates.

## Requirement: Track Management
The system MUST support an isolated "Track" lifecycle using the "Pull Request Pattern".
- **Files:** Each track MUST consist of exactly three files: `spec.md`, `plan.md`, and `decisions.md`.
- **Delta Spec:** The `spec.md` file serves as a delta (added/modified requirements) to be merged into Global Specs.

#### Scenario: Starting a new track
- **Given** an initialized CDD environment
- **When** I run `cdd start feature-name`
- **Then** a directory `.context/tracks/feature-name/` should be created
- **And** it should contain `spec.md`, `plan.md`, and `decisions.md` templates.

#### Scenario: Archiving a track
- **Given** an active track that has been integrated
- **When** I run `cdd archive feature-name`
- **Then** the track directory should be moved to `.context/archive/`.

## Requirement: Context Visibility
The system MUST provide commands to view and manage context.
- **`recite`**: Display `spec.md` and `plan.md` for a track.
- **`list`**: List all active tracks.
- **`log`**: Append decisions to a track's `decisions.md`.
- **`dump`**: Capture stdin into a track's `scratchpad.md`.
