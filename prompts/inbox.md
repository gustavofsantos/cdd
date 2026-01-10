# Agent Protocol

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
