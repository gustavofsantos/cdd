package cmd

import (
	"bytes"
	"cdd/internal/platform"
	"path/filepath"
	"testing"
)

func TestDeleteCmd(t *testing.T) {
	tests := []struct {
		name        string
		track       string
		setupFiles  map[string][]byte
		expectedOut string
		expectErr   bool
	}{
		{
			name:  "delete existing track",
			track: "task-1",
			setupFiles: map[string][]byte{
				filepath.Join(".context", "tracks", "task-1", "plan.md"): []byte("content"),
			},
			expectedOut: "Track 'task-1' deleted.\n",
			expectErr:   false,
		},
		{
			name:        "delete non-existent track",
			track:       "non-existent",
			setupFiles:  map[string][]byte{},
			expectedOut: "",
			expectErr:   true,
		},
		{
			name:  "delete archived track (should fail)",
			track: "archived-1",
			setupFiles: map[string][]byte{
				filepath.Join(".context", "archive", "archived-1", "plan.md"): []byte("content"),
			},
			expectedOut: "",
			expectErr:   true,
		},
		{
			name:        "prevent directory traversal",
			track:       "../archive/foo",
			setupFiles:  map[string][]byte{},
			expectedOut: "",
			expectErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := platform.NewMockFileSystem()
			for k, v := range tt.setupFiles {
				fs.WriteFile(k, v, 0644)
			}

			cmd := NewDeleteCmd(fs)
			buf := new(bytes.Buffer)
			cmd.SetOut(buf)

			// Use cmd.SetArgs to pass arguments
			cmd.SetArgs([]string{tt.track})

			err := cmd.Execute()

			if (err != nil) != tt.expectErr {
				t.Errorf("expected error %v, got %v", tt.expectErr, err)
			}

			if !tt.expectErr {
				if buf.String() != tt.expectedOut {
					t.Errorf("expected output %q, got %q", tt.expectedOut, buf.String())
				}

				// Verify deletion
				path := filepath.Join(".context", "tracks", tt.track)
				if _, err := fs.Stat(path); err == nil {
					t.Errorf("track directory still exists")
				}
			}
		})
	}
}
