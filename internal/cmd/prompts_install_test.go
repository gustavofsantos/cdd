package cmd

import (
	"cdd/internal/platform"
	"strings"
	"testing"
)

func TestPromptsInstall(t *testing.T) {
	fs := platform.NewMockFileSystem()
	cmd := NewPromptsCmd(fs)

	cmd.SetArgs([]string{"--install"})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Execute() failed: %v", err)
	}

	// Verify directory creation
	skillDir := ".agent/skills/cdd"
	info, err := fs.Stat(skillDir)
	if err != nil {
		t.Fatalf("failed to stat skill directory: %v", err)
	}
	if !info.IsDir() {
		t.Errorf("expected %s to be a directory", skillDir)
	}

	// Verify file creation
	skillFile := ".agent/skills/cdd/SKILL.md"
	content, err := fs.ReadFile(skillFile)
	if err != nil {
		t.Fatalf("failed to read skill file: %v", err)
	}

	expectedFrontmatter := "name: cdd"
	if !strings.Contains(string(content), expectedFrontmatter) {
		t.Errorf("expected frontmatter to contain '%s', got:\n%s", expectedFrontmatter, string(content))
	}

	expectedDescription := "Protocol for implementing software features using the Context-Driven Development methodology."
	if !strings.Contains(string(content), expectedDescription) {
		t.Errorf("expected description to contain '%s'", expectedDescription)
	}

	expectedRole := "**Role:** You are the CDD Engine."
	if !strings.Contains(string(content), expectedRole) {
		t.Errorf("expected role to be present in content")
	}
}
