package cmd

import (
	"bytes"
	"strings"
	"testing"

	"cdd/internal/platform"
)

func TestAgentsCmd_Help(t *testing.T) {
	buf := new(bytes.Buffer)
	mockFS := platform.NewMockFileSystem()
	agentsCmd := NewAgentsCmd(mockFS)
	agentsCmd.SetOut(buf)
	agentsCmd.SetErr(buf)

	agentsCmd.SetArgs([]string{"--help"})
	err := agentsCmd.Execute()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	output := buf.String()
	expected := "Usage:\n  agents [flags]"
	if !strings.Contains(output, expected) {
		t.Errorf("expected help output to contain usage, got %s", output)
	}

	if !strings.Contains(output, "EXAMPLES:") {
		t.Errorf("expected help output to contain EXAMPLES section")
	}
}
