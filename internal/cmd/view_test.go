package cmd

import (
	"cdd/internal/platform"
	"strings"
	"testing"
)

func TestBuildViewMarkdown(t *testing.T) {
	fs := platform.NewMockFileSystem()

	// Setup mock filesystem
	_ = fs.MkdirAll(".context/tracks/feature-1", 0755)
	_ = fs.WriteFile(".context/tracks/feature-1/plan.md", []byte("# Plan\n- [ ] Task 1\n- [x] Task 2"), 0644)
	_ = fs.WriteFile(".context/tracks/feature-1/spec.md", []byte("# Spec\nDetails here"), 0644)

	_ = fs.MkdirAll(".context/archive/20260101120000_old-feat", 0755)
	_ = fs.WriteFile(".context/archive/20260101120000_old-feat/plan.md", []byte("# Plan\n- [x] Done"), 0644)

	_ = fs.WriteFile(".context/inbox.md", []byte("Pending changes"), 0644)

	t.Run("Default Dashboard (Markdown)", func(t *testing.T) {
		viewInbox = false
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
		viewInbox = false
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
