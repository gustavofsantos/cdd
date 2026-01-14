package cmd

import (
	"testing"

	"cdd/internal/platform"
)

func TestAgentsMultiInstall(t *testing.T) {
	fs := platform.NewMockFileSystem()
	cmd := NewAgentsCmd(fs)

	cmd.SetArgs([]string{"--install", "--target", "agent"})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Execute() failed: %v", err)
	}

	skills := []string{"cdd", "cdd-analyst", "cdd-architect", "cdd-executor", "cdd-integrator"}
	for _, skill := range skills {
		skillFile := ".agent/skills/" + skill + "/SKILL.md"
		_, err := fs.Stat(skillFile)
		if err != nil {
			t.Errorf("skill file %s not found: %v", skillFile, err)
		}
	}
}
