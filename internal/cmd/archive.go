package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"cdd/internal/platform"

	"github.com/spf13/cobra"
)

func NewArchiveCmd(fs platform.FileSystem) *cobra.Command {
	return &cobra.Command{
		Use:   "archive [track-name]",
		Short: "Move completed track to history.",
		Long: `Archives the workspace after successful integration of changes.

When a track is archived:
1. All pending tasks in plan.md must be completed (marked with [x]).
2. The spec.md content is appended to .context/inbox.md for downstream processing.
3. Metadata (start/end time) is updated.
4. The track directory is moved from .context/tracks/ to .context/archive/.

EXAMPLES:
  $ cdd archive user-authentication
  $ cdd archive fix-bug-123`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			trackName := args[0]
			src := filepath.Join(".context/tracks", trackName)
			planFile := filepath.Join(src, "plan.md")

			destName := fmt.Sprintf("%s_%s", time.Now().Format("20060102150405"), trackName)
			dest := filepath.Join(".context/archive", destName)

			if _, err := fs.Stat(src); err != nil {
				return fmt.Errorf("Error: Track '%s' not found.", trackName)
			}

			// Validation: Check for unchecked items in plan.md
			if _, err := fs.Stat(planFile); err == nil {
				content, err := fs.ReadFile(planFile)
				if err == nil {
					if strings.Contains(string(content), "[ ]") {
						return fmt.Errorf("Error: Cannot archive '%s'. Pending tasks found in plan.md.", trackName)
					}
				}
			}

			// 1. Time Tracking
			metaFile := filepath.Join(src, "metadata.json")
			if _, err := fs.Stat(metaFile); err == nil {
				content, err := fs.ReadFile(metaFile)
				if err == nil {
					var meta map[string]interface{}
					if err := json.Unmarshal(content, &meta); err == nil {
						meta["archived_at"] = time.Now().Format(time.RFC3339)

						if startStr, ok := meta["started_at"].(string); ok {
							startTime, err := time.Parse(time.RFC3339, startStr)
							if err == nil {
								duration := time.Since(startTime)
								cmd.Printf("⏱️  Duration: %s\n", duration.Round(time.Second))
							}
						}

						metaBytes, err := json.MarshalIndent(meta, "", "  ")
						if err == nil {
							_ = fs.WriteFile(metaFile, metaBytes, 0644)
						}
					}
				}
			}

			// 2. Archive Spec to Inbox
			specFile := filepath.Join(src, "spec.md")
			if _, err := fs.Stat(specFile); err == nil {
				content, err := fs.ReadFile(specFile)
				if err != nil {
					return fmt.Errorf("Error reading spec.md: %v", err)
				}

				inboxFile := filepath.Join(".context", "inbox.md")
				f, err := fs.OpenFile(inboxFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					return fmt.Errorf("Error opening inbox.md: %v", err)
				}

				timestamp := time.Now().Format("2006-01-02 15:04:05")
				entry := fmt.Sprintf("\n---\n###### Archived at: %s | Track: %s\n\n%s\n", timestamp, trackName, string(content))

				if _, err := f.WriteString(entry); err != nil {
					_ = f.Close()
					return fmt.Errorf("Error appending to inbox.md: %v", err)
				}
				if err := f.Close(); err != nil {
					return fmt.Errorf("Error closing inbox.md: %v", err)
				}
			}

			// 3. Cleanup Legacy Files
			filesToDelete := []string{"scratchpad.md", "context_updates.md"}
			for _, f := range filesToDelete {
				_ = fs.Remove(filepath.Join(src, f))
			}

			// 2. Archive
			if err := fs.Rename(src, dest); err != nil {
				return fmt.Errorf("Error archiving track: %v", err)
			}
			cmd.Printf("📦 Archived workspace to '%s'\n", dest)
			return nil
		},
	}
}

func init() {
	rootCmd.AddCommand(NewArchiveCmd(platform.NewRealFileSystem()))
}
