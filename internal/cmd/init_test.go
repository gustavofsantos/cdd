package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestInitCommand(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "cdd-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	// Change working directory to temp dir
	oldWd, _ := os.Getwd()
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("failed to change dir: %v", err)
	}
	defer func() { _ = os.Chdir(oldWd) }()

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

	// Verify global files
	globalFiles := []string{
		".context/product.md",
		".context/tech-stack.md",
		".context/domain.md",
	}
	for _, f := range globalFiles {
		if _, err := os.Stat(filepath.Join(tmpDir, f)); os.IsNotExist(err) {
			t.Errorf("%s was not created", f)
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
		"current_state.md",
		"decisions.md",
	}
	for _, f := range trackFiles {
		filePath := filepath.Join(setupDir, f)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			t.Errorf(".context/tracks/setup/%s was not created", f)
		}
	}
}

func TestInitCmd_Help(t *testing.T) {
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"init", "--help"})
	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	output := buf.String()
	expected := "Usage:\n  cdd init [flags]"
	if !strings.Contains(output, expected) {
		// Sometimes in tests it doesn't have the parent name
		expected = "Usage:\n  init [flags]"
		if !strings.Contains(output, expected) {
			t.Errorf("expected help output to contain usage, got %s", output)
		}
	}

	if !strings.Contains(output, "EXAMPLES:") {
		t.Errorf("expected help output to contain EXAMPLES section")
	}
}
