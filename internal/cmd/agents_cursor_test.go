package cmd

import (
	"strings"
	"testing"

	"cdd/internal/platform"
)

func TestBuildCursorRulesContent(t *testing.T) {
	skills := []skill{
		{id: "cdd", name: "cdd", description: "Orchestrator", content: "---\nname: cdd\nversion: 1.0.0\n---\n# CDD Orchestrator"},
		{id: "cdd-analyst", name: "cdd-analyst", description: "Analyst", content: "---\nname: cdd-analyst\nversion: 1.0.0\n---\n# CDD Analyst"},
	}

	content := buildCursorRulesContent(skills)

	// Verify all skills are included
	if !strings.Contains(content, "# CDD Orchestrator") {
		t.Error("expected orchestrator content in cursor rules")
	}
	if !strings.Contains(content, "# CDD Analyst") {
		t.Error("expected analyst content in cursor rules")
	}

	// Verify version metadata is present
	if !strings.Contains(content, "version:") {
		t.Error("expected version metadata in cursor rules")
	}

	// Verify structure with separators
	if !strings.Contains(content, "---") {
		t.Error("expected markdown separators in cursor rules")
	}
}

func TestBuildCursorRulesContentEmpty(t *testing.T) {
	skills := []skill{}
	content := buildCursorRulesContent(skills)

	if content == "" {
		t.Error("expected non-empty content even for empty skills")
	}

	// Should still have version metadata
	if !strings.Contains(content, "version:") {
		t.Error("expected version metadata even for empty skills")
	}
}

func TestInstallCursorRules(t *testing.T) {
	fs := platform.NewMockFileSystem()
	skills := []skill{
		{id: "cdd", name: "cdd", description: "Orchestrator", content: "---\nname: cdd\nversion: 1.0.0\n---\n# CDD Orchestrator"},
		{id: "cdd-analyst", name: "cdd-analyst", description: "Analyst", content: "---\nname: cdd-analyst\nversion: 1.0.0\n---\n# CDD Analyst"},
	}

	cmd := NewAgentsCmd(fs)

	err := installCursorRules(cmd, fs, skills)
	if err != nil {
		t.Fatalf("installCursorRules failed: %v", err)
	}

	// Verify .cursorrules file exists
	cursorRulesFile := ".cursorrules"
	_, err = fs.Stat(cursorRulesFile)
	if err != nil {
		t.Errorf("expected .cursorrules file to exist: %v", err)
	}

	// Verify content
	content, err := fs.ReadFile(cursorRulesFile)
	if err != nil {
		t.Fatalf("failed to read .cursorrules: %v", err)
	}

	if !strings.Contains(string(content), "# CDD Orchestrator") {
		t.Error("expected orchestrator content in .cursorrules")
	}
}

func TestInstallCursorRulesIdempotent(t *testing.T) {
	fs := platform.NewMockFileSystem()
	skills := []skill{
		{id: "cdd", name: "cdd", description: "Orchestrator", content: "---\nname: cdd\nversion: 1.0.0\n---\n# CDD Orchestrator"},
	}

	cmd := NewAgentsCmd(fs)

	// First install
	err := installCursorRules(cmd, fs, skills)
	if err != nil {
		t.Fatalf("first installCursorRules failed: %v", err)
	}

	// Second install with same version should not overwrite
	err = installCursorRules(cmd, fs, skills)
	if err != nil {
		t.Fatalf("second installCursorRules failed: %v", err)
	}

	// Should NOT create backup
	backupPath := ".cursorrules.bak"
	_, err = fs.Stat(backupPath)
	if err == nil {
		t.Error("expected NO backup file when version is same")
	}
}

func TestInstallCursorRulesUpdate(t *testing.T) {
	fs := platform.NewMockFileSystem()

	// Setup old version
	cursorRulesFile := ".cursorrules"
	if err := fs.WriteFile(cursorRulesFile, []byte("---\nversion: 0.5.0\n---\nOld content"), 0644); err != nil {
		t.Fatalf("failed to write old .cursorrules: %v", err)
	}

	skills := []skill{
		{id: "cdd", name: "cdd", description: "Orchestrator", content: "---\nname: cdd\nversion: 1.0.0\n---\n# CDD Orchestrator"},
	}

	cmd := NewAgentsCmd(fs)

	// Install with new version
	err := installCursorRules(cmd, fs, skills)
	if err != nil {
		t.Fatalf("installCursorRules failed: %v", err)
	}

	// Verify backup created
	backupPath := ".cursorrules.bak"
	_, err = fs.Stat(backupPath)
	if err != nil {
		t.Errorf("expected backup file .cursorrules.bak: %v", err)
	}

	// Verify new content
	content, err := fs.ReadFile(cursorRulesFile)
	if err != nil {
		t.Fatalf("failed to read .cursorrules: %v", err)
	}

	if !strings.Contains(string(content), "# CDD Orchestrator") {
		t.Error("expected new content in .cursorrules")
	}
}
