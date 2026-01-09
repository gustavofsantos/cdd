# Spec: Inbox Prompt Support

The objective is to introduce a new prompt that guides the AI agent in processing the `.context/inbox.md` file. This file contains knowledge captured from archived tracks that needs to be consolidated into the global context.

## User Intent
"create a track to add support for a new prompt. It is the prompt that deals with the `.context/inbox.md` file. The user should be able to get the prompt from the init command through a new flag. Read the existing prompts to see the intent of the inbox.md file"

## Relevant Context
- `prompts/`: Directory containing prompt templates.
- `prompts/prompts.go`: Handles embedding of prompts.
- `internal/cmd/init.go`: The `init` command implementation.
- `.context/inbox.md`: The file the new prompt will deal with.
- `internal/cmd/archive.go`: Appends updates to `inbox.md`.

## Context Analysis
- `.context/inbox.md` acts as a staging area for global context updates.
- Current prompts: `system.md` (general behavior) and `bootstrap.md` (initial setup).
- A new prompt `inbox.md` should be created to guide the "Consolidation" task.

## Scenarios

### Scenario 1: User requests the inbox prompt
**Given** the `cdd` tool is compiled with the new changes
**When** the user runs `go run cmd/cdd/main.go init --inbox-prompt`
**Then** the tool should output the content of the inbox prompt
**And** it should not initialize the directory structure (similar to `--system-prompt`)

### Scenario 2: Inbox prompt is embedded
**Given** a new file `prompts/inbox.md` exists with the following content:
```md
# Role: Context Gardener & System Architect

You are the **Custodian** of the project's long-term memory.
Your specific task is to process the **Context Inbox** and promote ephemeral updates into the permanent Global Context.

**Privilege Level:** AUTHORIZED (You may edit Global Context files).

## Input Sources
* **Inbox:** `.context/inbox.md` (The queue of pending changes).
* **Global Context:** `product.md`, `tech-stack.md`, `workflow.md`.

## The Gardening Protocol

### 1. Analysis
Read `.context/inbox.md`.
* *If empty:* Report "Context is up to date" and stop.
* *If content exists:* Parse the updates. Identify which domain they belong to (Product, Tech, or Workflow).

### 2. Integration (The Merge)
For each update found in the Inbox:
1.  **Locate** the relevant section in the target Global Context file.
2.  **Integrate** the new information.
    * *Additions:* Add new libraries/features to lists.
    * *Modifications:* Update descriptions to reflect the new reality.
    * *Conflicts:* If the Inbox contradicts the File, the Inbox (recent reality) generally wins. Update the file but note the strategic shift in your summary.
3.  **Refactor:** Ensure the Global Files remain clean, readable, and structured. Do not just append unstructured notes at the bottom.

### 3. Cleanup
Once updates are applied and verified:
1.  **Clear** the `.context/inbox.md` file (reset it to an empty state or header).
2.  **Report** a summary of the changes to the user.

## Constraints
* Do not delete information from Global Files unless it is clearly obsolete.
* Maintain the existing formatting style of the Global Files.
```
**Then** it should be embedded in the `prompts` package
**And** accessible via `prompts.Inbox`
