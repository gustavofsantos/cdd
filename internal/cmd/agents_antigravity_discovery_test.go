package cmd

import (
	"path/filepath"
	"testing"

	"cdd/internal/platform"
)

func TestAntigravitySkillsAreDiscoverable(t *testing.T) {
	fs := platform.NewMockFileSystem()
	cmd := NewAgentsCmd(fs)

	// Install skills
	cmd.SetArgs([]string{"--install", "--target", "antigravity"})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("installation failed: %v", err)
	}

	// Simulate Antigravity discovery by checking that all skill folders contain valid SKILL.md
	skillIDs := []string{"cdd", "cdd-analyst", "cdd-architect", "cdd-executor", "cdd-integrator"}

	for _, skillID := range skillIDs {
		skillDir := filepath.Join(".agent", "skills", skillID)

		// Check directory exists
		info, err := fs.Stat(skillDir)
		if err != nil {
			t.Errorf("skill directory %s does not exist: %v", skillDir, err)
			continue
		}
		if !info.IsDir() {
			t.Errorf("expected %s to be a directory", skillDir)
			continue
		}

		// Check SKILL.md exists
		skillFile := filepath.Join(skillDir, "SKILL.md")
		content, err := fs.ReadFile(skillFile)
		if err != nil {
			t.Errorf("SKILL.md not found in %s: %v", skillDir, err)
			continue
		}

		// Verify it has valid Antigravity format
		err = validateAntigravitySkill(string(content))
		if err != nil {
			t.Errorf("skill %s failed Antigravity validation: %v", skillID, err)
		}

		// Verify it contains YAML frontmatter with name and description
		contentStr := string(content)
		if !hasYAMLField(contentStr, "name") {
			t.Errorf("skill %s missing 'name' in YAML frontmatter", skillID)
		}
		if !hasYAMLField(contentStr, "description") {
			t.Errorf("skill %s missing 'description' in YAML frontmatter", skillID)
		}
	}
}

func hasYAMLField(content, field string) bool {
	// Simple check for field: in YAML frontmatter
	start := 0
	if content[0:3] == "---" {
		start = 3
	}
	end := len(content)
	if idx := findString(content[start:], "---"); idx != -1 {
		end = start + idx
	}
	frontmatter := content[start:end]
	return findString(frontmatter, field+":") != -1
}

func findString(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
