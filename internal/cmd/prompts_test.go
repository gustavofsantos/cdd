package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestPromptsCommand(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "system prompt",
			args: []string{"prompts", "--system"},
			want: "AGENT SYSTEM PROMPT",
		},

		{
			name: "calibration prompt",
			args: []string{"prompts", "--calibration"},
			want: "CDD Workflow Calibrator",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bytes.NewBufferString("")
			rootCmd.SetOut(b)
			rootCmd.SetArgs(tt.args)

			err := rootCmd.Execute()
			if err != nil {
				t.Fatalf("Execute() failed: %v", err)
			}

			out := b.String()
			if out == "" {
				t.Errorf("expected output, got empty string")
			}
			// Optional: check if output contains the 'want' string
		})
	}
}

func TestPromptsCmd_Help(t *testing.T) {
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"prompts", "--help"})
	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	output := buf.String()
	expected := "Usage:\n  cdd prompts [flags]"
	if !strings.Contains(output, expected) {
		t.Errorf("expected help output to contain usage, got %s", output)
	}

	if !strings.Contains(output, "EXAMPLES:") {
		t.Errorf("expected help output to contain EXAMPLES section")
	}
}
