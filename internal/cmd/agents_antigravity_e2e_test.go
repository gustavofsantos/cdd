package cmd

import (
	"testing"

	"cdd/internal/platform"
)

func TestAgentsInstallAntigravityE2E(t *testing.T) {
	fs := platform.NewMockFileSystem()
	cmd := NewAgentsCmd(fs)

	// First test: install with antigravity target
	cmd.SetArgs([]string{"--install", "--target", "antigravity"})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Execute() failed: %v", err)
	}

	// Verify all five workflows are installed in .agent/workflows/
	expectedWorkflows := []string{
		".agent/workflows/cdd.md",
		".agent/workflows/cdd-analyst.md",
		".agent/workflows/cdd-architect.md",
		".agent/workflows/cdd-executor.md",
		".agent/workflows/cdd-integrator.md",
	}

	for _, workflowFile := range expectedWorkflows {
		_, err := fs.Stat(workflowFile)
		if err != nil {
			t.Errorf("failed to stat workflow file %s: %v", workflowFile, err)
		}

		// Verify content exists
		content, err := fs.ReadFile(workflowFile)
		if err != nil {
			t.Errorf("failed to read workflow file %s: %v", workflowFile, err)
		}
		if len(content) == 0 {
			t.Errorf("workflow file %s is empty", workflowFile)
		}
	}

	// Second test: run again with same version (should be idempotent)
	cmd2 := NewAgentsCmd(fs)
	cmd2.SetArgs([]string{"--install", "--target", "antigravity"})
	err = cmd2.Execute()
	if err != nil {
		t.Fatalf("second Execute() failed: %v", err)
	}

	// Verify workflows still exist and are valid
	for _, workflowFile := range expectedWorkflows {
		_, err := fs.Stat(workflowFile)
		if err != nil {
			t.Errorf("workflow file missing after second install: %v", err)
		}
	}
}

func TestAgentsInstallAntigravityAllWorkflowsValid(t *testing.T) {
	fs := platform.NewMockFileSystem()
	cmd := NewAgentsCmd(fs)

	cmd.SetArgs([]string{"--install", "--target", "antigravity"})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Execute() failed: %v", err)
	}

	// Verify each installed workflow passes validation
	workflowFiles := []string{
		".agent/workflows/cdd.md",
		".agent/workflows/cdd-analyst.md",
		".agent/workflows/cdd-architect.md",
		".agent/workflows/cdd-executor.md",
		".agent/workflows/cdd-integrator.md",
	}

	for _, workflowFile := range workflowFiles {
		content, err := fs.ReadFile(workflowFile)
		if err != nil {
			t.Errorf("failed to read %s: %v", workflowFile, err)
			continue
		}

		err = validateAntigravityWorkflow(string(content))
		if err != nil {
			t.Errorf("workflow %s failed validation: %v", workflowFile, err)
		}
	}
}
