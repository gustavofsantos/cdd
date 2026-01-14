package cmd

import (
	"bytes"
	"strings"
	"testing"

	"cdd/internal/platform"
)

func TestAgentsCmd_TargetValidationWithExamples(t *testing.T) {
	mockFS := platform.NewMockFileSystem()

	agentsCmd := NewAgentsCmd(mockFS)
	buf := new(bytes.Buffer)
	agentsCmd.SetOut(buf)
	agentsCmd.SetErr(buf)

	// Execute without specifying target or --all flag
	agentsCmd.SetArgs([]string{"--install"})
	_ = agentsCmd.Execute()

	output := buf.String()

	// Verify error message is shown
	if !strings.Contains(output, "Error") {
		t.Errorf("expected 'Error' in output, got: %s", output)
	}

	// Verify helpful message about target/all requirement
	if !strings.Contains(output, "must specify a target") && !strings.Contains(output, "--all") {
		t.Errorf("expected error message about --target or --all, got: %s", output)
	}

	// Verify examples are provided
	if !strings.Contains(output, "cdd agents --install --target") {
		t.Errorf("expected example with --target, got: %s", output)
	}

	if !strings.Contains(output, "cdd agents --install --all") {
		t.Errorf("expected example with --all, got: %s", output)
	}
}
