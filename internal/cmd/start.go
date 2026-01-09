package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start [track-name]",
	Short: "Create an isolated workspace (Track).",
	Long: `Creates an isolated workspace so multiple agents/tasks don't collide.
Usage: cdd start <track-name>`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		trackName := args[0]
		trackDir := filepath.Join(".context/tracks", trackName)

		if _, err := os.Stat(trackDir); !os.IsNotExist(err) {
			fmt.Printf("Error: Track '%s' exists.\n", trackName)
			os.Exit(1)
		}

		if err := os.MkdirAll(trackDir, 0755); err != nil {
			fmt.Printf("Error creating track directory: %v\n", err)
			os.Exit(1)
		}

		// Spec Template
		specTemplate := fmt.Sprintf(`# Specification: %s

## 1. User Intent (The Goal)
> [User Input Required]

## 2. Relevant Context (The Files)
> [Agent to Populate during Analysis]
- `+"`path/to/relevant/file.ext`"+`

## 3. Context Analysis (Agent Findings)
> [Agent to Populate]
> - Current Behavior:
> - Proposed Changes:

## 4. Scenarios (Acceptance Criteria)
> [Agent to Draft based on Intent - Gherkin Style Preferred]
> Feature: %s
>   Scenario: Happy Path
>     Given ...
>     When ...
>     Then ...
`, trackName, trackName)
		os.WriteFile(filepath.Join(trackDir, "spec.md"), []byte(specTemplate), 0644)

		// Context Updates Staging File
		updatesContent := "# Proposed Global Context Updates\n> Add notes here if product.md or tech-stack.md needs updating.\n"
		os.WriteFile(filepath.Join(trackDir, "context_updates.md"), []byte(updatesContent), 0644)

		// Plan Template
		planContent := fmt.Sprintf("# Plan for %s\n- [ ] 🗣️ Phase 0: Alignment & Analysis (Fill spec.md)\n- [ ] 📝 Phase 1: Approval (User signs off)\n", trackName)
		os.WriteFile(filepath.Join(trackDir, "plan.md"), []byte(planContent), 0644)

		// Decisions Log
		decisionsContent := fmt.Sprintf("# Decision Log\n> Created %s\n", time.Now().Format("Mon Jan 2 15:04:05 MST 2006"))
		os.WriteFile(filepath.Join(trackDir, "decisions.md"), []byte(decisionsContent), 0644)

		// Scratchpad
		scratchContent := fmt.Sprintf("# Scratchpad for %s\n> Dump raw logs here.\n", trackName)
		os.WriteFile(filepath.Join(trackDir, "scratchpad.md"), []byte(scratchContent), 0644)

		fmt.Printf("Track '%s' initialized.\n", trackName)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
