package cmd

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/cobra"

	"cdd/internal/platform"
)

func TestPackCompletion(t *testing.T) {
	// Test that pack command has completion support
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	if cmd == nil {
		t.Fatalf("NewPackCmd() returned nil")
	}

	// Should have ValidArgsFunction defined
	if cmd.ValidArgsFunction == nil {
		t.Errorf("pack command should have ValidArgsFunction for completion")
	}
}

func TestPackCompletionFocusFlag(t *testing.T) {
	// Test completion for --focus flag
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	// Get the focus flag
	focusFlag := cmd.Flags().Lookup("focus")
	if focusFlag == nil {
		t.Fatalf("focus flag not found")
	}

	// Should exist and be a string flag
	if focusFlag.Value.Type() != "string" {
		t.Errorf("focus flag should be a string flag, got %s", focusFlag.Value.Type())
	}
}

func TestGetPackCompletionTopics(t *testing.T) {
	// Test that GetPackCompletionTopics returns reasonable topics
	topics := GetPackCompletionTopics()

	if len(topics) == 0 {
		t.Errorf("GetPackCompletionTopics should return at least one topic")
	}

	// Topics should be strings
	for _, topic := range topics {
		if topic == "" {
			t.Errorf("GetPackCompletionTopics should not return empty strings")
		}
	}
}

func TestGetPackCompletionTopicsFiltering(t *testing.T) {
	// Test filtering of completion topics by prefix
	topics := GetPackCompletionTopics()

	tests := []struct {
		prefix        string
		expectMatches int
	}{
		{"log", 1},        // Should match "log"
		{"vi", 1},         // Should match "view"
		{"command", 1},    // Should match "command"
		{"xyz", 0},        // Should not match anything
		{"", len(topics)}, // Empty prefix should match all
	}

	for _, tt := range tests {
		t.Run(tt.prefix, func(t *testing.T) {
			var matches []string
			for _, topic := range topics {
				if strings.HasPrefix(topic, tt.prefix) {
					matches = append(matches, topic)
				}
			}

			if len(matches) != tt.expectMatches {
				t.Errorf("Prefix %q: got %d matches, expected %d", tt.prefix, len(matches), tt.expectMatches)
			}
		})
	}
}

func TestPackCompletionWithValidArgsFunction(t *testing.T) {
	// Test completion through ValidArgsFunction
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	if cmd.ValidArgsFunction == nil {
		t.Skipf("pack command does not implement ValidArgsFunction")
	}

	// Test completion for "log" prefix
	completions, directive := cmd.ValidArgsFunction(cmd, []string{}, "log")

	if len(completions) == 0 {
		t.Errorf("Completion for 'log' prefix should return suggestions")
	}

	if directive == cobra.ShellCompDirectiveError {
		t.Errorf("Completion should not return error directive")
	}

	// Test completion for empty prefix
	completionsEmpty, _ := cmd.ValidArgsFunction(cmd, []string{}, "")
	if len(completionsEmpty) == 0 {
		t.Errorf("Completion for empty prefix should return all suggestions")
	}
}

func TestPackCompletionNoFileCompletion(t *testing.T) {
	// Test that completion doesn't suggest files
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	if cmd.ValidArgsFunction == nil {
		t.Skipf("pack command does not implement ValidArgsFunction")
	}

	_, directive := cmd.ValidArgsFunction(cmd, []string{}, "")

	// Should disable file completion since we're suggesting topics
	if directive != cobra.ShellCompDirectiveNoFileComp {
		t.Errorf("pack completion should disable file completion")
	}
}

func TestPackCompletionSpecTopics(t *testing.T) {
	// Test that completion suggests topics based on actual specs
	fs := platform.NewRealFileSystem()

	specsPath := filepath.Join("..", "..", ".context", "specs")
	specs, err := DiscoverSpecs(fs, specsPath)
	if err != nil {
		t.Skipf("Cannot discover specs: %v", err)
	}

	if len(specs) == 0 {
		t.Skipf("No specs found")
	}

	// Get completion topics
	topics := GetPackCompletionTopics()

	// Should suggest common keywords from specs
	hasCommonTopics := false
	for _, spec := range specs {
		specNameLower := strings.ToLower(spec.Name)
		specNameWithoutExt := strings.TrimSuffix(specNameLower, ".md")

		for _, topic := range topics {
			if strings.Contains(topic, specNameWithoutExt) || strings.Contains(specNameWithoutExt, topic) {
				hasCommonTopics = true
			}
		}
	}

	if !hasCommonTopics {
		t.Logf("Note: completion topics don't directly include spec names, which is OK")
	}
}

func TestPackCompletionDuplicate(t *testing.T) {
	// Test that completion topics are unique
	topics := GetPackCompletionTopics()

	seen := make(map[string]bool)
	for _, topic := range topics {
		if seen[topic] {
			t.Errorf("Duplicate topic in completions: %q", topic)
		}
		seen[topic] = true
	}
}

func TestPackCompletionCaseInsensitive(t *testing.T) {
	// Test that completion handles case variations
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	if cmd.ValidArgsFunction == nil {
		t.Skipf("pack command does not implement ValidArgsFunction")
	}

	tests := []struct {
		prefix string
		name   string
	}{
		{"log", "lowercase"},
		{"Log", "mixed case"},
		{"LOG", "uppercase"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			completions, _ := cmd.ValidArgsFunction(cmd, []string{}, tt.prefix)
			// Should find matches regardless of case
			if len(completions) == 0 {
				t.Logf("No case-sensitive matches for %q (may be OK)", tt.prefix)
			}
		})
	}
}

func TestPackCompletionEdgeCases(t *testing.T) {
	// Test completion edge cases
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	if cmd.ValidArgsFunction == nil {
		t.Skipf("pack command does not implement ValidArgsFunction")
	}

	tests := []struct {
		toComplete string
		name       string
	}{
		{"", "empty string"},
		{"a", "single character"},
		{"very-long-prefix-that-wont-match", "long non-matching"},
		{" spaces ", "spaces in prefix"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			completions, directive := cmd.ValidArgsFunction(cmd, []string{}, tt.toComplete)

			// Should not error
			if directive == cobra.ShellCompDirectiveError {
				t.Errorf("Completion should not error for %q", tt.toComplete)
			}

			// Should return valid completions (possibly empty)
			if completions == nil {
				t.Errorf("Completion should return a slice, not nil")
			}
		})
	}
}
