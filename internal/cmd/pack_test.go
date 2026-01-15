package cmd

import (
	"bytes"
	"strings"
	"testing"

	"cdd/internal/platform"
)

func TestPackCmdBasic(t *testing.T) {
	// Test basic pack command invocation with --focus flag
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	if cmd == nil {
		t.Fatalf("NewPackCmd() returned nil")
	}

	if !strings.HasPrefix(cmd.Use, "pack") {
		t.Errorf("Expected Use to start with 'pack', got %q", cmd.Use)
	}

	if cmd.Short == "" {
		t.Errorf("Short description should not be empty")
	}

	if cmd.Long == "" {
		t.Errorf("Long description should not be empty")
	}
}

func TestPackCmdFlags(t *testing.T) {
	// Test that pack command has required flags
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	// Check for --focus flag
	focusFlag := cmd.Flags().Lookup("focus")
	if focusFlag == nil {
		t.Errorf("--focus flag not found")
	}

	// Check for --raw flag
	rawFlag := cmd.Flags().Lookup("raw")
	if rawFlag == nil {
		t.Errorf("--raw flag not found")
	}
}

func TestPackCmdMissingFocus(t *testing.T) {
	// Test that pack command requires --focus flag
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	var out bytes.Buffer
	cmd.SetOut(&out)

	err := cmd.RunE(cmd, []string{})
	if err == nil {
		t.Errorf("Expected error when --focus flag is missing")
	}
	if !strings.Contains(err.Error(), "focus") && !strings.Contains(err.Error(), "required") {
		t.Errorf("Error message should mention --focus flag, got: %v", err)
	}
}

func TestPackCmdWithFocus(t *testing.T) {
	// Test pack command with valid --focus flag
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	var out bytes.Buffer
	cmd.SetOut(&out)

	// Set focus flag
	if err := cmd.Flags().Set("focus", "log"); err != nil {
		t.Fatalf("Failed to set focus flag: %v", err)
	}

	err := cmd.RunE(cmd, []string{})
	if err != nil {
		t.Errorf("pack command with --focus should not error: %v", err)
	}

	output := out.String()
	if output == "" {
		t.Errorf("pack command should produce output")
	}
}

func TestPackCmdRawFlag(t *testing.T) {
	// Test that --raw flag produces plain text output
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	var out bytes.Buffer
	cmd.SetOut(&out)

	if err := cmd.Flags().Set("focus", "log"); err != nil {
		t.Fatalf("Failed to set focus flag: %v", err)
	}
	if err := cmd.Flags().Set("raw", "true"); err != nil {
		t.Fatalf("Failed to set raw flag: %v", err)
	}

	err := cmd.RunE(cmd, []string{})
	if err != nil {
		t.Errorf("pack command with --raw should not error: %v", err)
	}

	output := out.String()
	// Raw output should not contain ANSI escape codes or rich formatting
	if strings.Contains(output, "\x1b[") {
		t.Errorf("Raw output should not contain ANSI escape codes")
	}
}

func TestPackCmdEmptyResults(t *testing.T) {
	// Test pack command with topic that has no matches
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	var out bytes.Buffer
	cmd.SetOut(&out)

	// Use an extremely unlikely topic
	if err := cmd.Flags().Set("focus", "xyzzynotarealword123456"); err != nil {
		t.Fatalf("Failed to set focus flag: %v", err)
	}

	err := cmd.RunE(cmd, []string{})
	// Should not error, but should handle no results gracefully
	if err != nil {
		t.Errorf("pack command should handle no matches gracefully: %v", err)
	}

	output := out.String()
	// Should indicate no matches were found
	if output == "" {
		t.Errorf("pack command should output a message for no matches")
	}
}
