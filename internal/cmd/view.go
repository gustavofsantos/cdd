package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"cdd/internal/platform"

	"github.com/charmbracelet/glamour"
	"github.com/mattn/go-isatty"
	"github.com/spf13/cobra"
)

var (
	viewInbox    bool
	viewArchived bool
	viewSpec     bool
	viewPlan     bool
	viewRaw      bool
)

func NewViewCmd(fs platform.FileSystem) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "view [track-name]",
		Short: "Render track details.",
		Long: `Render track details.
Usage: 'cdd view' for dashboard, 'cdd view <track>' for details.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			markdown, err := buildViewMarkdown(fs, args)
			if err != nil {
				return err
			}

			// If not a TTY or --raw is set, print raw markdown/text
			if viewRaw || (!isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd())) {
				_, err := fmt.Fprint(cmd.OutOrStdout(), markdown)
				return err
			}

			out, err := glamour.Render(markdown, "dark")
			if err != nil {
				return fmt.Errorf("error rendering markdown: %v", err)
			}
			_, err = fmt.Fprint(cmd.OutOrStdout(), out)
			return err
		},
	}

	cmd.Flags().BoolVarP(&viewInbox, "inbox", "i", false, "Show context inbox")
	cmd.Flags().BoolVarP(&viewArchived, "archived", "a", false, "Show archived tracks")
	cmd.Flags().BoolVarP(&viewSpec, "spec", "s", false, "Show track specification")
	cmd.Flags().BoolVarP(&viewPlan, "plan", "p", false, "Show track plan")
	cmd.Flags().BoolVarP(&viewRaw, "raw", "r", false, "Output raw text (suitable for piping)")

	return cmd
}

func buildViewMarkdown(fs platform.FileSystem, args []string) (string, error) {
	var contentBuilder strings.Builder
	isRaw := viewRaw || (!isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd()))

	if len(args) == 0 {
		// Dashboard Mode
		if viewInbox {
			if !isRaw {
				contentBuilder.WriteString("# 📥 Context Inbox (Pending Updates)\n\n")
			}
			if content, err := fs.ReadFile(".context/inbox.md"); err == nil {
				contentBuilder.Write(content)
			} else if !isRaw {
				contentBuilder.WriteString("_Inbox empty._\n")
			}
		} else if viewArchived {
			if !isRaw {
				contentBuilder.WriteString("# 📦 Archived Tracks\n\n")
			}
			entries, err := fs.ReadDir(".context/archive")
			if err == nil {
				found := false
				for _, entry := range entries {
					if entry.IsDir() {
						cleanName := stripTimestamp(entry.Name())
						if isRaw {
							contentBuilder.WriteString(cleanName + "\n")
						} else {
							contentBuilder.WriteString(fmt.Sprintf("* **%s**\n", cleanName))
						}
						found = true
					}
				}
				if !found && !isRaw {
					contentBuilder.WriteString("_No archived tracks._\n")
				}
			} else if !isRaw {
				contentBuilder.WriteString("_No archive directory found._\n")
			}
		} else {
			if !isRaw {
				contentBuilder.WriteString("# 🟢 Active Tracks\n\n")
			}
			entries, err := fs.ReadDir(".context/tracks")
			if err == nil {
				found := false
				for _, entry := range entries {
					if entry.IsDir() {
						if isRaw {
							contentBuilder.WriteString(entry.Name() + "\n")
						} else {
							contentBuilder.WriteString(fmt.Sprintf("* **%s**\n", entry.Name()))
						}
						found = true
					}
				}
				if !found && !isRaw {
					contentBuilder.WriteString("_No active tracks._\n")
				}
			} else if !isRaw {
				contentBuilder.WriteString("_No tracks directory found._\n")
			}
		}
	} else {
		// Track Detail Mode
		trackName := args[0]
		var trackDir string

		if viewArchived {
			// Search for archived track (it has a timestamp prefix)
			archiveEntries, err := fs.ReadDir(".context/archive")
			if err == nil {
				for _, entry := range archiveEntries {
					if entry.IsDir() && strings.HasSuffix(entry.Name(), "_"+trackName) {
						trackDir = filepath.Join(".context/archive", entry.Name())
						trackName = entry.Name() // Use full name for display
						break
					}
				}
			}
		} else {
			trackDir = filepath.Join(".context/tracks", trackName)
		}

		if trackDir == "" || !dirExistsFS(fs, trackDir) {
			return "", fmt.Errorf("track '%s' not found", trackName)
		}

		contentBuilder.WriteString(fmt.Sprintf("# 🛤️ Track: %s\n\n", trackName))

		if viewSpec {
			contentBuilder.WriteString("## 📄 Specification (spec.md)\n")
			if content, err := fs.ReadFile(filepath.Join(trackDir, "spec.md")); err == nil {
				contentBuilder.Write(content)
			} else {
				contentBuilder.WriteString("_Missing spec.md_\n")
			}
		} else {
			// Default to plan/Next Tasks unless viewPlan is explicitly set
			contentBuilder.WriteString("## 📋 Next Tasks\n")
			if content, err := fs.ReadFile(filepath.Join(trackDir, "plan.md")); err == nil {
				tasks := extractNextTasks(string(content))
				if len(tasks) > 0 {
					contentBuilder.WriteString(tasks)
				} else {
					contentBuilder.WriteString("_No pending tasks._\n")
				}
			} else {
				contentBuilder.WriteString("_Missing plan.md_\n")
			}
		}
	}

	return contentBuilder.String(), nil
}

func extractNextTasks(content string) string {
	lines := strings.Split(content, "\n")
	var tasks []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.Contains(trimmed, "[ ]") {
			tasks = append(tasks, line)
		}
	}
	if len(tasks) == 0 {
		return ""
	}
	return strings.Join(tasks, "\n") + "\n"
}

func dirExistsFS(fs platform.FileSystem, path string) bool {
	info, err := fs.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func stripTimestamp(name string) string {
	parts := strings.SplitN(name, "_", 2)
	if len(parts) > 1 && len(parts[0]) == 14 {
		// Basic check if it's a timestamp (all digits)
		isTimestamp := true
		for _, c := range parts[0] {
			if c < '0' || c > '9' {
				isTimestamp = false
				break
			}
		}
		if isTimestamp {
			return parts[1]
		}
	}
	return name
}

func init() {
	rootCmd.AddCommand(NewViewCmd(platform.NewRealFileSystem()))
}
