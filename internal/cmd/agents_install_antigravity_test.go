package cmd

import (
	"testing"

	"cdd/internal/platform"
	"cdd/prompts"
)

func TestInstallAntigravityWorkflow(t *testing.T) {
	fs := platform.NewMockFileSystem()

	s := skill{
		id:          "cdd-system",
		name:        "cdd-system",
		description: "The Orchestrator",
		content:     prompts.System,
	}

	err := installAntigravityWorkflow(nil, fs, s)
	if err != nil {
		t.Fatalf("installAntigravityWorkflow() failed: %v", err)
	}

	// Verify file was created in .agent/workflows/cdd-system.md
	workflowFile := ".agent/workflows/cdd-system.md"
	_, err = fs.Stat(workflowFile)
	if err != nil {
		t.Fatalf("failed to stat workflow file: %v", err)
	}

	// Verify content was written
	content, err := fs.ReadFile(workflowFile)
	if err != nil {
		t.Fatalf("failed to read workflow file: %v", err)
	}
	if len(content) == 0 {
		t.Fatal("workflow file is empty")
	}
}

func TestInstallAntigravityWorkflow_CreatesDirectory(t *testing.T) {
	fs := platform.NewMockFileSystem()

	s := skill{
		id:          "cdd-analyst",
		name:        "cdd-analyst",
		description: "Analyst",
		content:     prompts.Analyst,
	}

	err := installAntigravityWorkflow(nil, fs, s)
	if err != nil {
		t.Fatalf("installAntigravityWorkflow() failed: %v", err)
	}

	// Verify directory was created
	workflowDir := ".agent/workflows"
	info, err := fs.Stat(workflowDir)
	if err != nil {
		t.Fatalf("failed to stat workflow directory: %v", err)
	}
	if !info.IsDir() {
		t.Fatalf("expected %s to be a directory", workflowDir)
	}
}

func TestInstallAntigravityWorkflow_ValidatesContent(t *testing.T) {
	fs := platform.NewMockFileSystem()

	// Create a workflow with invalid content (no frontmatter)
	s := skill{
		id:          "bad-workflow",
		name:        "bad-workflow",
		description: "Bad",
		content:     "This is not valid YAML frontmatter",
	}

	err := installAntigravityWorkflow(nil, fs, s)
	if err == nil {
		t.Fatal("expected installAntigravityWorkflow() to fail with invalid content")
	}

	if err.Error() != "invalid workflow content: workflow content must start with YAML frontmatter (---)" {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestInstallAntigravityWorkflow_ExistingFile(t *testing.T) {
	fs := platform.NewMockFileSystem()

	s := skill{
		id:          "cdd-test",
		name:        "cdd-test",
		description: "Test",
		content:     prompts.System,
	}

	// Install once
	err := installAntigravityWorkflow(nil, fs, s)
	if err != nil {
		t.Fatalf("first install failed: %v", err)
	}

	// Install again (same version should skip)
	err = installAntigravityWorkflow(nil, fs, s)
	if err != nil {
		t.Fatalf("second install failed: %v", err)
	}

	// Verify file still exists
	workflowFile := ".agent/workflows/cdd-test.md"
	_, err = fs.Stat(workflowFile)
	if err != nil {
		t.Fatalf("workflow file missing after second install: %v", err)
	}
}
