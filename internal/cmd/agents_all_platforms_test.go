package cmd

import (
	"bytes"
	"strings"
	"testing"

	"cdd/internal/platform"
)

func TestAgentsCmd_AllFlagInstallsAllPlatforms(t *testing.T) {
	// Mock file system
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

	// Verify skills are installed for all three platforms
	platforms := []string{".agent", ".claude", ".agents"}
	for _, platform := range platforms {
		if !strings.Contains(output, platform) {
			t.Errorf("expected --all to install for platform %s, but it wasn't mentioned in output: %s", platform, output)
		}
	}
}
