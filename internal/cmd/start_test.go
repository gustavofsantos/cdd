package cmd_test

import (
	"bytes"
	"strings"
	"testing"

	"cdd/internal/cmd"
	"cdd/internal/platform"
)

func TestStartCmd_CreatesTrack(t *testing.T) {
	fs := platform.NewMockFileSystem()
	command := cmd.NewStartCmd(fs)
	buf := new(bytes.Buffer)
	command.SetOut(buf)
	command.SetErr(buf)

	// Execute
	command.SetArgs([]string{"test-track"})
	err := command.Execute()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Verify Output
	output := buf.String()
	if !strings.Contains(output, "Track 'test-track' initialized.") {
		t.Errorf("expected output to contain 'Track 'test-track' initialized.', got '%s'", output)
	}

	// Verify File System
	if _, err := fs.Stat(".context/tracks/test-track/spec.md"); err != nil {
		t.Errorf("spec.md not created")
	}
	if _, err := fs.Stat(".context/tracks/test-track/plan.md"); err != nil {
		t.Errorf("plan.md not created")
	}

	// Verify Metadata
	metaPath := ".context/tracks/test-track/metadata.json"
	if _, err := fs.Stat(metaPath); err != nil {
		t.Errorf("metadata.json not created")
	} else {
		// Verify content
		content, _ := fs.ReadFile(metaPath)
		if !bytes.Contains(content, []byte("started_at")) {
			t.Errorf("metadata.json does not contain started_at")
		}
	}
}

func TestStartCmd_TrackExists(t *testing.T) {
	fs := platform.NewMockFileSystem()
	_ = fs.MkdirAll(".context/tracks/existing-track", 0755)
	// MockFS relies on file presence to detect directories or we need to update MkdirAll.
	// Simplest approach: create a file inside the directory.
	_ = fs.WriteFile(".context/tracks/existing-track/spec.md", []byte(""), 0644)

	command := cmd.NewStartCmd(fs)
	buf := new(bytes.Buffer)
	command.SetOut(buf)
	command.SetErr(buf)

	command.SetArgs([]string{"existing-track"})
	// Execute should fail?
	// If RunE returns error, command.Execute returns error.
	// The original code uses os.Exit(1). I will change it to return error.

	err := command.Execute()
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	// Check output contains error message?
	// Cobra prints error to Err.
	// Our refactored code should return error or print to cmd.ErrOrStderr.
	// Spec says: "Error: Track '%s' exists.\n"
}
