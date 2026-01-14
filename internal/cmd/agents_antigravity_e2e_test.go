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

	// Verify all five skills are installed in .agent/skills/
	expectedSkills := []string{
		".agent/skills/cdd/SKILL.md",
		".agent/skills/cdd-analyst/SKILL.md",
		".agent/skills/cdd-architect/SKILL.md",
		".agent/skills/cdd-executor/SKILL.md",
		".agent/skills/cdd-integrator/SKILL.md",
	}

	for _, skillFile := range expectedSkills {
		_, err := fs.Stat(skillFile)
		if err != nil {
			t.Errorf("failed to stat skill file %s: %v", skillFile, err)
		}

		// Verify content exists
		content, err := fs.ReadFile(skillFile)
		if err != nil {
			t.Errorf("failed to read skill file %s: %v", skillFile, err)
		}
		if len(content) == 0 {
			t.Errorf("skill file %s is empty", skillFile)
		}
	}

	// Second test: run again with same version (should be idempotent)
	cmd2 := NewAgentsCmd(fs)
	cmd2.SetArgs([]string{"--install", "--target", "antigravity"})
	err = cmd2.Execute()
	if err != nil {
		t.Fatalf("second Execute() failed: %v", err)
	}

	// Verify skills still exist and are valid
	for _, skillFile := range expectedSkills {
		_, err := fs.Stat(skillFile)
		if err != nil {
			t.Errorf("skill file missing after second install: %v", err)
		}
	}
}

func TestAgentsInstallAntigravityAllSkillsValid(t *testing.T) {
	fs := platform.NewMockFileSystem()
	cmd := NewAgentsCmd(fs)

	cmd.SetArgs([]string{"--install", "--target", "antigravity"})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Execute() failed: %v", err)
	}

	// Verify each installed skill passes validation
	skillDirs := []string{
		".agent/skills/cdd",
		".agent/skills/cdd-analyst",
		".agent/skills/cdd-architect",
		".agent/skills/cdd-executor",
		".agent/skills/cdd-integrator",
	}

	for _, skillDir := range skillDirs {
		skillFile := skillDir + "/SKILL.md"
		content, err := fs.ReadFile(skillFile)
		if err != nil {
			t.Errorf("failed to read %s: %v", skillFile, err)
			continue
		}

		err = validateAntigravitySkill(string(content))
		if err != nil {
			t.Errorf("skill %s failed validation: %v", skillDir, err)
		}
	}
}
