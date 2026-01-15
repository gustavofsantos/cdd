package cmd

import (
	"bytes"
	"strings"
	"testing"

	"cdd/internal/platform"
)

func verifyPlatformInstallations(t *testing.T, output string, platforms []string, skills []string) {
	for _, p := range platforms {
		if !strings.Contains(output, p) {
			t.Errorf("expected installation for platform %s, but not found in output", p)
		}
		for _, skill := range skills {
			skillFile := p + "/skills/" + skill + "/SKILL.md"
			if !strings.Contains(output, skillFile) {
				t.Errorf("expected skill file %s in output", skillFile)
			}
		}
	}
}

func TestAgentsCmd_AllFlagIntegration(t *testing.T) {
	mockFS := platform.NewMockFileSystem()

	agentsCmd := NewAgentsCmd(mockFS)
	buf := new(bytes.Buffer)
	agentsCmd.SetOut(buf)
	agentsCmd.SetErr(buf)

	// Execute with --all flag
	agentsCmd.SetArgs([]string{"--install", "--all"})
	_ = agentsCmd.Execute()

	output := buf.String()

	// Verify all directory-based platforms are installed
	expectedPlatforms := []string{".agent", ".claude", ".agents"}
	skills := []string{"cdd", "cdd-analyst", "cdd-architect", "cdd-executor", "cdd-integrator"}

	verifyPlatformInstallations(t, output, expectedPlatforms, skills)

	// Verify cursor and antigravity installations
	if !strings.Contains(output, ".cursor/rules") {
		t.Errorf("expected cursor rules installation in .cursor/rules, but not found in output")
	}

	if !strings.Contains(output, ".agent/skills") {
		t.Errorf("expected antigravity installation, but not found in output")
	}

	// Verify no errors occurred
	if strings.Contains(output, "Error") || strings.Contains(output, "error") {
		t.Errorf("expected no errors, but got: %s", output)
	}
}
