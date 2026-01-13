package cmd

import (
	"bytes"
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
			name: "executor prompt",
			args: []string{"prompts", "--executor"},
			want: "AGENT EXECUTOR PROMPT",
		},
		{
			name: "planner prompt",
			args: []string{"prompts", "--planner"},
			want: "AGENT PLANNER PROMPT",
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
