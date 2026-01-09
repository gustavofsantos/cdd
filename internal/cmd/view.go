package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view [track-name]",
	Short: "Render track details using 'glow'.",
	Long: `Render track details using 'glow' (User Tool Only).
Usage: 'cdd view' for dashboard, 'cdd view <track>' for details.`,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := exec.LookPath("glow"); err != nil {
			fmt.Println("Error: 'glow' is not installed. Please install it: [https://github.com/charmbracelet/glow](https://github.com/charmbracelet/glow)")
			os.Exit(1)
		}

		tmpFile := ".cdd_view_tmp.md"
		defer os.Remove(tmpFile)

		if len(args) == 0 {
			// Dashboard Mode
			f, err := os.Create(tmpFile)
			if err != nil {
				fmt.Printf("Error creating temp file: %v\n", err)
				os.Exit(1)
			}
			defer f.Close()

			f.WriteString("# 📂 Project Dashboard\n\n")
			f.WriteString("## 🌍 Global Context\n")
			if content, err := os.ReadFile(".context/product.md"); err == nil {
				f.Write(content)
			} else {
				f.WriteString("_No product.md found._\n")
			}

			f.WriteString("## 📥 Context Inbox (Pending Updates)\n")
			if content, err := os.ReadFile(".context/inbox.md"); err == nil {
				// tail -n 5 equivalent
				lines := splitLines(content)
				if len(lines) > 5 {
					lines = lines[len(lines)-5:]
				}
				for _, line := range lines {
					f.WriteString(line + "\n")
				}
			} else {
				f.WriteString("_Inbox empty._\n")
			}
			f.WriteString("\n## 🟢 Active Tracks\n")

			entries, err := os.ReadDir(".context/tracks")
			if err == nil {
				found := false
				for _, entry := range entries {
					if entry.IsDir() {
						f.WriteString(fmt.Sprintf("* **%s**\n", entry.Name()))
						found = true
					}
				}
				if !found {
					f.WriteString("_No active tracks._\n")
				}
			} else {
				f.WriteString("_No active tracks._\n")
			}
		} else {
			// Track Detail Mode
			trackName := args[0]
			trackDir := filepath.Join(".context/tracks", trackName)
			if _, err := os.Stat(trackDir); os.IsNotExist(err) {
				fmt.Printf("Error: Track '%s' not found.\n", trackName)
				os.Exit(1)
			}

			f, err := os.Create(tmpFile)
			if err != nil {
				fmt.Printf("Error creating temp file: %v\n", err)
				os.Exit(1)
			}
			defer f.Close()

			f.WriteString(fmt.Sprintf("# 🛤️ Track: %s\n\n", trackName))

			f.WriteString("## 📄 Specification\n")
			if content, err := os.ReadFile(filepath.Join(trackDir, "spec.md")); err == nil {
				f.Write(content)
			} else {
				f.WriteString("_Missing spec.md_\n")
			}

			f.WriteString("\n## 📋 Plan\n")
			if content, err := os.ReadFile(filepath.Join(trackDir, "plan.md")); err == nil {
				f.Write(content)
			} else {
				f.WriteString("_Missing plan.md_\n")
			}

			f.WriteString("\n## 📜 Recent Decisions\n")
			if content, err := os.ReadFile(filepath.Join(trackDir, "decisions.md")); err == nil {
				lines := splitLines(content)
				if len(lines) > 5 {
					lines = lines[len(lines)-5:]
				}
				for _, line := range lines {
					f.WriteString(line + "\n")
				}
			}
		}

		c := exec.Command("glow", tmpFile)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()
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
