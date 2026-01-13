# Specification: commit-often

## 1. User Intent (The Goal)
Implement a "commit often" approach in the system prompt. The idea is to track every significant change with a commit and link the commit hash in the plan, ensuring that the information and code changes can be retrieved in the future.

## 2. Relevant Context (The Files)
- `prompts/system.md`: The primary system prompt for the agent.
- `.context/workflow.md`: Documentation of the development flow.
- `GEMINI.md`: Contains the agent system prompt context.

## 3. Context Analysis (Agent Findings)
- Current system prompt focuses on TDD and CDD but lacks explicit instructions for frequent commits linked in the plan.
- The workflow documentation describes the CDD cycle but doesn't mention linking commit hashes in `plan.md`.
- Proposed Changes:
    - Add a "Commit Often" mandate to `prompts/system.md`.
    - Update the TDD loop in `prompts/system.md` to include a commit step with a hash link requirement.
    - Update `.context/workflow.md` to reflect this new practice.

## 4. Test Reference

- `prompts/system_test.go`: `TestSystemPromptHasCommitOften` verifies the presence of the "Commit Often" mandate.
