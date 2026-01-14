# Plan for view-command-tab-autocompletion

## Phase 0: Analysis
- [x] 🗣️ Phase 0: Alignment & Requirements (Fill `spec.md` using EARS)

## Phase 1: Architecture
- [x] 📝 Phase 1: Approval (User signs off)

## Phase 2: Implementation

### Autocompletion Infrastructure
- [x] Create a shell completion handler function that detects active tasks
- [x] Implement logic to count active tasks and return their list
- [x] Add test coverage for single active task scenario
- [x] Add test coverage for multiple active tasks scenario
- [x] Add test coverage for no active tasks scenario

### View Command Integration
- [x] Integrate completion handler into the `view` command's completion logic
- [x] Wire up tab completion to call the autocompletion handler
- [x] Test end-to-end tab completion with single active task
- [x] Test end-to-end tab completion with multiple active tasks
- [x] Test end-to-end tab completion with no active tasks

### Selection Menu (Multiple Tasks)
- [x] Implement selection menu rendering for multiple tasks
- [x] Add user selection handling logic
- [x] Test menu presentation and selection flow
