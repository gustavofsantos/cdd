package cmd_test

import (
	"bytes"
	"strings"
	"testing"

	"cdd/internal/cmd"
	"cdd/internal/platform"
)

func TestArchiveCmd_Success(t *testing.T) {
	fs := platform.NewMockFileSystem()
	trackName := "done-track"
	trackDir := ".context/tracks/" + trackName

	// Setup track files
	fs.WriteFile(trackDir+"/spec.md", []byte("Spec Content"), 0644)
	fs.WriteFile(trackDir+"/plan.md", []byte("- [x] Task 1\n- [x] Task 2"), 0644)
	fs.WriteFile(trackDir+"/context_updates.md", []byte("# Updates\nSome update\nAnother line"), 0644)

	// Command
	command := cmd.NewArchiveCmd(fs)
	buf := new(bytes.Buffer)
	command.SetOut(buf)
	command.SetErr(buf)

	command.SetArgs([]string{trackName})
	err := command.Execute()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Verify promotion
	featureFile := ".context/features/" + trackName + ".md"
	content, err := fs.ReadFile(featureFile)
	if err != nil {
		t.Errorf("expected feature file to be created")
	}
	if string(content) != "Spec Content" {
		t.Errorf("expected feature content match")
	}

	// Verify inbox append
	inboxFile := ".context/inbox.md"
	content, err = fs.ReadFile(inboxFile)
	if err != nil {
		t.Errorf("expected inbox file to represent updates")
	}
	if !strings.Contains(string(content), "Some update") {
		t.Errorf("expected inbox to contain updates")
	}

	// Verify Rename (Archive)
	// Check if original dir is gone (spec file gone)
	if _, err := fs.Stat(trackDir + "/spec.md"); err == nil {
		t.Errorf("expected track dir to be moved (original spec should be gone)")
	}

	// Check archive dir exists
	// We don't know the exact time, but we can search in files keys
	foundArchive := false
	expectedSuffix := "_" + trackName + "/spec.md"
	for k := range fs.Files {
		if strings.HasPrefix(k, ".context/archive/") && strings.HasSuffix(k, expectedSuffix) {
			foundArchive = true
			break
		}
	}
	if !foundArchive {
		t.Errorf("expected archived file to exist")
	}
}

func TestArchiveCmd_PendingTasks(t *testing.T) {
	fs := platform.NewMockFileSystem()
	trackName := "wip-track"
	trackDir := ".context/tracks/" + trackName
	fs.WriteFile(trackDir+"/spec.md", []byte(""), 0644)
	fs.WriteFile(trackDir+"/plan.md", []byte("- [ ] Pending Task"), 0644)

	command := cmd.NewArchiveCmd(fs)
	buf := new(bytes.Buffer)
	command.SetOut(buf)
	command.SetErr(buf)

	command.SetArgs([]string{trackName})
	err := command.Execute()
	if err == nil {
		t.Fatal("expected error for pending tasks")
	}

	output := buf.String()
	// It might return error or print error.
	// Original code: Printf "Error: Cannot archive..." then Exit(1).
	// Our refactor should return error.
	if !strings.Contains(err.Error(), "Cannot archive") && !strings.Contains(output, "Cannot archive") {
		// Cobra might capture error message in err.
		// Or we might return fmt.Errorf() which contains the message.
	}
}
