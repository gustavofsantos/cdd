package cmd_test

import (
	"bytes"
	"testing"

	"cdd/internal/cmd"
	"cdd/internal/platform"
)

func TestReciteCmd_DisplaysPlan(t *testing.T) {
	fs := platform.NewMockFileSystem()
	// Setup plan
	planContent := "# My Plan\n- [ ] Task 1"
	_ = fs.WriteFile(".context/tracks/active-track/plan.md", []byte(planContent), 0644)

	command := cmd.NewReciteCmd(fs)
	buf := new(bytes.Buffer)
	command.SetOut(buf)
	command.SetErr(buf)

	command.SetArgs([]string{"active-track"})
	err := command.Execute()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	output := buf.String()
	expectedHeader := "=== RECITATION: active-track ===\n"
	if !contains(output, expectedHeader) {
		t.Errorf("expected output to contain '%s', got '%s'", expectedHeader, output)
	}
	if !contains(output, planContent) {
		t.Errorf("expected output to contain plan content, got '%s'", output)
	}
}

func TestReciteCmd_InfersTrackWhenOnlyOneExists(t *testing.T) {
	fs := platform.NewMockFileSystem()
	// Setup single track
	planContent := "# Single Track Plan\n- [ ] Task 1"
	_ = fs.WriteFile(".context/tracks/only-track/plan.md", []byte(planContent), 0644)

	command := cmd.NewReciteCmd(fs)
	buf := new(bytes.Buffer)
	command.SetOut(buf)
	command.SetErr(buf)

	// Execute without track name argument
	command.SetArgs([]string{})
	err := command.Execute()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	output := buf.String()
	expectedHeader := "=== RECITATION: only-track ===\n"
	if !contains(output, expectedHeader) {
		t.Errorf("expected output to contain '%s', got '%s'", expectedHeader, output)
	}
	if !contains(output, planContent) {
		t.Errorf("expected output to contain plan content, got '%s'", output)
	}
}

func TestReciteCmd_RequiresTrackWhenMultipleExist(t *testing.T) {
	fs := platform.NewMockFileSystem()
	// Setup multiple tracks
	_ = fs.WriteFile(".context/tracks/track-one/plan.md", []byte("# Plan 1"), 0644)
	_ = fs.WriteFile(".context/tracks/track-two/plan.md", []byte("# Plan 2"), 0644)

	command := cmd.NewReciteCmd(fs)
	buf := new(bytes.Buffer)
	command.SetOut(buf)
	command.SetErr(buf)

	// Execute without track name argument
	command.SetArgs([]string{})
	err := command.Execute()

	if err == nil {
		t.Fatalf("expected error when multiple tracks exist without specifying one")
	}

	output := buf.String()
	if !contains(output, "Error: multiple active tracks found") {
		t.Errorf("expected error message about multiple tracks, got '%s'", output)
	}
	if !contains(output, "Select a task:") {
		t.Errorf("expected menu to be displayed, got '%s'", output)
	}
}

func TestReciteCmd_ErrorWhenNoTracksExist(t *testing.T) {
	fs := platform.NewMockFileSystem()

	command := cmd.NewReciteCmd(fs)
	buf := new(bytes.Buffer)
	command.SetOut(buf)
	command.SetErr(buf)

	// Execute without track name argument
	command.SetArgs([]string{})
	err := command.Execute()

	if err == nil {
		t.Fatalf("expected error when no tracks exist")
	}

	output := buf.String()
	if !contains(output, "Error: No active tracks found") {
		t.Errorf("expected error message about no tracks, got '%s'", output)
	}
}

func TestReciteCmd_Help(t *testing.T) {
	fs := platform.NewMockFileSystem()
	command := cmd.NewReciteCmd(fs)
	buf := new(bytes.Buffer)
	command.SetOut(buf)
	command.SetErr(buf)

	command.SetArgs([]string{"--help"})
	err := command.Execute()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	output := buf.String()
	expected := "Usage:\n  recite [track-name] [flags]"
	if !contains(output, expected) {
		t.Errorf("expected help output to contain usage, got %s", output)
	}

	if !contains(output, "EXAMPLES:") {
		t.Errorf("expected help output to contain EXAMPLES section")
	}
}

func TestReciteCmd_DisplaysOnlyNextChunk(t *testing.T) {
	fs := platform.NewMockFileSystem()
	planContent := `# Plan for upgrade-recite

## Phase 0: Analysis
- [x] Done item

## Phase 1: Architecture
- [ ] Next item
- [ ] Another next item

## Phase 2: Implementation
- [ ] Future item`

	_ = fs.WriteFile(".context/tracks/test-track/plan.md", []byte(planContent), 0644)

	command := cmd.NewReciteCmd(fs)
	buf := new(bytes.Buffer)
	command.SetOut(buf)
	command.SetErr(buf)

	command.SetArgs([]string{"test-track"})
	err := command.Execute()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	output := buf.String()

	// Should NOT contain Phase 0
	if contains(output, "## Phase 0: Analysis") {
		t.Errorf("expected output NOT to contain Phase 0, but it did")
	}

	// Should contain Phase 1
	if !contains(output, "## Phase 1: Architecture") {
		t.Errorf("expected output to contain Phase 1, but it didn't")
	}
	if !contains(output, "- [ ] Next item") {
		t.Errorf("expected output to contain 'Next item'")
	}

	// Should NOT contain Phase 2 (optional, but requested "only display the next chunk")
	if contains(output, "## Phase 2: Implementation") {
		t.Errorf("expected output NOT to contain Phase 2, but it did")
	}
}

func contains(s, substr string) bool {
	return bytes.Contains([]byte(s), []byte(substr))
}
