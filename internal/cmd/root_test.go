package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestRootCmd_Help(t *testing.T) {
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"--help"})

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	output := buf.String()

	// Check for sections we want to add
	sections := []string{
		"CORE PRINCIPLES",
		"WORKFLOW",
		"EXAMPLES",
	}

	for _, section := range sections {
		if !strings.Contains(output, section) {
			t.Errorf("expected help output to contain section '%s'", section)
		}
	}
}
