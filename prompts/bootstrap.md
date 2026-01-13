# Agent Protocol
**Context ID:** `setup`
**Role:** Domain Cartographer

I have just run `cdd init`. Your goal is to identify the **Bounded Contexts** (Business Domains) within this project and populate the Global Context files. Focus on *high-level structure* to avoid getting lost in the noise of a large brownfield project.

## Protocol (Strictly Sequential)
Run `cdd recite setup` to confirm your next step.

### Phase 1: Domain Survey (The Map)
**Objective:** Identify the architectural style and candidate Bounded Contexts.
1.  **Topography:** Run `ls -F` on the root.
    * *Ignore* generic noise: `scripts/`, `bin/`, `dist/`, `build/`, `.config/`, `node_modules/`, `vendor/`.
    * *Focus* on source roots: `src/`, `lib/`, `pkg/`, `internal/`, `services/`, `apps/`.
2.  **Workspace Detection:** Read configuration files (`package.json`, `go.work`, `pom.xml`, `lerna.json`) to detect Monorepo or Multi-Module structures.
3.  **Architecture Inference:**
    * *Layered (Technical Grouping):* Do you see `controllers/`, `models/`, `views/`? -> **Contexts are likely hidden inside these folders.**
    * *Modular (Domain Grouping):* Do you see `billing/`, `auth/`, `shipping/`? -> **These ARE the Contexts.**
4.  **Drafting:** Create a *provisional* `product.md` and `tech-stack.md`. List your candidate Bounded Contexts clearly.

### Phase 2: Domain Alignment (The Handshake)
**Objective:** Validate the business domains with the user.
1.  **Report:** Present the "Domain Map":
    * "Architecture Style: [Monolith/Microservices/Layered/Modular]"
    * "Detected Bounded Contexts: [List of potential domains, e.g., 'Auth', 'Payment', 'Legacy-Core']"
2.  **The Ask:** Ask the user:
    * "Do these Bounded Contexts accurately represent your business domains?"
    * "**Select your Primary Focus:** Which Context are we working on today? (We will ignore the others to reduce noise)."

### Phase 3: Context Deep Dive
**Objective:** Map ONLY the user's chosen Context.
1.  **Focused Scan:** Scan only the directories relevant to the selected Bounded Context.
    * *If Layered:* You may need to look at `controllers/BillingController.ts` and `models/Billing.ts`.
    * *If Modular:* Scan `services/billing/`.
2.  **Pattern Discovery:**
    * Identify the testing strategy *for this specific context*.
    * Identify the specific database or API patterns *for this specific context*.
3.  **Document:** Update `workflow.md` and `patterns.md` with findings specific to this domain.

### Phase 4: Finalization
1.  Write the final content to `.context/`.
2.  Run `cdd archive setup`.
