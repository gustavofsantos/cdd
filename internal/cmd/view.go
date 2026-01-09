package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view [track-name]",
	Short: "Render track details.",
	Long: `Render track details.
Usage: 'cdd view' for dashboard, 'cdd view <track>' for details.`,
	Run: func(cmd *cobra.Command, args []string) {
		var contentBuilder strings.Builder

		if len(args) == 0 {
			// Dashboard Mode
			contentBuilder.WriteString("# 📂 Project Dashboard\n\n")
			contentBuilder.WriteString("## 🌍 Global Context\n")
			if content, err := os.ReadFile(".context/product.md"); err == nil {
				contentBuilder.Write(content)
			} else {
				contentBuilder.WriteString("_No product.md found._\n")
			}

			contentBuilder.WriteString("## 📥 Context Inbox (Pending Updates)\n")
			if content, err := os.ReadFile(".context/inbox.md"); err == nil {
				// tail -n 5 equivalent
				lines := splitLines(content)
				if len(lines) > 5 {
					lines = lines[len(lines)-5:]
				}
				for _, line := range lines {
					contentBuilder.WriteString(line + "\n")
				}
			} else {
				contentBuilder.WriteString("_Inbox empty._\n")
			}
			contentBuilder.WriteString("\n## 🟢 Active Tracks\n")

			entries, err := os.ReadDir(".context/tracks")
			if err == nil {
				found := false
				for _, entry := range entries {
					if entry.IsDir() {
						contentBuilder.WriteString(fmt.Sprintf("* **%s**\n", entry.Name()))
						found = true
					}
				}
				if !found {
					contentBuilder.WriteString("_No active tracks._\n")
				}
			} else {
				contentBuilder.WriteString("_No active tracks._\n")
			}
		} else {
			// Track Detail Mode
			trackName := args[0]
			trackDir := filepath.Join(".context/tracks", trackName)
			if _, err := os.Stat(trackDir); os.IsNotExist(err) {
				fmt.Printf("Error: Track '%s' not found.\n", trackName)
				os.Exit(1)
			}

			contentBuilder.WriteString(fmt.Sprintf("# 🛤️ Track: %s\n\n", trackName))

			contentBuilder.WriteString("## 📄 Specification\n")
			if content, err := os.ReadFile(filepath.Join(trackDir, "spec.md")); err == nil {
				contentBuilder.Write(content)
			} else {
				contentBuilder.WriteString("_Missing spec.md_\n")
			}

			contentBuilder.WriteString("\n## 📋 Plan\n")
			if content, err := os.ReadFile(filepath.Join(trackDir, "plan.md")); err == nil {
				contentBuilder.Write(content)
			} else {
				contentBuilder.WriteString("_Missing plan.md_\n")
			}

			contentBuilder.WriteString("\n## 📜 Recent Decisions\n")
			if content, err := os.ReadFile(filepath.Join(trackDir, "decisions.md")); err == nil {
				lines := splitLines(content)
				if len(lines) > 5 {
					lines = lines[len(lines)-5:]
				}
				for _, line := range lines {
					contentBuilder.WriteString(line + "\n")
				}
			}
		}

		out, err := glamour.Render(contentBuilder.String(), "dark")
		if err != nil {
			fmt.Printf("Error rendering markdown: %v\n", err)
			os.Exit(1)
		}
		fmt.Print(out)
	},
}

func splitLines(content []byte) []string {
	// Simple split lines, handling carriage returns
	s := string(content)
	var lines []string
	for _, line := range splitLinesString(s) {
		lines = append(lines, line)
	}
	return lines
}

func splitLinesString(s string) []string {
	// A robust implementation would use bufio.Scanner
	// For simplicity, just split by newline
	var lines []string
	start := 0
	for i, c := range s {
		if c == '\n' {
			lines = append(lines, s[start:i])
			start = i + 1
		}
	}
	if start < len(s) {
		lines = append(lines, s[start:])
	}
	return lines
}

func init() {
	rootCmd.AddCommand(viewCmd)
}
