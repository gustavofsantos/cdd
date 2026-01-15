package cmd

import (
	"bytes"
	"strings"
	"testing"

	"cdd/internal/platform"
)

func TestPackLimitFlag(t *testing.T) {
	// Test that pack command has limit flag
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	if cmd == nil {
		t.Fatalf("NewPackCmd() returned nil")
	}

	// Check for --limit flag
	limitFlag := cmd.Flags().Lookup("limit")
	if limitFlag == nil {
		t.Errorf("--limit flag not found")
	}

	// Should be an integer flag
	if limitFlag.Value.Type() != "int" {
		t.Errorf("--limit flag should be int type, got %s", limitFlag.Value.Type())
	}
}

func TestPackLimitFlagDefault(t *testing.T) {
	// Test default limit value is -1 (no limit)
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	limitFlag := cmd.Flags().Lookup("limit")
	if limitFlag == nil {
		t.Fatalf("--limit flag not found")
	}

	// Default should be -1 (indicating no limit)
	defaultValue := limitFlag.DefValue
	if defaultValue != "-1" {
		t.Errorf("Default limit should be -1, got %s", defaultValue)
	}
}

func TestPackWithValidLimit(t *testing.T) {
	// Test pack command with valid limit flag
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	var out bytes.Buffer
	cmd.SetOut(&out)

	// Set both focus and limit flags
	cmd.Flags().Set("focus", "log")
	cmd.Flags().Set("limit", "5")

	err := cmd.RunE(cmd, []string{})
	if err != nil {
		t.Errorf("pack command with valid limit should not error: %v", err)
	}

	output := out.String()
	if output == "" {
		t.Errorf("pack command should produce output")
	}
}

func TestPackWithLimitZero(t *testing.T) {
	// Test pack command with limit 0 (show count only)
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	var out bytes.Buffer
	cmd.SetOut(&out)

	cmd.Flags().Set("focus", "log")
	cmd.Flags().Set("limit", "0")

	err := cmd.RunE(cmd, []string{})
	if err != nil {
		t.Errorf("pack command with limit 0 should not error: %v", err)
	}

	output := out.String()
	if output == "" {
		t.Errorf("pack command should produce output even with limit 0")
	}

	// Should mention match count
	if !strings.Contains(output, "Found") {
		t.Errorf("Output should show match count")
	}
}

func TestPackLimitNegativeValue(t *testing.T) {
	// Test pack command with negative limit (should be allowed, treated as no limit)
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	var out bytes.Buffer
	cmd.SetOut(&out)

	cmd.Flags().Set("focus", "log")
	cmd.Flags().Set("limit", "-999")

	err := cmd.RunE(cmd, []string{})
	if err != nil {
		t.Errorf("pack command with negative limit should not error: %v", err)
	}

	output := out.String()
	if output == "" {
		t.Errorf("pack command should produce output")
	}
}

func TestPackLimitHighValue(t *testing.T) {
	// Test pack command with limit higher than available results
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	var out bytes.Buffer
	cmd.SetOut(&out)

	cmd.Flags().Set("focus", "log")
	cmd.Flags().Set("limit", "999999")

	err := cmd.RunE(cmd, []string{})
	if err != nil {
		t.Errorf("pack command with high limit should not error: %v", err)
	}

	output := out.String()
	if output == "" {
		t.Errorf("pack command should produce output")
	}

	// Should still show all matches (up to available)
	if !strings.Contains(output, "Found") {
		t.Errorf("Output should show match count")
	}
}

func TestPackInvalidLimit(t *testing.T) {
	// Test pack command with invalid (non-numeric) limit
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	var out bytes.Buffer
	cmd.SetOut(&out)

	cmd.Flags().Set("focus", "log")
	
	// Try to set invalid limit value
	err := cmd.Flags().Set("limit", "not-a-number")
	if err == nil {
		t.Errorf("Setting invalid limit should error")
	}
}

func TestPackLimitWithRawFlag(t *testing.T) {
	// Test pack command with both limit and raw flags
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	var out bytes.Buffer
	cmd.SetOut(&out)

	cmd.Flags().Set("focus", "command")
	cmd.Flags().Set("limit", "3")
	cmd.Flags().Set("raw", "true")

	err := cmd.RunE(cmd, []string{})
	if err != nil {
		t.Errorf("pack command with limit and raw should not error: %v", err)
	}

	output := out.String()
	if output == "" {
		t.Errorf("pack command should produce output")
	}

	// Raw output should not have ANSI codes
	if strings.Contains(output, "\x1b[") {
		t.Errorf("Raw output should not have ANSI codes")
	}
}
