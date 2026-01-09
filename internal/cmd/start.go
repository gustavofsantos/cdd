package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"cdd/internal/platform"

	"github.com/spf13/cobra"
)

func NewStartCmd(fs platform.FileSystem) *cobra.Command {
	return &cobra.Command{
		Use:   "start [track-name]",
		Short: "Create an isolated workspace (Track).",
		Long: `Creates an isolated workspace so multiple agents/tasks don't collide.
Usage: cdd start <track-name>`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			trackName := args[0]
			trackDir := filepath.Join(".context/tracks", trackName)

			if _, err := fs.Stat(trackDir); !os.IsNotExist(err) {
				return fmt.Errorf("Error: Track '%s' exists.", trackName)
			}

			if err := fs.MkdirAll(trackDir, 0755); err != nil {
				return fmt.Errorf("Error creating track directory: %v", err)
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
			fs.WriteFile(filepath.Join(trackDir, "spec.md"), []byte(specTemplate), 0644)

			// Context Updates Staging File
			updatesContent := "# Proposed Global Context Updates\n> Add notes here if product.md or tech-stack.md needs updating.\n"
			fs.WriteFile(filepath.Join(trackDir, "context_updates.md"), []byte(updatesContent), 0644)

			// Plan Template
			planContent := fmt.Sprintf("# Plan for %s\n- [ ] 🗣️ Phase 0: Alignment & Analysis (Fill spec.md)\n- [ ] 📝 Phase 1: Approval (User signs off)\n", trackName)
			fs.WriteFile(filepath.Join(trackDir, "plan.md"), []byte(planContent), 0644)

			// Decisions Log
			decisionsContent := fmt.Sprintf("# Decision Log\n> Created %s\n", time.Now().Format("Mon Jan 2 15:04:05 MST 2006"))
			fs.WriteFile(filepath.Join(trackDir, "decisions.md"), []byte(decisionsContent), 0644)

			// Scratchpad
			scratchContent := fmt.Sprintf("# Scratchpad for %s\n> Dump raw logs here.\n", trackName)
			fs.WriteFile(filepath.Join(trackDir, "scratchpad.md"), []byte(scratchContent), 0644)

			cmd.Printf("Track '%s' initialized.\n", trackName)
			return nil
		},
	}
}

func init() {
	rootCmd.AddCommand(NewStartCmd(platform.NewRealFileSystem()))
}
