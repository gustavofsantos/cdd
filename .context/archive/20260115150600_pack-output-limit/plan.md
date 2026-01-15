# Plan for pack-output-limit

## Phase 0: Analysis
- [x] 🗣️ Phase 0: Alignment & Requirements (Fill `spec.md` using EARS)

## Phase 1: Architecture
- [x] 📝 Phase 1: Approval (User signs off)

## Phase 2: Implementation

### 2.1 Core Limit Logic
- [x] Implement LimitResults utility function that truncates a slice of ParagraphMatch by count
- [x] Write tests for LimitResults with various limit values
- [x] Write tests for edge cases (limit=0, limit > available, negative limit)

### 2.2 Pack Command Flag Integration
- [x] Add `--limit` flag to NewPackCmd (optional, default no limit)
- [x] Parse limit value and validate (must be non-negative integer)
- [x] Write tests for flag parsing and validation

### 2.3 Output Logic Integration
- [x] Integrate LimitResults into runPackCmd after filtering but before output
- [x] Update output header to show if results were truncated
- [x] Write integration tests with limited results

### 2.4 Output Formatting
- [x] Update buildPackMarkdown to accept limit parameter
- [x] When limit=0, show only match count header (no paragraphs)
- [x] When truncated, add "(showing X of Y matches)" message
- [x] Write tests for truncation message formatting

### 2.5 Toolbox Integration
- [x] Update pack toolbox wrapper to accept limit parameter
- [x] Update AMP_TOOLBOX.md documentation with limit examples
- [x] Test toolbox describe action includes limit argument

### 2.6 Testing & Documentation
- [x] Write comprehensive tests for limit behavior across different scenarios
- [x] Update pack.md spec with new limit requirement
- [x] Test shell completion still works with new flag
- [x] Write end-to-end test with various limit values
