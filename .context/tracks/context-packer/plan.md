# Plan for context-packer

## Phase 0: Analysis
- [x] 🗣️ Phase 0: Alignment & Requirements (Fill `spec.md` using EARS)

## Phase 1: Architecture
- [x] 📝 Phase 1: Approval (User signs off)

## Phase 2: Implementation

### 2.1 Core Infrastructure
- [x] Implement paragraph extraction utility that splits markdown by blank lines
- [x] Implement fuzzy string matching function (e.g., using levenshtein distance or simple substring scoring)
- [x] Write tests for paragraph extraction with various markdown formats
- [x] Write tests for fuzzy matching accuracy

### 2.2 Spec Reader & Search Engine
- [x] Implement spec file discovery (read all .md files from `.context/specs/`)
- [x] Implement topic-based paragraph filtering using fuzzy matching
- [x] Write tests for spec discovery with multiple files
- [x] Write tests for filtering accuracy across different topics
- [x] Implement search result ranking (score matching paragraphs)

### 2.3 Pack Command (CLI Interface)
- [x] Create `pack.go` command file with Cobra command structure
- [x] Implement `--focus <topic>` flag parsing
- [x] Implement command logic to orchestrate spec reading and filtering
- [x] Write integration tests for `cdd pack --focus <topic>` invocation
- [x] Test error handling when no specs are found

### 2.4 Output Formatting
- [ ] Implement markdown rendering output (using existing glamour pattern)
- [ ] Implement `--raw` flag for plain text output
- [ ] Implement "no matches found" message
- [ ] Write tests for both output formats
- [ ] Add output tests for empty result cases

### 2.5 Integration & Polish
- [ ] Register pack command in root.go
- [ ] Add pack command to help documentation
- [ ] Write end-to-end test simulating real workflow
- [ ] Test with existing `.context/specs/` files
- [ ] Add shell completion support for pack command
