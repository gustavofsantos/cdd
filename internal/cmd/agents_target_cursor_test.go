package cmd

import (
	"path/filepath"
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

	// Verify .cursor/rules directory exists
	rulesDir := filepath.Join(".cursor", "rules")
	_, err = fs.Stat(rulesDir)
	if err != nil {
		t.Errorf("expected .cursor/rules directory to exist: %v", err)
	}

	// Verify each skill has its own rule file
	expectedSkills := []string{"cdd", "cdd-analyst", "cdd-architect", "cdd-executor", "cdd-integrator"}
	for _, skill := range expectedSkills {
		ruleFile := filepath.Join(rulesDir, skill+".mdc")
		_, err = fs.Stat(ruleFile)
		if err != nil {
			t.Errorf("expected rule file %s to exist: %v", ruleFile, err)
		}

		// Verify content contains the skill
		content, err := fs.ReadFile(ruleFile)
		if err != nil {
			t.Fatalf("failed to read %s: %v", ruleFile, err)
		}

		if !strings.Contains(string(content), skill) {
			t.Errorf("expected skill '%s' in %s content", skill, ruleFile)
		}
	}
}
