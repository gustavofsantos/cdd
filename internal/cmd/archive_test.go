package cmd_test

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"cdd/internal/cmd"
	"cdd/internal/platform"
)

func TestArchiveCmd_Success(t *testing.T) {
	fs := platform.NewMockFileSystem()
	trackName := "done-track"
	trackDir := ".context/tracks/" + trackName

	// Setup track files
	_ = fs.WriteFile(trackDir+"/spec.md", []byte("Spec Content"), 0644)
	_ = fs.WriteFile(trackDir+"/plan.md", []byte("- [x] Task 1\n- [x] Task 2"), 0644)
	_ = fs.WriteFile(trackDir+"/context_updates.md", []byte("# Updates\nSome update\nAnother line"), 0644)

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

func TestArchiveCmd_CleanupAndTimestamp(t *testing.T) {
	fs := platform.NewMockFileSystem()
	trackName := "cleanup-track"
	trackDir := ".context/tracks/" + trackName

	// Setup track files including scratchpad
	_ = fs.WriteFile(trackDir+"/spec.md", []byte("Spec"), 0644)
	_ = fs.WriteFile(trackDir+"/plan.md", []byte("- [x] Done"), 0644)
	_ = fs.WriteFile(trackDir+"/scratchpad.md", []byte("Ephemeral content"), 0644)

	command := cmd.NewArchiveCmd(fs)
	buf := new(bytes.Buffer)
	command.SetOut(buf)

	command.SetArgs([]string{trackName})
	err := command.Execute()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Verify scratchpad is GONE in the archive
	foundScratchpad := false
	foundArchive := false
	var archiveDir string
	for k := range fs.Files {
		if strings.HasPrefix(k, ".context/archive/") && strings.HasSuffix(k, "_"+trackName+"/spec.md") {
			foundArchive = true
			archiveDir = strings.TrimSuffix(k, "/spec.md")
		}
		if strings.Contains(k, "scratchpad.md") {
			foundScratchpad = true
		}
	}

	if !foundArchive {
		t.Fatal("expected track to be archived")
	}
	if foundScratchpad {
		t.Errorf("expected scratchpad.md to be deleted before archiving")
	}

	// Verify timestamp format (YYYYMMDDHHMMSS)
	// Example: .context/archive/20260110150405_cleanup-track
	base := filepath.Base(archiveDir)
	timestamp := strings.Split(base, "_")[0]
	if len(timestamp) != 14 {
		t.Errorf("expected 14-digit timestamp (YYYYMMDDHHMMSS), got %s (%d digits)", timestamp, len(timestamp))
	}
}

func TestArchiveCmd_PendingTasks(t *testing.T) {
	fs := platform.NewMockFileSystem()
	trackName := "wip-track"
	trackDir := ".context/tracks/" + trackName
	_ = fs.WriteFile(trackDir+"/spec.md", []byte(""), 0644)
	_ = fs.WriteFile(trackDir+"/plan.md", []byte("- [ ] Pending Task"), 0644)

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
		t.Errorf("expected error message about pending tasks, got: %v", err)
	}
}

func TestArchiveCmd_TimeTracking(t *testing.T) {
	fs := platform.NewMockFileSystem()
	trackName := "time-track"
	trackDir := ".context/tracks/" + trackName

	// Setup track files
	_ = fs.WriteFile(trackDir+"/spec.md", []byte("Spec Content"), 0644)
	_ = fs.WriteFile(trackDir+"/plan.md", []byte("- [x] Done"), 0644)

	// Create metadata.json with a timestamp from 1 hour ago
	startedAt := time.Now().Add(-1 * time.Hour).Format(time.RFC3339)
	metadataContent := fmt.Sprintf(`{"started_at": "%s"}`, startedAt)
	_ = fs.WriteFile(trackDir+"/metadata.json", []byte(metadataContent), 0644)

	command := cmd.NewArchiveCmd(fs)
	buf := new(bytes.Buffer)
	command.SetOut(buf)

	command.SetArgs([]string{trackName})
	err := command.Execute()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "Duration:") {
		t.Errorf("expected output to contain duration, got: %s", output)
	}
}

func TestArchiveCmd_InboxIntegration(t *testing.T) {
	fs := platform.NewMockFileSystem()
	trackName := "inbox-track"
	trackDir := ".context/tracks/" + trackName

	specContent := "Spec Content for Inbox"
	_ = fs.WriteFile(trackDir+"/spec.md", []byte(specContent), 0644)
	_ = fs.WriteFile(trackDir+"/plan.md", []byte("- [x] Done"), 0644)

	// Pre-existing inbox
	_ = fs.MkdirAll(".context", 0755)
	_ = fs.WriteFile(".context/inbox.md", []byte("Existing Data"), 0644)

	command := cmd.NewArchiveCmd(fs)
	buf := new(bytes.Buffer)
	command.SetOut(buf)

	command.SetArgs([]string{trackName})
	err := command.Execute()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Verify inbox content
	inboxContent, err := fs.ReadFile(".context/inbox.md")
	if err != nil {
		t.Fatalf("expected inbox to exist")
	}
	contentStr := string(inboxContent)
	if !strings.Contains(contentStr, "Existing Data") {
		t.Errorf("expected original inbox content to be preserved")
	}
	if !strings.Contains(contentStr, specContent) {
		t.Errorf("expected spec content to be appended")
	}
	// Check for timestamp/header
	if !strings.Contains(contentStr, "Archived at:") {
		t.Errorf("expected header with timestamp")
	}
	if !strings.Contains(contentStr, "\n---\n") {
		t.Errorf("expected separator")
	}
}
