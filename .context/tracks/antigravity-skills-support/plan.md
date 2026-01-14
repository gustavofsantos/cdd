# Plan for antigravity-skills-support

## Phase 0: Analysis
- [x] 🗣️ Phase 0: Alignment & Requirements (Fill `spec.md` using EARS)

## Phase 1: Architecture
- [x] 📝 Phase 1: Approval (User signs off)

## Phase 2: Implementation

### Task 1: Add "antigravity" to target validation
Add "antigravity" as a valid `--target` option in the agents command.
- [x] Update the switch statement in `NewAgentsCmd` to handle `--target antigravity`
- [x] Add "antigravity" to help text and example commands
- [x] Test that unknown targets still warn appropriately

### Task 2: Create Antigravity skill formatter helper
Create a function to convert CDD skill content to Antigravity format (with YAML frontmatter).
- [ ] Extract YAML frontmatter from existing skill content
- [ ] Add `name` and `description` fields if missing
- [ ] Validate required Antigravity fields are present
- [ ] Test formatter with all five CDD skills

### Task 3: Implement installAntigravitySkill function
Create a dedicated installation function for Antigravity target (similar to `installSkill` but with Antigravity format validation).
- [ ] Create `installAntigravitySkill` function in agents.go
- [ ] Create `.agent/` directory if it doesn't exist
- [ ] Write skill to `.agent/skills/{skill-id}/SKILL.md` with proper formatting
- [ ] Handle existing skill overwrite scenarios with user feedback
- [ ] Test installation creates correct directory structure

### Task 4: Wire Antigravity target to installation pipeline
Connect the new Antigravity handler to the agents command flow.
- [ ] Update the switch statement to call `installAntigravitySkill` when target is "antigravity"
- [ ] Ensure all five skills are installed
- [ ] Test end-to-end: `cdd agents --install --target antigravity`
- [ ] Verify output messages and file structure

### Task 5: Integration and documentation
Finalize and document the feature.
- [ ] Update help text with antigravity target example
- [ ] Add version metadata handling for Antigravity skills
- [ ] Create a test that verifies skills can be discovered by Antigravity after installation
- [ ] Document the feature in README or CONTRIBUTING
