package cmd

import (
	"bytes"
	"strings"
	"testing"

	"cdd/internal/platform"
)

func TestBuildViewMarkdown(t *testing.T) {
	fs := platform.NewMockFileSystem()

	// Setup mock filesystem
	_ = fs.MkdirAll(".context/tracks/feature-1", 0755)
	_ = fs.WriteFile(".context/tracks/feature-1/plan.md", []byte("# Plan\n- [ ] Task 1\n- [x] Task 2"), 0644)
	_ = fs.WriteFile(".context/tracks/feature-1/spec.md", []byte("# Spec\nDetails here"), 0644)

	_ = fs.MkdirAll(".context/archive/20260101120000_old-feat", 0755)
	_ = fs.WriteFile(".context/archive/20260101120000_old-feat/plan.md", []byte("# Plan\n- [x] Done"), 0644)

	t.Run("Default Dashboard (Markdown)", func(t *testing.T) {

		viewArchived = false
		viewRaw = false // Forces TTY simulation in tests if we want markdown
		md, err := buildViewMarkdown(fs, []string{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		// Note: In tests, isatty.IsTerminal usually returns false, so it defaults to raw.
		// We expect the raw list "feature-1\n" by default in test environment.
		if !strings.Contains(md, "feature-1") {
			t.Errorf("expected 'feature-1' in output, got %q", md)
		}
	})

	t.Run("Raw Mode Active Tracks", func(t *testing.T) {
		viewRaw = true
		viewArchived = false
		md, err := buildViewMarkdown(fs, []string{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		expected := "feature-1\n"
		if md != expected {
			t.Errorf("expected %q, got %q", expected, md)
		}
	})

	t.Run("Raw Mode Archived Tracks", func(t *testing.T) {
		viewRaw = true
		viewArchived = true
		md, err := buildViewMarkdown(fs, []string{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		expected := "old-feat\n"
		if md != expected {
			t.Errorf("expected %q, got %q", expected, md)
		}
	})

	t.Run("Track Plan (Default)", func(t *testing.T) {
		viewRaw = false

		viewArchived = false
		viewSpec = false
		md, err := buildViewMarkdown(fs, []string{"feature-1"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !strings.Contains(md, "Track: feature-1") {
			t.Errorf("expected 'Track: feature-1' in output, got %q", md)
		}
		if !strings.Contains(md, "Task 1") {
			t.Errorf("expected 'Task 1' (pending) in output, got %q", md)
		}
	})

	t.Run("Track Spec", func(t *testing.T) {
		viewRaw = false
		viewSpec = true
		md, err := buildViewMarkdown(fs, []string{"feature-1"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !strings.Contains(md, "Specification") {
			t.Errorf("expected 'Specification' in output, got %q", md)
		}
		if !strings.Contains(md, "Details here") {
			t.Errorf("expected 'Details here' in output, got %q", md)
		}
	})
}

func TestViewCmd_Help(t *testing.T) {
	fs := platform.NewMockFileSystem()
	command := NewViewCmd(fs)
	buf := new(bytes.Buffer)
	command.SetOut(buf)
	command.SetErr(buf)

	command.SetArgs([]string{"--help"})
	err := command.Execute()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	output := buf.String()
	expected := "Usage:\n  view [track-name] [flags]"
	if !strings.Contains(output, expected) {
		t.Errorf("expected help output to contain usage, got %s", output)
	}

	if !strings.Contains(output, "EXAMPLES:") {
		t.Errorf("expected help output to contain EXAMPLES section")
	}
}

func TestViewCmd_Completion(t *testing.T) {
	fs := platform.NewMockFileSystem()

	// Setup mock filesystem with active tracks
	_ = fs.WriteFile(".context/tracks/feature-alpha/plan.md", []byte(""), 0644)
	_ = fs.WriteFile(".context/tracks/feature-beta/plan.md", []byte(""), 0644)

	t.Run("suggests all active tracks for completion", func(t *testing.T) {
		command := NewViewCmd(fs)
		if command.ValidArgsFunction != nil {
			// Call the function to get suggestions
			args, _ := command.ValidArgsFunction(command, []string{}, "")
			if len(args) != 2 {
				t.Errorf("expected 2 suggestions, got %d: %v", len(args), args)
			}
			// Check that both track names are in suggestions
			expected := map[string]bool{"feature-alpha": false, "feature-beta": false}
			for _, arg := range args {
				if _, ok := expected[arg]; ok {
					expected[arg] = true
				}
			}
			for track, found := range expected {
				if !found {
					t.Errorf("expected track '%s' in suggestions", track)
				}
			}
		} else {
			t.Errorf("expected ValidArgsFunction to be set on view command")
		}
	})

	t.Run("filters completion suggestions by prefix", func(t *testing.T) {
		command := NewViewCmd(fs)
		if command.ValidArgsFunction != nil {
			args, _ := command.ValidArgsFunction(command, []string{}, "feature-a")
			if len(args) != 1 {
				t.Errorf("expected 1 suggestion for 'feature-a', got %d: %v", len(args), args)
			}
			if len(args) > 0 && args[0] != "feature-alpha" {
				t.Errorf("expected 'feature-alpha', got %s", args[0])
			}
		}
	})
}

func TestViewCmd_SingleTaskCompletion(t *testing.T) {
	fs := platform.NewMockFileSystem()

	// Setup mock filesystem with single active track
	_ = fs.WriteFile(".context/tracks/my-single-task/plan.md", []byte(""), 0644)

	t.Run("suggests single task for completion", func(t *testing.T) {
		command := NewViewCmd(fs)
		if command.ValidArgsFunction != nil {
			args, _ := command.ValidArgsFunction(command, []string{}, "")
			if len(args) != 1 {
				t.Errorf("expected 1 suggestion, got %d: %v", len(args), args)
			}
			if len(args) > 0 && args[0] != "my-single-task" {
				t.Errorf("expected 'my-single-task', got %s", args[0])
			}
		}
	})
}

func TestViewCmd_NoTaskCompletion(t *testing.T) {
	fs := platform.NewMockFileSystem()

	// No active tracks

	t.Run("returns empty suggestions when no tasks exist", func(t *testing.T) {
		command := NewViewCmd(fs)
		if command.ValidArgsFunction != nil {
			args, _ := command.ValidArgsFunction(command, []string{}, "")
			if len(args) != 0 {
				t.Errorf("expected 0 suggestions, got %d: %v", len(args), args)
			}
		}
	})
}
