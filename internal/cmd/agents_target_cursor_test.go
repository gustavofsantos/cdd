package cmd

import (
	"strings"
	"testing"

	"cdd/internal/platform"
)

func TestAgentsInstallCursorTarget(t *testing.T) {
	fs := platform.NewMockFileSystem()
	cmd := NewAgentsCmd(fs)

	cmd.SetArgs([]string{"--install", "--target", "cursor"})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Execute() failed: %v", err)
	}

	// Verify .cursorrules file exists
	cursorRulesFile := ".cursorrules"
	_, err = fs.Stat(cursorRulesFile)
	if err != nil {
		t.Errorf("expected .cursorrules file to exist: %v", err)
	}

	// Verify content has all skills
	content, err := fs.ReadFile(cursorRulesFile)
	if err != nil {
		t.Fatalf("failed to read .cursorrules: %v", err)
	}

	contentStr := string(content)
	expectedSkills := []string{"cdd", "cdd-analyst", "cdd-architect", "cdd-executor", "cdd-integrator"}
	for _, skill := range expectedSkills {
		if !strings.Contains(contentStr, skill) {
			t.Errorf("expected skill '%s' in .cursorrules content", skill)
		}
	}
}
