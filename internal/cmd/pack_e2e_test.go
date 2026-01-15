package cmd

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"cdd/internal/platform"
)

func TestPackE2EWorkflow(t *testing.T) {
	// End-to-end test: complete workflow from command execution to output
	fs := platform.NewRealFileSystem()

	// Scenario: Developer learning about the log command
	cmd := NewPackCmd(fs)
	var out bytes.Buffer
	cmd.SetOut(&out)

	// User searches for "log" topic
	if err := cmd.Flags().Set("focus", "log"); err != nil {
		t.Fatalf("Failed to set focus flag: %v", err)
	}

	// Execute the command
	err := cmd.RunE(cmd, []string{})
	if err != nil {
		t.Fatalf("E2E workflow failed: %v", err)
	}

	output := out.String()

	// Verify entire workflow result
	if output == "" {
		t.Fatalf("E2E workflow produced no output")
	}

	// Verify output contains expected elements
	assertions := []struct {
		contains string
		name     string
	}{
		{"Context Pack", "header"},
		{"log", "topic"},
		{"relevant paragraphs", "match count"},
		{"Score:", "relevance scores"},
		{"From", "spec source"},
	}

	for _, assertion := range assertions {
		if !strings.Contains(strings.ToLower(output), strings.ToLower(assertion.contains)) {
			t.Errorf("E2E workflow: missing %s - output should contain %q", assertion.name, assertion.contains)
		}
	}
}

func TestPackE2EMultipleQueries(t *testing.T) {
	// End-to-end test: sequential queries
	fs := platform.NewRealFileSystem()

	queries := []string{"log", "view", "command"}

	for i, query := range queries {
		cmd := NewPackCmd(fs)
		var out bytes.Buffer
		cmd.SetOut(&out)

		if err := cmd.Flags().Set("focus", query); err != nil {
			t.Fatalf("Failed to set focus flag: %v", err)
		}

		err := cmd.RunE(cmd, []string{})
		if err != nil {
			t.Errorf("E2E query %d (%q) failed: %v", i+1, query, err)
		}

		output := out.String()
		if output == "" {
			t.Errorf("E2E query %d (%q) produced no output", i+1, query)
		}

		// Each query should mention the topic
		if !strings.Contains(strings.ToLower(output), strings.ToLower(query)) {
			t.Errorf("E2E query %d (%q) output should mention the query", i+1, query)
		}
	}
}

func TestPackE2ERawAndFormattedOutput(t *testing.T) {
	// End-to-end test: same query with different output modes
	fs := platform.NewRealFileSystem()

	// Query with raw output
	cmdRaw := NewPackCmd(fs)
	var outRaw bytes.Buffer
	cmdRaw.SetOut(&outRaw)
	if err := cmdRaw.Flags().Set("focus", "log"); err != nil {
		t.Fatalf("Failed to set focus flag: %v", err)
	}
	if err := cmdRaw.Flags().Set("raw", "true"); err != nil {
		t.Fatalf("Failed to set raw flag: %v", err)
	}

	err := cmdRaw.RunE(cmdRaw, []string{})
	if err != nil {
		t.Errorf("E2E raw output failed: %v", err)
	}

	// Query with formatted output
	cmdFormatted := NewPackCmd(fs)
	var outFormatted bytes.Buffer
	cmdFormatted.SetOut(&outFormatted)
	if err := cmdFormatted.Flags().Set("focus", "log"); err != nil {
		t.Fatalf("Failed to set focus flag: %v", err)
	}

	err = cmdFormatted.RunE(cmdFormatted, []string{})
	if err != nil {
		t.Errorf("E2E formatted output failed: %v", err)
	}

	rawOutput := outRaw.String()
	formattedOutput := outFormatted.String()

	// Both should have content
	if rawOutput == "" || formattedOutput == "" {
		t.Fatalf("E2E: both raw and formatted output should have content")
	}

	// Raw should not have ANSI codes
	if strings.Contains(rawOutput, "\x1b[") {
		t.Errorf("E2E raw output should not have ANSI codes")
	}

	// Both should mention the search topic
	if !strings.Contains(strings.ToLower(rawOutput), "log") {
		t.Errorf("E2E raw output should mention 'log'")
	}
	if !strings.Contains(strings.ToLower(formattedOutput), "log") {
		t.Errorf("E2E formatted output should mention 'log'")
	}
}

func TestPackE2EErrorHandling(t *testing.T) {
	// End-to-end test: error scenarios
	fs := platform.NewRealFileSystem()

	// Scenario 1: Missing required --focus flag
	cmd1 := NewPackCmd(fs)
	var out1 bytes.Buffer
	cmd1.SetOut(&out1)

	err := cmd1.RunE(cmd1, []string{})
	if err == nil {
		t.Errorf("E2E: missing --focus should error")
	}
	if !strings.Contains(err.Error(), "focus") {
		t.Errorf("E2E: error should mention --focus flag")
	}

	// Scenario 2: Empty topic
	cmd2 := NewPackCmd(fs)
	var out2 bytes.Buffer
	cmd2.SetOut(&out2)
	if err := cmd2.Flags().Set("focus", ""); err != nil {
		t.Fatalf("Failed to set focus flag: %v", err)
	}

	err = cmd2.RunE(cmd2, []string{})
	if err == nil {
		t.Errorf("E2E: empty --focus should error")
	}

	// Scenario 3: Topic with no matches (handled gracefully)
	cmd3 := NewPackCmd(fs)
	var out3 bytes.Buffer
	cmd3.SetOut(&out3)
	if err := cmd3.Flags().Set("focus", "zzz999kkk"); err != nil {
		t.Fatalf("Failed to set focus flag: %v", err)
	}

	err = cmd3.RunE(cmd3, []string{})
	if err != nil {
		t.Errorf("E2E: no matches should not error: %v", err)
	}

	output3 := out3.String()
	if output3 == "" {
		t.Errorf("E2E: no matches should still produce output")
	}
	if !strings.Contains(output3, "No matches") {
		t.Errorf("E2E: no matches output should indicate this")
	}
}

func TestPackE2ESpecsIntegration(t *testing.T) {
	// End-to-end test: verify integration with real specs
	fs := platform.NewRealFileSystem()

	// First, discover specs manually to verify they exist
	specsPath := filepath.Join("..", "..", ".context", "specs")
	specs, err := DiscoverSpecs(fs, specsPath)
	if err != nil {
		t.Skipf("E2E: cannot find specs directory, skipping: %v", err)
	}

	if len(specs) == 0 {
		t.Skipf("E2E: no specs found, skipping")
	}

	// Now run pack command and verify it uses those specs
	cmd := NewPackCmd(fs)
	var out bytes.Buffer
	cmd.SetOut(&out)

	// Use a broad query that should match across specs
	if err := cmd.Flags().Set("focus", "command"); err != nil {
		t.Fatalf("Failed to set focus flag: %v", err)
	}

	err = cmd.RunE(cmd, []string{})
	if err != nil {
		t.Errorf("E2E integration failed: %v", err)
	}

	output := out.String()

	// Should reference multiple specs
	// At least some specs should appear in output
	// (not all necessarily match "command")

	// Output should have meaningful content
	if len(output) < 100 {
		t.Errorf("E2E integration: output seems too short: %d bytes", len(output))
	}
}

func TestPackE2EPerformance(t *testing.T) {
	// End-to-end test: verify performance is reasonable
	fs := platform.NewRealFileSystem()

	cmd := NewPackCmd(fs)
	var out bytes.Buffer
	cmd.SetOut(&out)

	if err := cmd.Flags().Set("focus", "log"); err != nil {
		t.Fatalf("Failed to set focus flag: %v", err)
	}

	// Should complete reasonably quickly (< 500ms)
	done := make(chan error, 1)
	go func() {
		done <- cmd.RunE(cmd, []string{})
	}()

	select {
	case err := <-done:
		if err != nil {
			t.Errorf("E2E performance test failed: %v", err)
		}
	case <-time.After(500 * time.Millisecond):
		t.Errorf("E2E performance test: command took too long (> 500ms)")
	}

	output := out.String()
	if output == "" {
		t.Errorf("E2E performance test: no output produced")
	}
}

func TestPackE2EOutputConsistency(t *testing.T) {
	// End-to-end test: verify output is deterministic
	fs := platform.NewRealFileSystem()

	// Run command twice with same parameters
	var outputs []string

	for run := 0; run < 2; run++ {
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
			t.Errorf("E2E consistency run %d failed: %v", run+1, err)
		}

		outputs = append(outputs, out.String())
	}

	// Outputs should be identical
	if outputs[0] != outputs[1] {
		t.Errorf("E2E consistency: outputs differ between runs")
		t.Logf("Run 1 length: %d", len(outputs[0]))
		t.Logf("Run 2 length: %d", len(outputs[1]))
	}
}

func TestPackE2EInteractiveUsage(t *testing.T) {
	// End-to-end test: simulate interactive exploration
	fs := platform.NewRealFileSystem()

	// Scenario: Developer explores different topics
	topics := []struct {
		query       string
		description string
	}{
		{"log", "learning about logging"},
		{"view", "understanding visualization"},
		{"command", "exploring CLI commands"},
	}

	for _, topic := range topics {
		t.Run(topic.description, func(t *testing.T) {
			cmd := NewPackCmd(fs)
			var out bytes.Buffer
			cmd.SetOut(&out)

			if err := cmd.Flags().Set("focus", topic.query); err != nil {
				t.Fatalf("Failed to set focus flag: %v", err)
			}

			err := cmd.RunE(cmd, []string{})
			if err != nil {
				t.Errorf("Failed to query %q: %v", topic.query, err)
			}

			output := out.String()

			// Verify output structure
			checks := []struct {
				present  bool
				contains string
			}{
				{strings.Contains(output, "#"), "markdown header"},
				{strings.Contains(strings.ToLower(output), "context"), "context pack header"},
				{strings.Contains(output, "---"), "match separators"},
			}

			for _, check := range checks {
				if !check.present {
					t.Errorf("Output should contain %s", check.contains)
				}
			}
		})
	}
}
