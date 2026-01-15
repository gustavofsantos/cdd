# Plan for amp-toolbox

## Phase 0: Analysis
- [x] 🗣️ Phase 0: Alignment & Requirements (Fill `spec.md` using EARS)

## Phase 1: Architecture
- [x] 📝 Phase 1: Approval (User signs off)

## Phase 2: Implementation

### Build Toolbox Wrapper System
- [x] Create a shared Go package for toolbox tool generation that builds a tool from a cobra command
- [x] Implement describe action: extract name, description, and args from cobra.Command and output JSON

### Build Individual Toolbox Executables
- [x] Create `cdd-init` wrapper that calls `cdd init` via the toolbox system
- [x] Create `cdd-start` wrapper that calls `cdd start <track>` with name argument
- [x] Create `cdd-recite` wrapper that calls `cdd recite <track>` with track name argument
- [x] Create `cdd-log` wrapper that calls `cdd log <track> <message>` with track and message arguments
- [x] Create `cdd-archive` wrapper that calls `cdd archive <track>` with track name argument
- [x] Create `cdd-view` wrapper that calls `cdd view [track]` with optional track name and flags
- [x] Create `cdd-agents` wrapper that calls `cdd agents --install` with install and target flags
- [x] Create `cdd-delete` wrapper that calls `cdd delete <track>` with track name argument
- [x] Create `cdd-version` wrapper that calls `cdd version` with no arguments

### Integration & Distribution
- [x] Update build system to compile all toolbox wrappers into a toolbox directory
- [x] Verify all wrappers output correct JSON on describe action
- [x] Verify all wrappers correctly parse stdin arguments on execute action
- [x] Document how to set up `AMP_TOOLBOX` environment variable to use the toolbox
