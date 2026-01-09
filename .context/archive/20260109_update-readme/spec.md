# Specification: update-readme

## 1. User Intent (The Goal)
The user wants to refactor the project's `README.md` to make it more professional and "sellable".
Specifically:
1. Move installation details to a separate `INSTALLATION.md` file and link it in `README.md`.
2. Detail the rationale for the project and define the target audience, incorporating the user's specific philosophy on AI-assisted development.

## 2. Relevant Context (The Files)
- `README.md`: The main entry point for the project documentation.
- `INSTALLATION.md`: New file to be created.
- `GEMINI.md`: Contains the core CDD philosophy rules.

## 3. Context Analysis (Agent Findings)
- Current Behavior: `README.md` contains a significant section on installation (lines 13-55) which takes up space and might be better in a separate file for a "sales-focused" landing page.
- Current Behavior: `README.md` lacks the personal and practical rationale provided by the creator.
- Proposed Changes:
    - Create `INSTALLATION.md` and move the installation section there.
    - Update `README.md` with a "Why CDD?" (Rationale) section that highlights:
        - **The Strategist vs. The Tactician**: The user remains the strategist (the "boss"), while the AI acts as the tactician (the "worker").
        - **Context Engineering**: Solving the "getting lost" problem in large, noisy, legacy projects.
        - **Cost Efficiency**: Leveraging small, fast, and cheap models (Gemini Flash, Claude Haiku) through better process rather than brute-force reasoning.
        - **Lineage & Evolution**: Built on lessons learned from **OpenSpec**, **Conductor**, and **Manus**, and inspired by how tools like **Cursor**, **Windsurf**, **Claude Code**, and **Antigravity** handle planning. It explores a "different direction" by prioritizing strict context management to extract premium value from budget models.
    - Update `README.md` with an "Audience" section:
        - Experienced engineers navigating large-scale or brownfield projects.
        - Developers seeking to maximize AI ROI (especially in regions where API costs are high).
        - AI-assisted teams practicing XP/TDD.
    - Add a "Getting Started" section in `README.md` that links to `INSTALLATION.md` and describes `cdd init`.

## 4. Scenarios (Acceptance Criteria)
Feature: README Enhancement
  Scenario: Refactor Installation Instructions
    Given I have a `README.md` with installation details
    When I create `INSTALLATION.md` with those details
    And I replace the installation section in `README.md` with a link to `INSTALLATION.md`
    Then `README.md` should be cleaner and `INSTALLATION.md` should contain all setup steps.

  Scenario: Add Rationale and Audience
    Given the project needs to communicate its unique value proposition
    When I add a "Why CDD?" section detailing the strategist/tactician split and cost-efficiency
    And I add an "Audience" section for experienced engineers and cost-conscious developers
    Then the `README.md` should feel professional, compelling, and grounded in real-world engineering challenges.
