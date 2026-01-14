package cmd

import (
	"testing"
)

func TestExtractVersion(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    string
	}{
		{
			name:    "version with quotes",
			content: "---\nname: cdd\nversion: \"1.0.0\"\n---\nContent",
			want:    "1.0.0",
		},
		{
			name:    "version without quotes",
			content: "---\nname: cdd\nversion: 1.5.2\n---\nContent",
			want:    "1.5.2",
		},
		{
			name:    "version with metadata prefix",
			content: "---\nname: cdd\nmetadata:\n  version: \"2.3.1\"\n---\nContent",
			want:    "2.3.1",
		},
		{
			name:    "no version found",
			content: "---\nname: cdd\n---\nNo version here",
			want:    "0.0.0",
		},
		{
			name:    "empty content",
			content: "",
			want:    "0.0.0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractVersion(tt.content)
			if got != tt.want {
				t.Errorf("extractVersion() = %q, want %q", got, tt.want)
			}
		})
	}
}
