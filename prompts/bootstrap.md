# BOOTSTRAP PROTOCOL
**Context ID:** `initialization`
**Role:** Domain Cartographer
**Objective:** Replace the `[Bootstrap: ...]` placeholders in the Global Context files with concrete data derived from the codebase.

## 1. Topography (The Scan)
Perform a shallow scan of the codebase to build a mental map.
1.  **Run:** `ls -R src/` (or `ls -R lib/`, `ls -R app/` depending on language).
2.  **Read:** `package.json`, `go.mod`, `pom.xml`, or `requirements.txt` to identify the **Tech Stack**.
3.  **Read:** `README.md` (if exists) to understand the **Product Vision**.

## 2. Analysis (The Mapping)
Based on the file structure, deduce the following:
* **Architecture Style:** Is it a Monolith (layered folders), Microservices (separate repos/folders), or Modular Monolith (domain folders)?
* **Bounded Contexts:** Identify the high-level business domains (e.g., "Checkout", "Identity", "Inventory").
* **Ubiquitous Language:** Extract 3-5 key domain terms used in class/file names.

## 3. The Injection (The Output)
You will now edit the Global Context files. **DO NOT** rewrite the whole file. **ONLY** replace the blockquotes starting with `> [Bootstrap: ...]`.

### Target: `.context/product.md`
* **Core Value:** detailed summary of what this software does based on the README.
* **Ubiquitous Language:** Fill the table with the terms you found.

### Target: `.context/architecture.md`
* **High-Level Design:** Describe the system boundary (e.g., "A monolithic Rails app serving a React frontend").
* **Architectural Pattern:** State the inferred pattern (e.g., "MVC", "Hexagonal", "Clean Architecture").
* **Core Components:** List the top-level modules or bounded contexts found in step 2.

### Target: `.context/tech-stack.md`
* **Core Foundations:** Fill in Language, Framework, and Database versions found in config files.
* **Libraries:** List the major libraries used for Styling, State, and Testing.

## 4. Final Verification
Ask the user:
> "I have mapped the domain and populated the Global Context.
> **Detected Style:** [Style]
> **Detected Contexts:** [List]
>
> Please review `.context/product.md`, `.context/architecture.md`, and `.context/tech-stack.md`.
> Are these definitions correct before we begin the first Track?"
