# Plan for cursor-rules-support

## Phase 0: Analysis
- [x] 🗣️ Phase 0: Alignment & Requirements (Fill `spec.md` using EARS)

## Phase 1: Architecture
- [x] 📝 Phase 1: Approval (User signs off)

## Phase 2: Implementation (Completed)

### Completion
- All 7 TDD tasks completed
- All tests passing
- Manual testing verified cursor installation works correctly
- Help text updated with cursor examples

## Phase 3: Integration

### Design
1. **installCursorRules()** function: Similar to `installSkill()` but writes concatenated content to `.cursorrules` file
2. **buildCursorRulesContent()** function: Merges all 5 skills into single file with version metadata
3. **Extract version** from combined content (use max version or new composite version)
4. Support idempotent updates with version checking and backup
5. Update help text and flag validation to accept "cursor" target

### Tasks
- [x] 1. Extract version extraction logic into reusable function
- [x] 2. Create buildCursorRulesContent() to concatenate skills with markdown structure
- [x] 3. Implement installCursorRules() function with version checking and backup logic
- [x] 4. Update Run handler to support cursor target and call installCursorRules
- [x] 5. Update help text to document cursor target
- [x] 6. Add test for cursor installation
- [x] 7. Add test for cursor version checking and idempotency
