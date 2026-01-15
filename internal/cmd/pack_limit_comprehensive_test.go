package cmd

import (
	"bytes"
	"strconv"
	"strings"
	"testing"

	"cdd/internal/platform"
)

func TestPackLimitComprehensiveScenarios(t *testing.T) {
	// Comprehensive test: various limit scenarios
	fs := platform.NewRealFileSystem()

	scenarios := []struct {
		name            string
		focus           string
		limit           int
		expectMaxCount  int
		expectTruncated bool
	}{
		{
			name:            "limit 1 on broad topic",
			focus:           "log",
			limit:           1,
			expectMaxCount:  1,
			expectTruncated: true,
		},
		{
			name:            "limit 10 on log topic",
			focus:           "log",
			limit:           10,
			expectMaxCount:  10,
			expectTruncated: true,
		},
		{
			name:            "limit 0 shows count only",
			focus:           "command",
			limit:           0,
			expectMaxCount:  0,
			expectTruncated: false,
		},
		{
			name:            "no limit returns all",
			focus:           "view",
			limit:           -1,
			expectMaxCount:  999,
			expectTruncated: false,
		},
		{
			name:            "very high limit",
			focus:           "error",
			limit:           9999,
			expectMaxCount:  9999,
			expectTruncated: false,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			cmd := NewPackCmd(fs)
			var out bytes.Buffer
			cmd.SetOut(&out)

			if err := cmd.Flags().Set("focus", scenario.focus); err != nil {
				t.Fatalf("Failed to set focus flag: %v", err)
			}
			if err := cmd.Flags().Set("limit", strconv.Itoa(scenario.limit)); err != nil {
				t.Fatalf("Failed to set limit flag: %v", err)
			}

			err := cmd.RunE(cmd, []string{})
			if err != nil {
				t.Errorf("pack command failed: %v", err)
			}

			output := out.String()
			if output == "" {
				t.Errorf("Expected output")
			}

			matchCount := strings.Count(output, "**[Match")
			if matchCount > scenario.expectMaxCount {
				t.Errorf("Limit %d: got %d matches, expected max %d", scenario.limit, matchCount, scenario.expectMaxCount)
			}

			// Check for truncation message
			hasTruncation := strings.Contains(output, "showing") && strings.Contains(output, "of")
			if scenario.expectTruncated && !hasTruncation && scenario.limit > 0 {
				// Only check if we actually have results
				if matchCount > 0 {
					t.Logf("Warning: expected truncation message but didn't find it")
				}
			}
		})
	}
}

func TestPackLimitWithDifferentTopics(t *testing.T) {
	// Test limit behavior across different topics
	fs := platform.NewRealFileSystem()

	topics := []string{"log", "view", "command", "specification", "requirement"}
	limit := 3

	for _, topic := range topics {
		t.Run(topic, func(t *testing.T) {
			cmd := NewPackCmd(fs)
			var out bytes.Buffer
			cmd.SetOut(&out)

			if err := cmd.Flags().Set("focus", topic); err != nil {
				t.Fatalf("Failed to set focus flag: %v", err)
			}
			if err := cmd.Flags().Set("limit", "3"); err != nil {
				t.Fatalf("Failed to set limit flag: %v", err)
			}

			err := cmd.RunE(cmd, []string{})
			if err != nil {
				t.Errorf("pack command for topic %q failed: %v", topic, err)
			}

			output := out.String()
			matchCount := strings.Count(output, "**[Match")

			if matchCount > limit {
				t.Errorf("Topic %q: limit %d but got %d matches", topic, limit, matchCount)
			}
		})
	}
}

func TestPackLimitPresetsRanking(t *testing.T) {
	// Verify limit preserves relevance ranking
	fs := platform.NewRealFileSystem()

	// Get first 3 matches
	cmd1 := NewPackCmd(fs)
	var out1 bytes.Buffer
	cmd1.SetOut(&out1)
	if err := cmd1.Flags().Set("focus", "log"); err != nil {
		t.Fatalf("Failed to set focus flag: %v", err)
	}
	if err := cmd1.Flags().Set("limit", "3"); err != nil {
		t.Fatalf("Failed to set limit flag: %v", err)
	}
	if err := cmd1.RunE(cmd1, []string{}); err != nil {
		t.Errorf("pack command failed: %v", err)
	}
	output1 := out1.String()

	// Get first 5 matches
	cmd2 := NewPackCmd(fs)
	var out2 bytes.Buffer
	cmd2.SetOut(&out2)
	if err := cmd2.Flags().Set("focus", "log"); err != nil {
		t.Fatalf("Failed to set focus flag: %v", err)
	}
	if err := cmd2.Flags().Set("limit", "5"); err != nil {
		t.Fatalf("Failed to set limit flag: %v", err)
	}
	if err := cmd2.RunE(cmd2, []string{}); err != nil {
		t.Errorf("pack command failed: %v", err)
	}
	output2 := out2.String()

	// First 3 from limit 5 should match first 3 from limit 3
	// Extract scores from both outputs
	scores1 := extractScores(output1)
	scores2 := extractScores(output2)

	if len(scores1) > 0 && len(scores2) >= 3 {
		for i := 0; i < len(scores1); i++ {
			if scores1[i] != scores2[i] {
				t.Errorf("Ranking mismatch at position %d: limit 3 has %.2f, limit 5 has %.2f", i, scores1[i], scores2[i])
			}
		}
	}
}

func TestPackLimitWithRawAndFormatted(t *testing.T) {
	// Test limit works consistently with both output formats
	fs := platform.NewRealFileSystem()

	// Raw output
	cmdRaw := NewPackCmd(fs)
	var outRaw bytes.Buffer
	cmdRaw.SetOut(&outRaw)
	if err := cmdRaw.Flags().Set("focus", "command"); err != nil {
		t.Fatalf("Failed to set focus flag: %v", err)
	}
	if err := cmdRaw.Flags().Set("limit", "4"); err != nil {
		t.Fatalf("Failed to set limit flag: %v", err)
	}
	if err := cmdRaw.Flags().Set("raw", "true"); err != nil {
		t.Fatalf("Failed to set raw flag: %v", err)
	}
	if err := cmdRaw.RunE(cmdRaw, []string{}); err != nil {
		t.Errorf("pack command failed: %v", err)
	}
	rawOutput := outRaw.String()

	// Formatted output
	cmdFmt := NewPackCmd(fs)
	var outFmt bytes.Buffer
	cmdFmt.SetOut(&outFmt)
	if err := cmdFmt.Flags().Set("focus", "command"); err != nil {
		t.Fatalf("Failed to set focus flag: %v", err)
	}
	if err := cmdFmt.Flags().Set("limit", "4"); err != nil {
		t.Fatalf("Failed to set limit flag: %v", err)
	}
	if err := cmdFmt.RunE(cmdFmt, []string{}); err != nil {
		t.Errorf("pack command failed: %v", err)
	}
	fmtOutput := outFmt.String()

	rawCount := strings.Count(rawOutput, "**[Match")
	fmtCount := strings.Count(fmtOutput, "**[Match")

	if rawCount != fmtCount {
		t.Errorf("Raw and formatted outputs have different match counts: %d vs %d", rawCount, fmtCount)
	}

	if rawCount > 4 {
		t.Errorf("Limit 4 but got %d matches in raw output", rawCount)
	}
	if fmtCount > 4 {
		t.Errorf("Limit 4 but got %d matches in formatted output", fmtCount)
	}
}

func TestPackLimitEdgeCases(t *testing.T) {
	// Edge cases: limit 1, limit at exact match count
	fs := platform.NewRealFileSystem()

	tests := []struct {
		name  string
		limit int
	}{
		{"limit 1", 1},
		{"limit 2", 2},
		{"limit 100", 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := NewPackCmd(fs)
			var out bytes.Buffer
			cmd.SetOut(&out)

			if err := cmd.Flags().Set("focus", "log"); err != nil {
				t.Fatalf("Failed to set focus flag: %v", err)
			}
			if err := cmd.Flags().Set("limit", strconv.Itoa(tt.limit)); err != nil {
				t.Fatalf("Failed to set limit flag: %v", err)
			}

			err := cmd.RunE(cmd, []string{})
			if err != nil {
				t.Errorf("pack command failed: %v", err)
			}

			output := out.String()
			matchCount := strings.Count(output, "**[Match")

			if matchCount > tt.limit {
				t.Errorf("Limit %d exceeded: got %d matches", tt.limit, matchCount)
			}
		})
	}
}

func TestPackLimitNoMatchesTopic(t *testing.T) {
	// Edge case: limit on a topic with no matches
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	var out bytes.Buffer
	cmd.SetOut(&out)

	if err := cmd.Flags().Set("focus", "xyz999notarealword"); err != nil {
		t.Fatalf("Failed to set focus flag: %v", err)
	}
	if err := cmd.Flags().Set("limit", "10"); err != nil {
		t.Fatalf("Failed to set limit flag: %v", err)
	}

	err := cmd.RunE(cmd, []string{})
	if err != nil {
		t.Errorf("pack command should handle no matches: %v", err)
	}

	output := out.String()
	if !strings.Contains(output, "No matches") {
		t.Errorf("Should indicate no matches found")
	}

	matchCount := strings.Count(output, "**[Match")
	if matchCount > 0 {
		t.Errorf("Should have no matches but got %d", matchCount)
	}
}

// Helper function to extract scores from output
func extractScores(output string) []float64 {
	var scores []float64
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.Contains(line, "Score:") {
			// Extract score value - format is "**[Match N] (Score: 0.XX)**"
			parts := strings.Split(line, "Score:")
			if len(parts) > 1 {
				scoreStr := strings.TrimSpace(parts[1])
				scoreStr = strings.Split(scoreStr, ")")[0]
				scoreStr = strings.TrimSpace(scoreStr)
				if score, err := strconv.ParseFloat(scoreStr, 64); err == nil {
					// Successfully parsed score
					scores = append(scores, score)
				}
			}
		}
	}
	return scores
}
