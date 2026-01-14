package cmd

import (
	"fmt"
	"io"
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
		Long: `Record a permanent decision, error, or architectural choice.

The log command appends a timestamped entry to the track's decisions.md file. 
This file serves as a permanent record of the "Why" behind changes.

You can provide the message as an argument, or pipe it via STDIN. 
If only one track is active, the track name is optional when using STDIN.

EXAMPLES:
  $ cdd log user-auth "Switched to JWT for session management"
  $ cdd log "Fixed intermittent CI failure"
  $ cdd log << 'EOF'
    ADR: Use Redis for caching.
    Reasoning: Performance requirements for the new dashboard.
    EOF`,
		Args: cobra.RangeArgs(0, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			var trackName string
			var msg string

			// Check for stdin
			var hasStdin bool
			in := cmd.InOrStdin()

			// If it is actual os.Stdin, check for pipe/redirection to avoid hanging
			if in == os.Stdin {
				stat, _ := os.Stdin.Stat()
				hasStdin = (stat.Mode() & os.ModeCharDevice) == 0
			} else {
				// For tests where we replaced input
				hasStdin = true
			}

			if hasStdin {
				input, err := io.ReadAll(in)
				if err != nil {
					return fmt.Errorf("error reading from stdin: %v", err)
				}
				msg = string(input)
			}

			if len(args) == 2 {
				trackName = args[0]
				msg = args[1]
			} else if len(args) == 1 {
				trackName = args[0]
				if !hasStdin {
					return fmt.Errorf("message is required")
				}
			} else if len(args) == 0 {
				if !hasStdin {
					return fmt.Errorf("track name and message are required")
				}
				// Infer track name
				tracks, err := fs.ReadDir(".context/tracks")
				if err != nil {
					return fmt.Errorf("error reading tracks directory: %v", err)
				}
				var activeTracks []string
				for _, t := range tracks {
					if t.IsDir() {
						activeTracks = append(activeTracks, t.Name())
					}
				}

				if len(activeTracks) == 1 {
					trackName = activeTracks[0]
				} else if len(activeTracks) == 0 {
					return fmt.Errorf("no active tracks found")
				} else {
					return fmt.Errorf("multiple active tracks found (%d), please specify one", len(activeTracks))
				}
			}

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
