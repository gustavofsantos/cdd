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

	t.Run("Default Dashboard", func(t *testing.T) {
		viewInbox = false
		viewArchived = false
		md, err := buildViewMarkdown(fs, []string{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !strings.Contains(md, "Active Tracks") {
			t.Errorf("expected 'Active Tracks' in output, got %q", md)
		}
		if !strings.Contains(md, "feature-1") {
			t.Errorf("expected 'feature-1' in output, got %q", md)
		}
	})

	t.Run("Inbox Flag", func(t *testing.T) {
		viewInbox = true
		viewArchived = false
		md, err := buildViewMarkdown(fs, []string{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !strings.Contains(md, "Context Inbox") {
			t.Errorf("expected 'Context Inbox' in output, got %q", md)
		}
		if !strings.Contains(md, "Pending changes") {
			t.Errorf("expected 'Pending changes' in output, got %q", md)
		}
	})

	t.Run("Archived Flag (Dashboard)", func(t *testing.T) {
		viewInbox = false
		viewArchived = true
		md, err := buildViewMarkdown(fs, []string{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !strings.Contains(md, "Archived Tracks") {
			t.Errorf("expected 'Archived Tracks' in output, got %q", md)
		}
		if !strings.Contains(md, "20260101120000_old-feat") {
			t.Errorf("expected '20260101120000_old-feat' in output, got %q", md)
		}
	})

	t.Run("Track Plan (Default)", func(t *testing.T) {
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
		if strings.Contains(md, "Task 2") {
			t.Errorf("did not expect 'Task 2' (completed) in output, got %q", md)
		}
	})

	t.Run("Track Spec", func(t *testing.T) {
		viewInbox = false
		viewArchived = false
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

	t.Run("Archived Track", func(t *testing.T) {
		viewInbox = false
		viewArchived = true
		viewSpec = false
		md, err := buildViewMarkdown(fs, []string{"old-feat"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !strings.Contains(md, "old-feat") {
			t.Errorf("expected 'old-feat' in output, got %q", md)
		}
		if !strings.Contains(md, "No pending tasks") {
			t.Errorf("expected 'No pending tasks' in output, got %q", md)
		}
	})
}
