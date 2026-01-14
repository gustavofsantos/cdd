package cmd

import (
	"strings"
	"testing"

	"cdd/internal/platform"
)

func TestAgentsInstall(t *testing.T) {
	fs := platform.NewMockFileSystem()
	cmd := NewAgentsCmd(fs)

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

func TestAgentsMigration_Legacy(t *testing.T) {
	fs := platform.NewMockFileSystem()
	cmd := NewAgentsCmd(fs)

	// Setup legacy file (no version)
	skillDir := ".agent/skills/cdd"
	if err := fs.MkdirAll(skillDir, 0755); err != nil {
		t.Fatalf("failed to create skill directory: %v", err)
	}
	legacyPath := skillDir + "/SKILL.md"
	if err := fs.WriteFile(legacyPath, []byte("---\nname: cdd\n---\nOld Content"), 0644); err != nil {
		t.Fatalf("failed to write legacy file: %v", err)
	}

	cmd.SetArgs([]string{"--install"})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Execute() failed: %v", err)
	}

	// Verify Backup Created
	backupPath := legacyPath + ".bak"
	_, err = fs.Stat(backupPath)
	if err != nil {
		t.Errorf("expected backup file %s to exist", backupPath)
	}

	// Verify New Content has Version
	content, _ := fs.ReadFile(legacyPath)
	if !strings.Contains(string(content), "version: ") {
		t.Error("expected new file to have version in frontmatter")
	}
}

func TestAgentsUpToDate(t *testing.T) {
	fs := platform.NewMockFileSystem()
	cmd := NewAgentsCmd(fs)

	// Setup current file (simulate current version)
	skillDir := ".agent/skills/cdd"
	if err := fs.MkdirAll(skillDir, 0755); err != nil {
		t.Fatalf("failed to create skill directory: %v", err)
	}
	currentPath := skillDir + "/SKILL.md"
	if err := fs.WriteFile(currentPath, []byte("---\nname: cdd\nmetadata:\n  version: \"2\"\n---\nNew Content"), 0644); err != nil {
		t.Fatalf("failed to write current file: %v", err)
	}

	cmd.SetArgs([]string{"--install"})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Execute() failed: %v", err)
	}

	// Should NOT create backup (idempotent)
	backupPath := currentPath + ".bak"
	_, err = fs.Stat(backupPath)
	if err == nil {
		t.Error("expected NO backup file when up to date")
	}
}
