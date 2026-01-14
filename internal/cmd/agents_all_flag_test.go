package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestAgentsCmd_AllFlagIsRecognized(t *testing.T) {
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"agents", "--install", "--all"})
	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("expected no error with --all flag, got %v", err)
	}

	output := buf.String()
	// The flag should be recognized and not produce an unknown flag error
	if strings.Contains(output, "unknown flag") || strings.Contains(output, "flag provided but not defined") {
		t.Errorf("--all flag was not recognized: %s", output)
	}
}
