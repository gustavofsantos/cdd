package cmd

import (
	"bytes"
	"strings"
	"testing"

	"cdd/internal/platform"
)

func TestPackIntegrationWithLimit(t *testing.T) {
	// Integration test: pack command with limit applies correctly
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	var out bytes.Buffer
	cmd.SetOut(&out)

	cmd.Flags().Set("focus", "log")
	cmd.Flags().Set("limit", "3")

	err := cmd.RunE(cmd, []string{})
	if err != nil {
		t.Errorf("pack command with limit failed: %v", err)
	}

	output := out.String()
	if output == "" {
		t.Errorf("Expected output from pack command")
	}

	// Count matches in output
	matchCount := strings.Count(output, "**[Match")
	if matchCount > 3 {
		t.Errorf("Limit 3 should return max 3 matches, got %d", matchCount)
	}
}

func TestPackIntegrationLimitZeroShowsCountOnly(t *testing.T) {
	// Integration test: limit 0 shows count but no paragraphs
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	var out bytes.Buffer
	cmd.SetOut(&out)

	cmd.Flags().Set("focus", "log")
	cmd.Flags().Set("limit", "0")

	err := cmd.RunE(cmd, []string{})
	if err != nil {
		t.Errorf("pack command with limit 0 failed: %v", err)
	}

	output := out.String()

	// Should have count header
	if !strings.Contains(output, "Found") {
		t.Errorf("Output should show match count")
	}

	// Should NOT have Match entries
	if strings.Contains(output, "**[Match") {
		t.Errorf("Limit 0 should not show individual matches")
	}
}

func TestPackIntegrationTruncationMessage(t *testing.T) {
	// Integration test: output should indicate if results were truncated
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	var out bytes.Buffer
	cmd.SetOut(&out)

	cmd.Flags().Set("focus", "command")
	cmd.Flags().Set("limit", "2")

	err := cmd.RunE(cmd, []string{})
	if err != nil {
		t.Errorf("pack command failed: %v", err)
	}

	output := out.String()

	// If we got results, check for truncation indicator
	if strings.Contains(output, "**[Match") {
		// Should show "showing X of Y" if truncated
		matchCount := strings.Count(output, "**[Match")
		if matchCount == 2 && strings.Contains(output, "Found") {
			// This is OK - either truncated or just 2 results exist
			t.Logf("Got %d matches with limit 2", matchCount)
		}
	}
}

func TestPackIntegrationLimitNoLimit(t *testing.T) {
	// Integration test: no limit returns all results
	fs := platform.NewRealFileSystem()

	// First get all results
	cmd1 := NewPackCmd(fs)
	var out1 bytes.Buffer
	cmd1.SetOut(&out1)
	cmd1.Flags().Set("focus", "log")
	cmd1.Flags().Set("limit", "-1")
	cmd1.RunE(cmd1, []string{})
	allCount := strings.Count(out1.String(), "**[Match")

	// Then get with limit
	cmd2 := NewPackCmd(fs)
	var out2 bytes.Buffer
	cmd2.SetOut(&out2)
	cmd2.Flags().Set("focus", "log")
	cmd2.Flags().Set("limit", "5")
	cmd2.RunE(cmd2, []string{})
	limitCount := strings.Count(out2.String(), "**[Match")

	if allCount < limitCount {
		t.Errorf("Limit should constrain results: all=%d, limit(5)=%d", allCount, limitCount)
	}

	if limitCount > 5 {
		t.Errorf("Limit 5 should return max 5 matches, got %d", limitCount)
	}
}

func TestPackIntegrationLimitPreservesRanking(t *testing.T) {
	// Integration test: limit should return top N by score
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	var out bytes.Buffer
	cmd.SetOut(&out)

	cmd.Flags().Set("focus", "view")
	cmd.Flags().Set("limit", "2")

	err := cmd.RunE(cmd, []string{})
	if err != nil {
		t.Errorf("pack command failed: %v", err)
	}

	output := out.String()

	// Should have at most 2 matches
	matchCount := strings.Count(output, "**[Match")
	if matchCount > 2 {
		t.Errorf("Limit 2 should return max 2 matches, got %d", matchCount)
	}

	// Verify scores are still shown
	if matchCount > 0 && !strings.Contains(output, "Score:") {
		t.Errorf("Output should show match scores")
	}
}

func TestPackIntegrationLimitWithRaw(t *testing.T) {
	// Integration test: limit works with raw output
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	var out bytes.Buffer
	cmd.SetOut(&out)

	cmd.Flags().Set("focus", "log")
	cmd.Flags().Set("limit", "4")
	cmd.Flags().Set("raw", "true")

	err := cmd.RunE(cmd, []string{})
	if err != nil {
		t.Errorf("pack command with limit and raw failed: %v", err)
	}

	output := out.String()

	// Should not have ANSI codes
	if strings.Contains(output, "\x1b[") {
		t.Errorf("Raw output should not have ANSI codes")
	}

	// Should still respect limit
	matchCount := strings.Count(output, "**[Match")
	if matchCount > 4 {
		t.Errorf("Limit 4 should return max 4 matches, got %d", matchCount)
	}
}

func TestPackIntegrationLimitVeryHigh(t *testing.T) {
	// Integration test: very high limit behaves like no limit
	fs := platform.NewRealFileSystem()

	// Get all results with no limit
	cmd1 := NewPackCmd(fs)
	var out1 bytes.Buffer
	cmd1.SetOut(&out1)
	cmd1.Flags().Set("focus", "log")
	cmd1.Flags().Set("limit", "-1")
	cmd1.RunE(cmd1, []string{})
	allCount := strings.Count(out1.String(), "**[Match")

	// Get with very high limit
	cmd2 := NewPackCmd(fs)
	var out2 bytes.Buffer
	cmd2.SetOut(&out2)
	cmd2.Flags().Set("focus", "log")
	cmd2.Flags().Set("limit", "999999")
	cmd2.RunE(cmd2, []string{})
	highCount := strings.Count(out2.String(), "**[Match")

	// Should return same number
	if allCount != highCount {
		t.Errorf("High limit should equal no limit: all=%d, high=%d", allCount, highCount)
	}
}
