package cmd

import (
	"testing"

	"cdd/internal/platform"
)

func TestAgentsInstallClaude(t *testing.T) {
	fs := platform.NewMockFileSystem()
	cmd := NewAgentsCmd(fs)

	cmd.SetArgs([]string{"--install", "--target", "claude"})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Execute() failed: %v", err)
	}

	// Verify directory creation in .claude
	skillDir := ".claude/skills/cdd"
	_, err = fs.Stat(skillDir)
	if err != nil {
		t.Fatalf("failed to stat skill directory: %v", err)
	}
}

func TestAgentsInstallPluralAgents(t *testing.T) {
	fs := platform.NewMockFileSystem()
	cmd := NewAgentsCmd(fs)

	cmd.SetArgs([]string{"--install", "--target", "agents"})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Execute() failed: %v", err)
	}

	// Verify directory creation in .agents
	skillDir := ".agents/skills/cdd"
	_, err = fs.Stat(skillDir)
	if err != nil {
		t.Fatalf("failed to stat skill directory: %v", err)
	}
}

func TestAgentsInstallAntigravity(t *testing.T) {
	fs := platform.NewMockFileSystem()
	cmd := NewAgentsCmd(fs)

	cmd.SetArgs([]string{"--install", "--target", "antigravity"})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Execute() failed: %v", err)
	}

	// Verify directory creation in .agent/workflows
	workflowDir := ".agent/workflows"
	_, err = fs.Stat(workflowDir)
	if err != nil {
		t.Fatalf("failed to stat workflow directory: %v", err)
	}

	// Verify workflow file exists
	workflowFile := ".agent/workflows/cdd.md"
	_, err = fs.Stat(workflowFile)
	if err != nil {
		t.Fatalf("failed to stat workflow file: %v", err)
	}
}
