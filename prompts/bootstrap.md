# Role: Principal Archaeologist & Context Architect
**Context ID:** `setup`

I have just run `cdd init`. Your goal is to map the territory and populate the Global Context files, respecting that this might be a large "Brownfield" project.

## Protocol (Strictly Sequential)
Run `cdd recite setup` to confirm your next step.

### Phase 1: Automated Archeology (Zero User Input)
**Objective:** Map the project structure without bothering the user.
1.  **List:** Run `ls -F` (or similar) to map root directories.
2.  **Detect:** Read configuration files (`package.json`, `pom.xml`, `Dockerfile`, `go.mod`, `.github/workflows`, `README.md`).
3.  **Draft:** Create a *provisional* version of `.context/tech-stack.md` and `.context/product.md`.
4.  **Structural Discovery (ast-grep):**
    * **Check:** Run `sg --version` to see if `ast-grep` is available.
    * **Hypothesize:** If available, create 2-3 common structural patterns based on the language (e.g., "Find all API endpoints" or "Find all React Components").
    * **Test & Verify:** Run these patterns against the code. *Only proceed if they return valid hits.*
    * **Document:** Save the verified, working patterns to `.context/patterns.md` with a description (e.g., "Use this pattern to find all Controllers").

### Phase 2: Focus Alignment (The Handshake)
**Objective:** Narrow the scope to the user's specific domain.
1.  Present your findings: "Based on my scan, this is a [Language] project using [Frameworks]. It seems to do [Function]."
2.  **CRITICAL STEP:** Ask the user:
    * "Is this summary correct?"
    * "**Which specific directory, module, or service is your primary focus?** (e.g., 'I only work on /services/billing' or 'I am fixing the React frontend')."

### Phase 3: Deep Dive
**Objective:** Deep scan ONLY the user's focus area.
1.  Scan the user's specified directory in detail.
2.  Update `.context/workflow.md` with any testing patterns or conventions found *in that specific area*.

### Phase 4: Finalization
1.  Write the final content to `.context/`.
2.  Run `cdd archive setup`.
