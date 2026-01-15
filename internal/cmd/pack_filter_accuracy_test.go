package cmd

import (
	"path/filepath"
	"strings"
	"testing"

	"cdd/internal/platform"
)

func TestFilterAccuracyDifferentTopics(t *testing.T) {
	// Test filtering with different topics on same specs
	specs := []Spec{
		{
			Name: "log.md",
			Content: []byte(`# Log Command

The log command records permanent decisions.

## Requirements
The system shall record messages to the decisions.md file.
When the user invokes cdd log, the system shall append a timestamp.`),
		},
		{
			Name: "view.md",
			Content: []byte(`# View Command

The view command displays track information.

## Dashboard
The system shall render a dashboard with active tracks.
When viewing archived tracks, show the archive status.`),
		},
	}

	tests := []struct {
		topic        string
		minScore     float64
		expectFromLog bool
		minMatches   int // At least this many
	}{
		{"log", 0.5, true, 2},      // Should find log-related content
		{"view", 0.5, false, 2},    // Should find view-related content
		{"command", 0.5, true, 2},  // Should find both (both are commands)
		{"decision", 0.5, true, 1}, // Should find decision-related (only in log)
		{"dashboard", 0.5, false, 1}, // Should find dashboard (only in view)
		{"track", 0.5, true, 1},    // "track" in view content
	}

	for _, tt := range tests {
		t.Run(tt.topic, func(t *testing.T) {
			matches := FilterParagraphs(specs, tt.topic, tt.minScore)
			if len(matches) < tt.minMatches {
				t.Errorf("Topic %q: got %d matches, expected at least %d", tt.topic, len(matches), tt.minMatches)
			}

			// Verify source spec accuracy
			hasLog := false
			for _, match := range matches {
				if match.SpecName == "log.md" {
					hasLog = true
				}
			}
			if hasLog != tt.expectFromLog {
				t.Errorf("Topic %q: log matches expected=%v, got=%v", tt.topic, tt.expectFromLog, hasLog)
			}
		})
	}
}

func TestFilterAccuracyScoreThresholds(t *testing.T) {
	// Test that threshold filtering is accurate
	specs := []Spec{
		{
			Name: "test.md",
			Content: []byte(`# Authentication

User authentication is critical.

The system uses JWT tokens for authentication.

Password hashing with bcrypt.

Database storage of user credentials.`),
		},
	}

	topic := "auth"

	// Lower threshold should return more or equal results
	lowThreshold := FilterParagraphs(specs, topic, 0.3)
	mediumThreshold := FilterParagraphs(specs, topic, 0.5)
	highThreshold := FilterParagraphs(specs, topic, 0.8)

	if len(lowThreshold) < len(mediumThreshold) {
		t.Errorf("Lower threshold should return >= matches: %d vs %d", len(lowThreshold), len(mediumThreshold))
	}
	if len(mediumThreshold) < len(highThreshold) {
		t.Errorf("Medium threshold should return >= matches as high threshold: %d vs %d", len(mediumThreshold), len(highThreshold))
	}

	// All high-threshold matches should be in lower thresholds
	for _, hMatch := range highThreshold {
		found := false
		for _, mMatch := range mediumThreshold {
			if hMatch.Paragraph == mMatch.Paragraph {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("High threshold match not found in medium threshold")
		}
	}
}

func TestFilterAccuracyWithRealSpecs(t *testing.T) {
	// Test with real spec files from the project
	fs := platform.NewRealFileSystem()
	specsPath := filepath.Join("..", "..", ".context", "specs")

	specs, err := DiscoverSpecs(fs, specsPath)
	if err != nil {
		t.Fatalf("DiscoverSpecs() failed: %v", err)
	}

	if len(specs) == 0 {
		t.Fatalf("No specs discovered for accuracy testing")
	}

	tests := []struct {
		topic           string
		minScore        float64
		expectMatches   int
		expectSpecNames []string
	}{
		{"log", 0.5, 2, []string{"log.md"}},
		{"view", 0.5, 2, []string{"view.md"}},
		{"command", 0.5, 4, []string{"log.md", "view.md"}},
	}

	for _, tt := range tests {
		t.Run(tt.topic, func(t *testing.T) {
			matches := FilterParagraphs(specs, tt.topic, tt.minScore)
			if len(matches) < tt.expectMatches {
				t.Errorf("Topic %q: got %d matches, expected at least %d", tt.topic, len(matches), tt.expectMatches)
			}

			// Verify all matches have scores >= minScore
			for _, match := range matches {
				if match.Score < tt.minScore {
					t.Errorf("Match score %.2f below threshold %.2f", match.Score, tt.minScore)
				}
			}

			// Verify at least one match comes from expected specs
			foundExpected := false
			for _, match := range matches {
				for _, expectedSpec := range tt.expectSpecNames {
					if match.SpecName == expectedSpec {
						foundExpected = true
						break
					}
				}
			}
			if len(tt.expectSpecNames) > 0 && !foundExpected {
				t.Errorf("No matches found from expected specs: %v", tt.expectSpecNames)
			}
		})
	}
}

func TestFilterAccuracySortingByScore(t *testing.T) {
	// Test that results are sorted by score (highest first)
	specs := []Spec{
		{
			Name: "test.md",
			Content: []byte(`# Log Command

The log command is useful.

Log entries are recorded.

Record permanent decisions.

Logging is important.`),
		},
	}

	topic := "log"
	matches := FilterParagraphs(specs, topic, 0.3)

	if len(matches) < 2 {
		t.Fatalf("Expected multiple matches, got %d", len(matches))
	}

	// Verify scores are in descending order
	for i := 0; i < len(matches)-1; i++ {
		if matches[i].Score < matches[i+1].Score {
			t.Errorf("Scores not sorted: %.2f < %.2f at positions %d, %d",
				matches[i].Score, matches[i+1].Score, i, i+1)
		}
	}
}

func TestFilterAccuracyPartialMatches(t *testing.T) {
	// Test partial word and substring matching
	specs := []Spec{
		{
			Name: "test.md",
			Content: []byte(`# Authentication

User authentication works.

Authorization levels exist.

Authenticated users proceed.

Auto-authentication is disabled.`),
		},
	}

	topic := "auth"

	matches := FilterParagraphs(specs, topic, 0.4)

	// Should find several paragraphs containing "auth" variations
	if len(matches) < 3 {
		t.Errorf("Expected at least 3 matches for 'auth', got %d", len(matches))
	}

	// Verify all matches contain "auth" (case-insensitive)
	for _, match := range matches {
		if !strings.Contains(strings.ToLower(match.Paragraph), "auth") {
			t.Errorf("Match doesn't contain 'auth': %q", match.Paragraph)
		}
	}
}

func TestFilterAccuracyMultipleSpecsScoring(t *testing.T) {
	// Test that matches from different specs are scored correctly
	specs := []Spec{
		{
			Name: "spec1.md",
			Content: []byte(`# Config

Configuration management.`),
		},
		{
			Name: "spec2.md",
			Content: []byte(`# Configuration

The configuration system is central.

Configure your settings.`),
		},
	}

	topic := "config"
	matches := FilterParagraphs(specs, topic, 0.5)

	// Should have matches from both specs
	spec1Found := false
	spec2Found := false
	for _, match := range matches {
		if match.SpecName == "spec1.md" {
			spec1Found = true
		}
		if match.SpecName == "spec2.md" {
			spec2Found = true
		}
	}

	if !spec1Found {
		t.Errorf("No matches found from spec1.md")
	}
	if !spec2Found {
		t.Errorf("No matches found from spec2.md")
	}
}

func TestFilterAccuracyEmptyParagraphs(t *testing.T) {
	// Test filtering handles specs with many small paragraphs correctly
	specs := []Spec{
		{
			Name: "test.md",
			Content: []byte(`# Title

A

B

Log message here.

C

D

Logging system.`),
		},
	}

	topic := "log"
	matches := FilterParagraphs(specs, topic, 0.5)

	// Should find meaningful paragraphs but filter out single-char ones
	for _, match := range matches {
		if len(strings.TrimSpace(match.Paragraph)) < 3 {
			t.Errorf("Filtered match too short: %q", match.Paragraph)
		}
	}
}
