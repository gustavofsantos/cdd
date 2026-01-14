package cmd

import (
	"testing"

	"cdd/internal/platform"
	"cdd/prompts"
)

func TestInstallAntigravitySkill(t *testing.T) {
	fs := platform.NewMockFileSystem()

	s := skill{
		id:          "cdd-system",
		name:        "cdd-system",
		description: "The Orchestrator",
		content:     prompts.System,
	}

	err := installAntigravitySkill(nil, fs, s)
	if err != nil {
		t.Fatalf("installAntigravitySkill() failed: %v", err)
	}

	// Verify file was created in .agent/skills/cdd-system/SKILL.md
	skillFile := ".agent/skills/cdd-system/SKILL.md"
	_, err = fs.Stat(skillFile)
	if err != nil {
		t.Fatalf("failed to stat skill file: %v", err)
	}

	// Verify content was written
	content, err := fs.ReadFile(skillFile)
	if err != nil {
		t.Fatalf("failed to read skill file: %v", err)
	}
	if len(content) == 0 {
		t.Fatal("skill file is empty")
	}
}

func TestInstallAntigravitySkill_CreatesDirectory(t *testing.T) {
	fs := platform.NewMockFileSystem()

	s := skill{
		id:          "cdd-analyst",
		name:        "cdd-analyst",
		description: "Analyst",
		content:     prompts.Analyst,
	}

	err := installAntigravitySkill(nil, fs, s)
	if err != nil {
		t.Fatalf("installAntigravitySkill() failed: %v", err)
	}

	// Verify directory was created
	skillDir := ".agent/skills/cdd-analyst"
	info, err := fs.Stat(skillDir)
	if err != nil {
		t.Fatalf("failed to stat skill directory: %v", err)
	}
	if !info.IsDir() {
		t.Fatalf("expected %s to be a directory", skillDir)
	}
}

func TestInstallAntigravitySkill_ValidatesContent(t *testing.T) {
	fs := platform.NewMockFileSystem()

	// Create a skill with invalid content (no frontmatter)
	s := skill{
		id:          "bad-skill",
		name:        "bad-skill",
		description: "Bad",
		content:     "This is not valid YAML frontmatter",
	}

	err := installAntigravitySkill(nil, fs, s)
	if err == nil {
		t.Fatal("expected installAntigravitySkill() to fail with invalid content")
	}

	if err.Error() != "invalid skill content: skill content must start with YAML frontmatter (---)" {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestInstallAntigravitySkill_ExistingFile(t *testing.T) {
	fs := platform.NewMockFileSystem()

	s := skill{
		id:          "cdd-test",
		name:        "cdd-test",
		description: "Test",
		content:     prompts.System,
	}

	// Install once
	err := installAntigravitySkill(nil, fs, s)
	if err != nil {
		t.Fatalf("first install failed: %v", err)
	}

	// Install again (same version should skip)
	err = installAntigravitySkill(nil, fs, s)
	if err != nil {
		t.Fatalf("second install failed: %v", err)
	}

	// Verify file still exists
	skillFile := ".agent/skills/cdd-test/SKILL.md"
	_, err = fs.Stat(skillFile)
	if err != nil {
		t.Fatalf("skill file missing after second install: %v", err)
	}
}
