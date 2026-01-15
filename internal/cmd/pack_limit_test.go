package cmd

import (
	"testing"
)

func TestLimitResults(t *testing.T) {
	tests := []struct {
		name        string
		matches     []ParagraphMatch
		limit       int
		expectCount int
		expectError bool
	}{
		{
			name: "no limit (negative)",
			matches: []ParagraphMatch{
				{Paragraph: "para1", SpecName: "spec1.md", Score: 0.9},
				{Paragraph: "para2", SpecName: "spec2.md", Score: 0.8},
				{Paragraph: "para3", SpecName: "spec3.md", Score: 0.7},
			},
			limit:       -1,
			expectCount: 3,
			expectError: false,
		},
		{
			name: "limit to 2",
			matches: []ParagraphMatch{
				{Paragraph: "para1", SpecName: "spec1.md", Score: 0.9},
				{Paragraph: "para2", SpecName: "spec2.md", Score: 0.8},
				{Paragraph: "para3", SpecName: "spec3.md", Score: 0.7},
			},
			limit:       2,
			expectCount: 2,
			expectError: false,
		},
		{
			name: "limit 0 (no content)",
			matches: []ParagraphMatch{
				{Paragraph: "para1", SpecName: "spec1.md", Score: 0.9},
				{Paragraph: "para2", SpecName: "spec2.md", Score: 0.8},
			},
			limit:       0,
			expectCount: 0,
			expectError: false,
		},
		{
			name: "limit higher than available",
			matches: []ParagraphMatch{
				{Paragraph: "para1", SpecName: "spec1.md", Score: 0.9},
			},
			limit:       10,
			expectCount: 1,
			expectError: false,
		},
		{
			name:        "empty matches",
			matches:     []ParagraphMatch{},
			limit:       5,
			expectCount: 0,
			expectError: false,
		},
		{
			name: "single match with limit 1",
			matches: []ParagraphMatch{
				{Paragraph: "para1", SpecName: "spec1.md", Score: 0.9},
			},
			limit:       1,
			expectCount: 1,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := LimitResults(tt.matches, tt.limit)

			if tt.expectError && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if len(result) != tt.expectCount {
				t.Errorf("Expected %d results, got %d", tt.expectCount, len(result))
			}

			// Verify order is preserved
			for i := 0; i < len(result); i++ {
				if result[i] != tt.matches[i] {
					t.Errorf("Result order not preserved at index %d", i)
				}
			}
		})
	}
}

func TestLimitResultsPreservesScores(t *testing.T) {
	// Verify that limiting preserves the original scores
	matches := []ParagraphMatch{
		{Paragraph: "para1", SpecName: "spec1.md", Score: 0.95},
		{Paragraph: "para2", SpecName: "spec2.md", Score: 0.87},
		{Paragraph: "para3", SpecName: "spec3.md", Score: 0.72},
	}

	result, err := LimitResults(matches, 2)
	if err != nil {
		t.Fatalf("LimitResults failed: %v", err)
	}

	if len(result) != 2 {
		t.Errorf("Expected 2 results, got %d", len(result))
	}

	if result[0].Score != 0.95 {
		t.Errorf("First result score should be 0.95, got %f", result[0].Score)
	}
	if result[1].Score != 0.87 {
		t.Errorf("Second result score should be 0.87, got %f", result[1].Score)
	}
}
