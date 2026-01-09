package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var archiveCmd = &cobra.Command{
	Use:   "archive [track-name]",
	Short: "Move completed track to history and promote knowledge.",
	Long: `Promotes knowledge to Living Docs and archives the workspace.
Constraint: Cannot archive if there are pending tasks ([ ]).`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		trackName := args[0]
		src := filepath.Join(".context/tracks", trackName)
		planFile := filepath.Join(src, "plan.md")
		specFile := filepath.Join(src, "spec.md")
		updateFile := filepath.Join(src, "context_updates.md")

		destName := fmt.Sprintf("%s_%s", time.Now().Format("20060102"), trackName)
		dest := filepath.Join(".context/archive", destName)
		featureDest := filepath.Join(".context/features", trackName+".md")
		inboxFile := ".context/inbox.md"

		if _, err := os.Stat(src); os.IsNotExist(err) {
			fmt.Printf("Error: Track '%s' not found.\n", trackName)
			os.Exit(1)
		}

		// Validation: Check for unchecked items
		if _, err := os.Stat(planFile); err == nil {
			content, err := os.ReadFile(planFile)
			if err == nil {
				if strings.Contains(string(content), "[ ]") {
					fmt.Printf("Error: Cannot archive '%s'. Pending tasks found in plan.md.\n", trackName)
					os.Exit(1)
				}
			}
		}

		// 1. Promote Spec to Living Documentation
		if _, err := os.Stat(specFile); err == nil {
			input, err := os.ReadFile(specFile)
			if err == nil {
				os.WriteFile(featureDest, input, 0644)
				fmt.Printf("✅ Promoted spec to Living Docs: %s\n", featureDest)
			}
		}

		// 2. Handle Context Updates (Inbox Pattern)
		if _, err := os.Stat(updateFile); err == nil {
			info, _ := os.Stat(updateFile)
			if info.Size() > 0 {
				content, err := os.ReadFile(updateFile)
				if err == nil {
					lines := strings.Split(strings.TrimSpace(string(content)), "\n")
					// Check if file has content other than header (heuristic: > 2 lines)
					if len(lines) > 2 {
						f, err := os.OpenFile(inboxFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err == nil {
							defer f.Close()
							header := fmt.Sprintf("\n\n## Updates from Track: %s (%s)\n", trackName, time.Now().Format("Mon Jan 2 15:04:05 MST 2006"))
							f.WriteString(header)
							f.Write(content)
							fmt.Printf("✅ Appended context updates to %s\n", inboxFile)
						}
					}
				}
			}
		}

		// 3. Archive
		if err := os.Rename(src, dest); err != nil {
			fmt.Printf("Error archiving track: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("📦 Archived workspace to '%s'\n", dest)
	},
}

func init() {
	rootCmd.AddCommand(archiveCmd)
}
