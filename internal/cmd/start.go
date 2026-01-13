package cmd

import (
	"encoding/json"
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
		Long: `Creates an isolated workspace following the Lean CDD v4.1 protocol.
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

			// Spec Template (v4.1 Delta Spec)
			specTemplate := fmt.Sprintf(`# Track: %s
**Target Spec:** .context/specs/TODO/spec.md

## Context
(Links to relevant files)

## Proposed Changes
### ADDED Requirements
* **Requirement: Feature Name**
    * The system SHALL ...
    * #### Scenario: Happy Path
        * Given ...
        * When ...
        * Then ...

### MODIFIED Requirements
* **Requirement: Existing Feature**
    * (Copy current text and show changes)
`, trackName)
			if err := fs.WriteFile(filepath.Join(trackDir, "spec.md"), []byte(specTemplate), 0644); err != nil {
				return fmt.Errorf("failed to write spec.md: %w", err)
			}

			// Plan Template (v4.1 TDD Steps)
			planContent := fmt.Sprintf("# Plan for %s\n[ ] 🔴 Test: (Red)\n[ ] 🟢 Impl: (Green)\n[ ] 🔵 Refactor: (Refactor)\n", trackName)
			if err := fs.WriteFile(filepath.Join(trackDir, "plan.md"), []byte(planContent), 0644); err != nil {
				return fmt.Errorf("failed to write plan.md: %w", err)
			}

			// Decisions Log
			decisionsContent := fmt.Sprintf("# Decision Log\n> Created %s\n", time.Now().Format("Mon Jan 2 15:04:05 MST 2006"))
			if err := fs.WriteFile(filepath.Join(trackDir, "decisions.md"), []byte(decisionsContent), 0644); err != nil {
				return fmt.Errorf("failed to write decisions.md: %w", err)
			}

			// Metadata for Time Tracking (Internal)
			metadata := map[string]interface{}{
				"started_at": time.Now().Format(time.RFC3339),
			}
			metaBytes, err := json.MarshalIndent(metadata, "", "  ")
			if err != nil {
				return fmt.Errorf("failed to marshal metadata: %w", err)
			}
			if err := fs.WriteFile(filepath.Join(trackDir, "metadata.json"), metaBytes, 0644); err != nil {
				return fmt.Errorf("failed to write metadata.json: %w", err)
			}

			cmd.Printf("Track '%s' initialized.\n活跃 (Active) Track created with 3-file structure.\n", trackName)
			return nil
		},
	}
}

func init() {
	rootCmd.AddCommand(NewStartCmd(platform.NewRealFileSystem()))
}
