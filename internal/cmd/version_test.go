package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestVersionCmd_Help(t *testing.T) {
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"version", "--help"})
	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	output := buf.String()
	expected := "Usage:\n  cdd version [flags]"
	if !strings.Contains(output, expected) {
		// Handled as standalone in some test contexts
		expected = "Usage:\n  version [flags]"
		if !strings.Contains(output, expected) {
			t.Errorf("expected help output to contain usage, got %s", output)
		}
	}

	if !strings.Contains(output, "EXAMPLES:") {
		t.Errorf("expected help output to contain EXAMPLES section")
	}
}
