package cmd

import (
	"bytes"
	"strings"
	"testing"

	"cdd/internal/platform"
)

func TestAgentsCmd_NoDefaultTargetFallback(t *testing.T) {
	mockFS := platform.NewMockFileSystem()

	agentsCmd := NewAgentsCmd(mockFS)
	buf := new(bytes.Buffer)
	agentsCmd.SetOut(buf)
	agentsCmd.SetErr(buf)

	// Execute without specifying target or --all flag
	agentsCmd.SetArgs([]string{"--install"})
	_ = agentsCmd.Execute()

	output := buf.String()

	// Should not successfully install without target or --all flag
	if strings.Contains(output, "installed at") {
		t.Errorf("expected no installation without target or --all flag, but got output: %s", output)
	}

	// Should indicate an error occurred
	if !strings.Contains(output, "Error") && !strings.Contains(output, "error") && !strings.Contains(output, "required") {
		t.Errorf("expected error message when target is missing, but got: %s", output)
	}
}
