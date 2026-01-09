package cmd_test

import (
	"bytes"
	"testing"

	"cdd/internal/cmd"
	"cdd/internal/platform"
)

func TestListCmd_ListsActiveTracks(t *testing.T) {
	fs := platform.NewMockFileSystem()
	// Setup tracks
	// MockFS relies on file presence for directories.
	fs.WriteFile(".context/tracks/track1/spec.md", []byte(""), 0644)
	fs.WriteFile(".context/tracks/track2/spec.md", []byte(""), 0644)
	// Files that are not directories should be ignored?
	fs.WriteFile(".context/tracks/README.md", []byte(""), 0644)

	command := cmd.NewListCmd(fs)
	buf := new(bytes.Buffer)
	command.SetOut(buf)
	command.SetErr(buf)

	err := command.Execute()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	output := buf.String()
	if !contains(output, "ACTIVE TRACKS:") {
		t.Errorf("expected header 'ACTIVE TRACKS:', got '%s'", output)
	}
	if !contains(output, "- track1") {
		t.Errorf("expected track1, got '%s'", output)
	}
	if !contains(output, "- track2") {
		t.Errorf("expected track2, got '%s'", output)
	}
	// README.md is a file, should not be listed as a track (which are folders)
	// The original code checks `entry.IsDir()`.
	// My Mock `ReadDir` sets `isDir` based on traversal.
	// If I have `.context/tracks/README.md`, `ReadDir` will return `README.md` as false isDir.
	// So it should work.
	if contains(output, "- README.md") {
		t.Errorf("expected README.md to be ignored, got '%s'", output)
	}
}

func TestListCmd_ListsArchivedTracks(t *testing.T) {
	fs := platform.NewMockFileSystem()
	fs.WriteFile(".context/archive/old-track/spec.md", []byte(""), 0644)

	command := cmd.NewListCmd(fs)
	buf := new(bytes.Buffer)
	command.SetOut(buf)
	command.SetErr(buf)

	command.SetArgs([]string{"--archived"})
	err := command.Execute()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	output := buf.String()
	if !contains(output, "ARCHIVED TRACKS:") {
		t.Errorf("expected header, got '%s'", output)
	}
	if !contains(output, "- old-track") {
		t.Errorf("expected old-track, got '%s'", output)
	}
}
