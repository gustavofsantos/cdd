package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log [track-name] [message]",
	Short: "Record a permanent decision or error.",
	Long: `Record a permanent decision or error.
'Mask, Don't Remove'. We must keep evidence of failures to avoid repeating them.
Usage: cdd log <track-name> <message>`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		trackName := args[0]
		msg := args[1]
		logFile := filepath.Join(".context/tracks", trackName, "decisions.md")

		f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Error opening log file: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()

		timestamp := time.Now().Format("2006-01-02 15:04:05")
		entry := fmt.Sprintf("[%s] %s\n", timestamp, msg)
		if _, err := f.WriteString(entry); err != nil {
			fmt.Printf("Error writing to log: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Logged to %s\n", logFile)
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
}
