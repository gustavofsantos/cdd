package cmd

import (
	"testing"
)

func TestFilterParagraphs(t *testing.T) {
	tests := []struct {
		name           string
		specs          []Spec
		topic          string
		minScore       float64
		expectMatches  int
		expectSpecName string
	}{
		{
			name: "filter by log topic",
			specs: []Spec{
				{
					Name: "log.md",
					Content: []byte(`# Log Command Specification

## Overview
The log command records permanent decisions.

## Requirements
When the user invokes cdd log, the system shall record messages.`),
				},
			},
			topic:         "log",
			minScore:      0.5,
			expectMatches: 3, // overview + requirements + message
		},
		{
			name: "filter with no matches",
			specs: []Spec{
				{
					Name: "log.md",
					Content: []byte(`# Log Command

Content about logging decisions.`),
				},
			},
			topic:         "authentication",
			minScore:      0.5,
			expectMatches: 0,
		},
		{
			name: "filter across multiple specs",
			specs: []Spec{
				{
					Name: "log.md",
					Content: []byte(`# Log Command
The log command records decisions.`),
				},
				{
					Name: "view.md",
					Content: []byte(`# View Command
The view command displays tracks.`),
				},
			},
			topic:         "command",
			minScore:      0.5,
			expectMatches: 2, // one from each spec
		},
		{
			name: "filter with high score threshold",
			specs: []Spec{
				{
					Name: "log.md",
					Content: []byte(`# Log Command

The log command records decisions and errors.

## Requirements
The system shall record messages to file.`),
				},
			},
			topic:         "log",
			minScore:      0.85,
			expectMatches: 2, // both title and first content paragraph match at high threshold
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matches := FilterParagraphs(tt.specs, tt.topic, tt.minScore)
			if len(matches) != tt.expectMatches {
				t.Errorf("FilterParagraphs() returned %d matches, expected %d", len(matches), tt.expectMatches)
				for i, m := range matches {
					t.Logf("  Match %d: %q (score: %.2f) from %s", i+1, m.Paragraph[:min(50, len(m.Paragraph))], m.Score, m.SpecName)
				}
			}
		})
	}
}

func TestFilterParagraphsEmptySpecs(t *testing.T) {
	matches := FilterParagraphs([]Spec{}, "query", 0.5)
	if len(matches) != 0 {
		t.Errorf("FilterParagraphs with empty specs should return no matches, got %d", len(matches))
	}
}

func TestFilterParagraphsEmptyTopic(t *testing.T) {
	specs := []Spec{
		{
			Name:    "test.md",
			Content: []byte("Some content"),
		},
	}
	matches := FilterParagraphs(specs, "", 0.5)
	if len(matches) != 0 {
		t.Errorf("FilterParagraphs with empty topic should return no matches, got %d", len(matches))
	}
}

func TestParagraphMatchStruct(t *testing.T) {
	match := ParagraphMatch{
		Paragraph: "Test paragraph",
		SpecName:  "test.md",
		Score:     0.85,
	}

	if match.Paragraph != "Test paragraph" {
		t.Errorf("Paragraph field incorrect")
	}
	if match.SpecName != "test.md" {
		t.Errorf("SpecName field incorrect")
	}
	if match.Score != 0.85 {
		t.Errorf("Score field incorrect")
	}
}

// Helper function for tests
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
