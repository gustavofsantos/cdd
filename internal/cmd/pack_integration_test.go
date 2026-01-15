package cmd

import (
	"bytes"
	"strings"
	"testing"

	"cdd/internal/platform"
)

func TestPackCmdIntegrationBasic(t *testing.T) {
	// Integration test: pack command with real specs
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	var out bytes.Buffer
	cmd.SetOut(&out)

	if err := cmd.Flags().Set("focus", "log"); err != nil {
		t.Fatalf("Failed to set focus flag: %v", err)
	}

	err := cmd.RunE(cmd, []string{})
	if err != nil {
		t.Errorf("pack command integration failed: %v", err)
	}

	output := out.String()
	if output == "" {
		t.Errorf("Expected output from pack command")
	}

	// Should contain the topic in output
	if !strings.Contains(output, "log") && !strings.Contains(output, "Log") {
		t.Errorf("Output should contain topic 'log'")
	}

	// Should have context pack header
	if !strings.Contains(output, "Context Pack") {
		t.Errorf("Output should contain 'Context Pack' header")
	}
}

func TestPackCmdIntegrationMultipleMatches(t *testing.T) {
	// Integration test: verify multiple matches are returned and formatted
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	var out bytes.Buffer
	cmd.SetOut(&out)

	if err := cmd.Flags().Set("focus", "command"); err != nil {
		t.Fatalf("Failed to set focus flag: %v", err)
	}

	err := cmd.RunE(cmd, []string{})
	if err != nil {
		t.Errorf("pack command failed: %v", err)
	}

	output := out.String()

	// Should mention number of matches
	if !strings.Contains(output, "relevant paragraphs") {
		t.Errorf("Output should mention number of paragraphs found")
	}

	// Should have spec file headers
	if !strings.Contains(output, "From") {
		t.Errorf("Output should have 'From' spec headers")
	}

	// Should have match separators
	if !strings.Contains(output, "---") {
		t.Errorf("Output should have match separators")
	}
}

func TestPackCmdIntegrationScoreDisplay(t *testing.T) {
	// Integration test: verify match scores are displayed
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	var out bytes.Buffer
	cmd.SetOut(&out)

	if err := cmd.Flags().Set("focus", "log"); err != nil {
		t.Fatalf("Failed to set focus flag: %v", err)
	}

	err := cmd.RunE(cmd, []string{})
	if err != nil {
		t.Errorf("pack command failed: %v", err)
	}

	output := out.String()

	// Should display scores
	if !strings.Contains(output, "Score:") {
		t.Errorf("Output should display match scores")
	}
}

func TestPackCmdIntegrationRawOutput(t *testing.T) {
	// Integration test: verify --raw produces plain text
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	var out bytes.Buffer
	cmd.SetOut(&out)

	if err := cmd.Flags().Set("focus", "view"); err != nil {
		t.Fatalf("Failed to set focus flag: %v", err)
	}
	if err := cmd.Flags().Set("raw", "true"); err != nil {
		t.Fatalf("Failed to set raw flag: %v", err)
	}

	err := cmd.RunE(cmd, []string{})
	if err != nil {
		t.Errorf("pack command with --raw failed: %v", err)
	}

	output := out.String()

	// Should not have ANSI codes
	if strings.Contains(output, "\x1b[") {
		t.Errorf("Raw output should not have ANSI escape codes")
	}

	// Should still have content
	if !strings.Contains(output, "View") && !strings.Contains(output, "view") {
		t.Errorf("Raw output should contain search term")
	}
}

func TestPackCmdIntegrationNoMatchesMessage(t *testing.T) {
	// Integration test: graceful handling of no matches
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	var out bytes.Buffer
	cmd.SetOut(&out)

	// Use a topic unlikely to match anything
	if err := cmd.Flags().Set("focus", "xyzzyzzzzqqqq9999"); err != nil {
		t.Fatalf("Failed to set focus flag: %v", err)
	}

	err := cmd.RunE(cmd, []string{})
	if err != nil {
		t.Errorf("pack command should not error on no matches: %v", err)
	}

	output := out.String()

	// Should indicate no matches
	if !strings.Contains(output, "No matches") && !strings.Contains(output, "no matches") {
		t.Errorf("Output should indicate no matches found, got: %s", output[:min(100, len(output))])
	}
}

func TestPackCmdIntegrationTopicVariations(t *testing.T) {
	// Integration test: test various topic searches
	fs := platform.NewRealFileSystem()

	topics := []struct {
		topic        string
		expectOutput string
	}{
		{"log", "log"},
		{"view", "view"},
		{"command", "command"},
		{"specification", ""},
	}

	for _, tt := range topics {
		t.Run(tt.topic, func(t *testing.T) {
			cmd := NewPackCmd(fs)
			var out bytes.Buffer
			cmd.SetOut(&out)

			if err := cmd.Flags().Set("focus", tt.topic); err != nil {
				t.Fatalf("Failed to set focus flag: %v", err)
			}

			err := cmd.RunE(cmd, []string{})
			if err != nil {
				t.Errorf("pack command for topic %q failed: %v", tt.topic, err)
			}

			output := out.String()
			if output == "" {
				t.Errorf("pack command for topic %q produced no output", tt.topic)
			}

			// Should always have the topic in output
			if !strings.Contains(strings.ToLower(output), strings.ToLower(tt.topic)) {
				t.Errorf("Output for topic %q should mention the topic", tt.topic)
			}
		})
	}
}

func TestPackCmdIntegrationOutputFormatting(t *testing.T) {
	// Integration test: verify output is properly formatted markdown
	fs := platform.NewRealFileSystem()
	cmd := NewPackCmd(fs)

	var out bytes.Buffer
	cmd.SetOut(&out)

	if err := cmd.Flags().Set("focus", "log"); err != nil {
		t.Fatalf("Failed to set focus flag: %v", err)
	}
	if err := cmd.Flags().Set("raw", "true"); err != nil {
		t.Fatalf("Failed to set raw flag: %v", err)
	}

	err := cmd.RunE(cmd, []string{})
	if err != nil {
		t.Errorf("pack command failed: %v", err)
	}

	output := out.String()

	// Should have markdown headers
	if !strings.Contains(output, "#") {
		t.Errorf("Output should contain markdown headers")
	}

	// Should have proper structure
	lines := strings.Split(output, "\n")
	if len(lines) < 5 {
		t.Errorf("Output should have meaningful length, got %d lines", len(lines))
	}
}

func TestPackCmdIntegrationConsistency(t *testing.T) {
	// Integration test: verify consistent results on multiple runs
	fs := platform.NewRealFileSystem()
	topic := "log"

	var out1 bytes.Buffer
	cmd1 := NewPackCmd(fs)
	cmd1.SetOut(&out1)
	if err := cmd1.Flags().Set("focus", topic); err != nil {
		t.Fatalf("Failed to set focus flag: %v", err)
	}
	if err := cmd1.RunE(cmd1, []string{}); err != nil {
		t.Errorf("pack command failed: %v", err)
	}

	var out2 bytes.Buffer
	cmd2 := NewPackCmd(fs)
	cmd2.SetOut(&out2)
	if err := cmd2.Flags().Set("focus", topic); err != nil {
		t.Fatalf("Failed to set focus flag: %v", err)
	}
	if err := cmd2.RunE(cmd2, []string{}); err != nil {
		t.Errorf("pack command failed: %v", err)
	}

	output1 := out1.String()
	output2 := out2.String()

	// Should produce identical output on consecutive runs
	if output1 != output2 {
		t.Errorf("pack command should produce consistent results across runs")
	}

	// Count matches in both outputs
	count1 := strings.Count(output1, "Match")
	count2 := strings.Count(output2, "Match")

	if count1 != count2 {
		t.Errorf("Match counts differ between runs: %d vs %d", count1, count2)
	}
}

func TestPackCmdIntegrationRealWorldScenario(t *testing.T) {
	// Integration test: realistic workflow scenario
	fs := platform.NewRealFileSystem()

	scenarios := []struct {
		name  string
		topic string
	}{
		{"developer learning about log command", "log"},
		{"understanding view functionality", "view"},
		{"command documentation", "command"},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			cmd := NewPackCmd(fs)
			var out bytes.Buffer
			cmd.SetOut(&out)

			if err := cmd.Flags().Set("focus", scenario.topic); err != nil {
				t.Fatalf("Failed to set focus flag: %v", err)
			}

			err := cmd.RunE(cmd, []string{})
			if err != nil {
				t.Errorf("Scenario %q failed: %v", scenario.name, err)
			}

			output := out.String()

			// Verify useful output is returned
			if len(output) < 50 {
				t.Errorf("Scenario %q: output too short (%d chars)", scenario.name, len(output))
			}

			// Verify output structure
			if !strings.Contains(output, "#") {
				t.Errorf("Scenario %q: output should have markdown structure", scenario.name)
			}

			if !strings.Contains(output, "Score:") && !strings.Contains(output, "No matches") {
				t.Errorf("Scenario %q: output should have scores or no-match message", scenario.name)
			}
		})
	}
}
