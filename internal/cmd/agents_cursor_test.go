package cmd

import (
	"path/filepath"
	"strings"
	"testing"

	"cdd/internal/platform"
)

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

	// Verify .cursor/rules directory exists
	rulesDir := filepath.Join(".cursor", "rules")
	_, err = fs.Stat(rulesDir)
	if err != nil {
		t.Errorf("expected .cursor/rules directory to exist: %v", err)
	}

	// Verify individual rule files exist
	cddRuleFile := filepath.Join(rulesDir, "cdd.mdc")
	_, err = fs.Stat(cddRuleFile)
	if err != nil {
		t.Errorf("expected cdd.mdc rule file to exist: %v", err)
	}

	analystRuleFile := filepath.Join(rulesDir, "cdd-analyst.mdc")
	_, err = fs.Stat(analystRuleFile)
	if err != nil {
		t.Errorf("expected cdd-analyst.mdc rule file to exist: %v", err)
	}

	// Verify content
	content, err := fs.ReadFile(cddRuleFile)
	if err != nil {
		t.Fatalf("failed to read cdd.mdc: %v", err)
	}

	if !strings.Contains(string(content), "# CDD Orchestrator") {
		t.Error("expected orchestrator content in cdd.mdc")
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
	backupPath := filepath.Join(".cursor", "rules", "cdd.mdc.bak")
	_, err = fs.Stat(backupPath)
	if err == nil {
		t.Error("expected NO backup file when version is same")
	}
}

func TestInstallCursorRulesUpdate(t *testing.T) {
	fs := platform.NewMockFileSystem()

	rulesDir := filepath.Join(".cursor", "rules")
	cddRuleFile := filepath.Join(rulesDir, "cdd.mdc")

	// Create rules directory and setup old version
	if err := fs.MkdirAll(rulesDir, 0755); err != nil {
		t.Fatalf("failed to create rules directory: %v", err)
	}
	if err := fs.WriteFile(cddRuleFile, []byte("---\nname: cdd\nversion: 0.5.0\n---\nOld content"), 0644); err != nil {
		t.Fatalf("failed to write old cdd.mdc: %v", err)
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
	backupPath := filepath.Join(rulesDir, "cdd.mdc.bak")
	_, err = fs.Stat(backupPath)
	if err != nil {
		t.Errorf("expected backup file cdd.mdc.bak: %v", err)
	}

	// Verify new content
	content, err := fs.ReadFile(cddRuleFile)
	if err != nil {
		t.Fatalf("failed to read cdd.mdc: %v", err)
	}

	if !strings.Contains(string(content), "# CDD Orchestrator") {
		t.Error("expected new content in cdd.mdc")
	}
}
