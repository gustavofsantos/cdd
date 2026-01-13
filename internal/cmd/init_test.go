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

	// Verify .context directory and files were created
	contextDir := filepath.Join(tmpDir, ".context")
	if _, err := os.Stat(contextDir); os.IsNotExist(err) {
		t.Errorf(".context directory was not created")
	}

	files := []string{
		"product.md",
		"tech-stack.md",
		"workflow.md",
		"patterns.md",
		"inbox.md",
	}
	for _, f := range files {
		if _, err := os.Stat(filepath.Join(contextDir, f)); os.IsNotExist(err) {
			t.Errorf(".context/%s was not created", f)
		}
	}
}
