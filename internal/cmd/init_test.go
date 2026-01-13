package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestInitCommand(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "cdd-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Change working directory to temp dir
	oldWd, _ := os.Getwd()
	os.Chdir(tmpDir)
	defer os.Chdir(oldWd)

	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"init"})

	err = rootCmd.Execute()
	if err != nil {
		t.Fatalf("Execute() failed: %v", err)
	}

	// Verify .context directory structure was created
	contextDir := filepath.Join(tmpDir, ".context")
	if _, err := os.Stat(contextDir); os.IsNotExist(err) {
		t.Errorf(".context directory was not created")
	}

	// Verify subdirectories
	subdirs := []string{
		".context/tracks",
		".context/archive",
		".context/specs",
	}
	for _, d := range subdirs {
		if _, err := os.Stat(filepath.Join(tmpDir, d)); os.IsNotExist(err) {
			t.Errorf("%s directory was not created", d)
		}
	}

	// Verify setup track was created
	setupDir := filepath.Join(tmpDir, ".context/tracks/setup")
	if _, err := os.Stat(setupDir); os.IsNotExist(err) {
		t.Errorf(".context/tracks/setup directory was not created")
	}

	// Verify setup track files
	trackFiles := []string{
		"spec.md",
		"plan.md",
		"decisions.md",
	}
	for _, f := range trackFiles {
		filePath := filepath.Join(setupDir, f)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			t.Errorf(".context/tracks/setup/%s was not created", f)
		}
	}
}
