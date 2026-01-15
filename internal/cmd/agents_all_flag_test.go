package cmd

import (
	"bytes"
	"cdd/internal/platform"
	"strings"
	"testing"
)

func TestAgentsCmd_AllFlagIsRecognized(t *testing.T) {
	mockFS := platform.NewMockFileSystem()
	agentsCmd := NewAgentsCmd(mockFS)
	buf := new(bytes.Buffer)
	agentsCmd.SetOut(buf)
	agentsCmd.SetErr(buf)

	agentsCmd.SetArgs([]string{"--install", "--all"})
	err := agentsCmd.Execute()
	if err != nil {
		t.Fatalf("expected no error with --all flag, got %v", err)
	}

	output := buf.String()
	// The flag should be recognized and not produce an unknown flag error
	if strings.Contains(output, "unknown flag") || strings.Contains(output, "flag provided but not defined") {
		t.Errorf("--all flag was not recognized: %s", output)
	}
}
