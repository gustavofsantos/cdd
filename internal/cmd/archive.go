package cmd

import (
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

			destName := fmt.Sprintf("%s_%s", time.Now().Format("20060102"), trackName)
			dest := filepath.Join(".context/archive", destName)
			featureDest := filepath.Join(".context/features", trackName+".md")
			inboxFile := ".context/inbox.md"

			if _, err := fs.Stat(src); !os.IsNotExist(err) && err != nil {
				// Check logic: fs.Stat succeeds -> err=nil.
				// fs.Stat fails -> Check if NotExist.
				// If NotExist, error.
			}
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

			// 1. Promote Spec to Living Documentation
			if _, err := fs.Stat(specFile); err == nil {
				input, err := fs.ReadFile(specFile)
				if err == nil {
					fs.WriteFile(featureDest, input, 0644)
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
								f.WriteString(header)
								f.WriteString(string(content))
								f.Close()
								cmd.Printf("✅ Appended context updates to %s\n", inboxFile)

								// Trigger cleanup reminder
								CheckInboxSize(fs, cmd)
							}
						}
					}
				}
			}

			// 3. Archive
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
