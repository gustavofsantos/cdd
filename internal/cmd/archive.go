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
		Short: "Move completed track to history and promote knowledge.",
		Long: `Promotes knowledge to Living Docs and archives the workspace.
Constraint: Cannot archive if there are pending tasks ([ ]).`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			trackName := args[0]
			src := filepath.Join(".context/tracks", trackName)
			planFile := filepath.Join(src, "plan.md")
			specFile := filepath.Join(src, "spec.md")
			updateFile := filepath.Join(src, "context_updates.md")
			scratchFile := filepath.Join(src, "scratchpad.md")

			destName := fmt.Sprintf("%s_%s", time.Now().Format("20060102150405"), trackName)
			dest := filepath.Join(".context/archive", destName)
			featureDest := filepath.Join(".context/features", trackName+".md")
			inboxFile := ".context/inbox.md"

			// Simplified check:
			if _, err := fs.Stat(src); err != nil {
				return fmt.Errorf("Error: Track '%s' not found.", trackName)
			}

			// Validation: Check for unchecked items
			if _, err := fs.Stat(planFile); err == nil {
				content, err := fs.ReadFile(planFile)
				if err == nil {
					if strings.Contains(string(content), "[ ]") {
						return fmt.Errorf("Error: Cannot archive '%s'. Pending tasks found in plan.md.", trackName)
					}
				}
			}

			// 0. Cleanup ephemeral files
			_ = fs.Remove(scratchFile)

			// 1. Promote Spec to Living Documentation
			if _, err := fs.Stat(specFile); err == nil {
				input, err := fs.ReadFile(specFile)
				if err == nil {
					if err := fs.WriteFile(featureDest, input, 0644); err != nil {
						return fmt.Errorf("Error promoting spec: %v", err)
					}
					cmd.Printf("✅ Promoted spec to Living Docs: %s\n", featureDest)
				}
			}

			// 2. Handle Context Updates (Inbox Pattern)
			if _, err := fs.Stat(updateFile); err == nil {
				info, _ := fs.Stat(updateFile)
				if info.Size() > 0 {
					content, err := fs.ReadFile(updateFile)
					if err == nil {
						lines := strings.Split(strings.TrimSpace(string(content)), "\n")
						// Check if file has content other than header (heuristic: > 2 lines)
						if len(lines) > 2 {
							f, err := fs.OpenFile(inboxFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
							if err == nil {
								header := fmt.Sprintf("\n\n## Updates from Track: %s (%s)\n", trackName, time.Now().Format("Mon Jan 2 15:04:05 MST 2006"))
								if _, err := f.WriteString(header); err != nil {
									_ = f.Close()
									return fmt.Errorf("Error writing header to inbox: %v", err)
								}
								if _, err := f.WriteString(string(content)); err != nil {
									_ = f.Close()
									return fmt.Errorf("Error writing content to inbox: %v", err)
								}
								if err := f.Close(); err != nil {
									return fmt.Errorf("Error closing inbox file: %v", err)
								}
								cmd.Printf("✅ Appended context updates to %s\n", inboxFile)

								// Trigger cleanup reminder
								CheckInboxSize(fs, cmd)
							}
						}
					}
				}
			}

			// 3. Time Tracking (Calculate Duration)
			metaFile := filepath.Join(src, "metadata.json")
			if _, err := fs.Stat(metaFile); err == nil {
				content, err := fs.ReadFile(metaFile)
				if err == nil {
					var meta map[string]interface{}
					if err := json.Unmarshal(content, &meta); err == nil {
						if startStr, ok := meta["started_at"].(string); ok {
							startTime, err := time.Parse(time.RFC3339, startStr)
							if err == nil {
								duration := time.Since(startTime)
								cmd.Printf("⏱️  Duration: %s\n", duration.Round(time.Second))
							}
						}
					}
				}
			}

			// 4. Archive
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
