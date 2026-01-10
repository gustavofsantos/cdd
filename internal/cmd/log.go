package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"cdd/internal/platform"

	"github.com/spf13/cobra"
)

func NewLogCmd(fs platform.FileSystem) *cobra.Command {
	return &cobra.Command{
		Use:   "log [track-name] [message]",
		Short: "Record a permanent decision or error.",
		Long: `Record a permanent decision or error.
'Mask, Don't Remove'. We must keep evidence of failures to avoid repeating them.
Usage: cdd log <track-name> <message>`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			trackName := args[0]
			msg := args[1]
			logFile := filepath.Join(".context/tracks", trackName, "decisions.md")

			f, err := fs.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return fmt.Errorf("Error opening log file: %v", err)
			}
			defer func() { _ = f.Close() }()

			timestamp := time.Now().Format("2006-01-02 15:04:05")
			entry := fmt.Sprintf("[%s] %s\n", timestamp, msg)
			if _, err := f.WriteString(entry); err != nil {
				return fmt.Errorf("Error writing to log: %v", err)
			}

			cmd.Printf("Logged to %s\n", logFile)
			return nil
		},
	}
}

func init() {
	rootCmd.AddCommand(NewLogCmd(platform.NewRealFileSystem()))
}
