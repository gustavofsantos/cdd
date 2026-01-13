package cmd

import (
	"bytes"
	"testing"
)

func TestInitFlags(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "executor prompt flag",
			args: []string{"init", "--executor-prompt"},
			want: "Executor Prompt", // Just a partial check if needed, or check if it's non-empty
		},
		{
			name: "planner prompt flag",
			args: []string{"init", "--planner-prompt"},
			want: "Planner Prompt",
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
		})
	}
}
