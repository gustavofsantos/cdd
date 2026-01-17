package cmd

import (
	"strings"
	"testing"

	"cdd/internal/platform"
)

func TestAgentsInstallGemini(t *testing.T) {
	fs := platform.NewMockFileSystem()
	cmd := NewAgentsCmd(fs)

	cmd.SetArgs([]string{"--install", "--target", "gemini"})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Execute() failed: %v", err)
	}

	// Verify directory creation
	skillDir := ".gemini/skills/cdd"
	info, err := fs.Stat(skillDir)
	if err != nil {
		t.Fatalf("failed to stat skill directory: %v", err)
	}
	if !info.IsDir() {
		t.Errorf("expected %s to be a directory", skillDir)
	}

	// Verify file creation
	skillFile := ".gemini/skills/cdd/SKILL.md"
	content, err := fs.ReadFile(skillFile)
	if err != nil {
		t.Fatalf("failed to read skill file: %v", err)
	}

	expectedFrontmatter := "name: cdd"
	if !strings.Contains(string(content), expectedFrontmatter) {
		t.Errorf("expected frontmatter to contain '%s', got:\n%s", expectedFrontmatter, string(content))
	}
}
