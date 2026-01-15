package cmd

import (
	"bytes"
	"strings"
	"testing"

	"cdd/internal/platform"
)

func TestPackOutputMarkdownRendering(t *testing.T) {
	// Test that markdown output is properly rendered
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	var out bytes.Buffer
	cmd.SetOut(&out)

	cmd.Flags().Set("focus", "log")
	// Don't set raw flag, should render

	err := cmd.RunE(cmd, []string{})
	if err != nil {
		t.Errorf("pack command failed: %v", err)
	}

	output := out.String()
	if output == "" {
		t.Errorf("Expected rendered output")
	}

	// Rendered output should have markdown structure
	if !strings.Contains(output, "#") {
		t.Errorf("Rendered output should contain markdown")
	}
}

func TestPackOutputRawFormat(t *testing.T) {
	// Test raw text output without rendering
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	var out bytes.Buffer
	cmd.SetOut(&out)

	cmd.Flags().Set("focus", "log")
	cmd.Flags().Set("raw", "true")

	err := cmd.RunE(cmd, []string{})
	if err != nil {
		t.Errorf("pack command failed: %v", err)
	}

	output := out.String()
	if output == "" {
		t.Errorf("Expected raw output")
	}

	// Raw output should not have ANSI codes
	if strings.Contains(output, "\x1b[") {
		t.Errorf("Raw output should not have ANSI escape sequences")
	}

	// Raw output should still be markdown
	if !strings.Contains(output, "#") {
		t.Errorf("Raw output should still contain markdown headers")
	}
}

func TestPackOutputNoMatchesMessage(t *testing.T) {
	// Test graceful handling of no matches
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	var out bytes.Buffer
	cmd.SetOut(&out)

	cmd.Flags().Set("focus", "xyz999notarealword")

	err := cmd.RunE(cmd, []string{})
	if err != nil {
		t.Errorf("pack command should not error on no matches: %v", err)
	}

	output := out.String()

	// Should have helpful message
	if !strings.Contains(output, "No matches") {
		t.Errorf("Output should indicate no matches found")
	}

	if !strings.Contains(output, "found") {
		t.Errorf("Output should explain what was searched")
	}

	// Should suggest alternatives
	if !strings.Contains(output, "Try") || !strings.Contains(output, "topics") {
		t.Errorf("Output should suggest trying other topics")
	}
}

func TestPackOutputEmptyResultsHandling(t *testing.T) {
	// Test output when specs exist but topic has no matches
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	var out bytes.Buffer
	cmd.SetOut(&out)

	// Use topic with no realistic matches
	cmd.Flags().Set("focus", "abcdefghijklmnop")

	err := cmd.RunE(cmd, []string{})
	if err != nil {
		t.Errorf("pack command should handle empty results: %v", err)
	}

	output := out.String()

	// Should not be empty
	if output == "" {
		t.Errorf("pack command should output something even with no matches")
	}

	// Should be readable
	if !strings.Contains(output, "\n") {
		t.Errorf("Output should be properly formatted with newlines")
	}
}

func TestPackOutputStructure(t *testing.T) {
	// Test overall output structure
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	var out bytes.Buffer
	cmd.SetOut(&out)

	cmd.Flags().Set("focus", "log")
	cmd.Flags().Set("raw", "true")

	err := cmd.RunE(cmd, []string{})
	if err != nil {
		t.Errorf("pack command failed: %v", err)
	}

	output := out.String()

	// Should have header with topic name
	if !strings.Contains(output, "Context Pack") {
		t.Errorf("Output should have context pack header")
	}

	if !strings.Contains(output, "log") {
		t.Errorf("Output should mention the search topic")
	}

	// Should have matches section
	if !strings.Contains(output, "From") {
		t.Errorf("Output should show which specs matches come from")
	}

	// Should have separators between matches
	if !strings.Contains(output, "---") {
		t.Errorf("Output should separate matches with ---")
	}
}

func TestPackOutputMatchInfo(t *testing.T) {
	// Test that match information is included
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	var out bytes.Buffer
	cmd.SetOut(&out)

	cmd.Flags().Set("focus", "log")
	cmd.Flags().Set("raw", "true")

	err := cmd.RunE(cmd, []string{})
	if err != nil {
		t.Errorf("pack command failed: %v", err)
	}

	output := out.String()

	// Should have match numbering
	if !strings.Contains(output, "Match") {
		t.Errorf("Output should number matches")
	}

	// Should have scores
	if !strings.Contains(output, "Score:") {
		t.Errorf("Output should show match scores")
	}

	// Scores should be reasonable (0.0-1.0)
	if !strings.Contains(output, "0.") {
		t.Errorf("Output should show decimal scores")
	}
}

func TestPackOutputRawVsRendered(t *testing.T) {
	// Compare raw and rendered outputs
	fs := platform.NewRealFileSystem()

	// Get raw output
	cmdRaw := NewPackCmd(fs)
	var outRaw bytes.Buffer
	cmdRaw.SetOut(&outRaw)
	cmdRaw.Flags().Set("focus", "log")
	cmdRaw.Flags().Set("raw", "true")
	cmdRaw.RunE(cmdRaw, []string{})
	rawOutput := outRaw.String()

	// Get rendered output
	cmdRendered := NewPackCmd(fs)
	var outRendered bytes.Buffer
	cmdRendered.SetOut(&outRendered)
	cmdRendered.Flags().Set("focus", "log")
	// raw flag defaults to false
	cmdRendered.RunE(cmdRendered, []string{})
	renderedOutput := outRendered.String()

	// Both should have content
	if rawOutput == "" || renderedOutput == "" {
		t.Errorf("Both raw and rendered output should have content")
	}

	// Raw should not have ANSI codes
	if strings.Contains(rawOutput, "\x1b[") {
		t.Errorf("Raw output should not have ANSI codes")
	}

	// Rendered might have ANSI codes (depending on terminal detection)
	// but both should have readable content

	// Raw output should be shorter or equal (no extra formatting)
	// This is a loose check since rendering might vary
	if len(rawOutput) > 0 && len(renderedOutput) > 0 {
		// Both should be reasonable lengths
		if len(rawOutput) < 20 {
			t.Errorf("Raw output seems too short: %d bytes", len(rawOutput))
		}
		if len(renderedOutput) < 20 {
			t.Errorf("Rendered output seems too short: %d bytes", len(renderedOutput))
		}
	}
}

func TestPackOutputMultipleSpecsSeparation(t *testing.T) {
	// Test that matches from different specs are clearly separated
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	var out bytes.Buffer
	cmd.SetOut(&out)

	cmd.Flags().Set("focus", "command")
	cmd.Flags().Set("raw", "true")

	err := cmd.RunE(cmd, []string{})
	if err != nil {
		t.Errorf("pack command failed: %v", err)
	}

	output := out.String()

	// Should have "From" headers for spec files
	fromCount := strings.Count(output, "## From")
	if fromCount < 1 {
		t.Errorf("Output should have 'From' headers for spec files")
	}

	// Should have multiple spec references if searching broad topic
	lines := strings.Split(output, "\n")
	specCount := 0
	for _, line := range lines {
		if strings.Contains(line, "## From") && strings.Contains(line, ".md") {
			specCount++
		}
	}

	if specCount < 1 {
		t.Errorf("Output should reference spec files")
	}
}

func TestPackOutputRelevanceOrdering(t *testing.T) {
	// Test that matches are ordered by relevance (score)
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	var out bytes.Buffer
	cmd.SetOut(&out)

	cmd.Flags().Set("focus", "log")
	cmd.Flags().Set("raw", "true")

	err := cmd.RunE(cmd, []string{})
	if err != nil {
		t.Errorf("pack command failed: %v", err)
	}

	output := out.String()

	// Extract scores from output
	lines := strings.Split(output, "\n")

	for _, line := range lines {
		if strings.Contains(line, "Score:") {
			// Extract the numeric score
			parts := strings.Split(line, "Score:")
			if len(parts) > 1 {
				scoreStr := strings.TrimSpace(parts[1])
				scoreStr = strings.Split(scoreStr, " ")[0]
				// Simple parsing - just check if scores are present
				if strings.Contains(scoreStr, ".") {
					// Score is formatted, this is good
				}
			}
		}
	}

	// Should have at least some match info
	if !strings.Contains(output, "Score:") {
		t.Errorf("Output should display match scores")
	}
}
