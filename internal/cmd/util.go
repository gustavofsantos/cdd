package cmd

import (
	"strings"

	"cdd/internal/platform"

	"github.com/spf13/cobra"
)

// CheckInboxSize checks the line count of .context/inbox.md and prints a suggestion if it exceeds 50 lines.
func CheckInboxSize(fs platform.FileSystem, cmd *cobra.Command) {
	inboxFile := ".context/inbox.md"
	content, err := fs.ReadFile(inboxFile)
	if err != nil {
		return // Inbox might not exist yet, or error reading
	}

	lines := strings.Split(string(content), "\n")
	lineCount := len(lines)
	// Some Split implementations might return extra line if ends with \n,
	// but for a reminder this is fine.
	if lineCount > 50 {
		cmd.Printf("\n⚠️ Your .context/inbox.md has %d lines. It's getting large! Run 'cdd init --inbox-prompt' to get the cleaner prompt and run it with your agent to consolidate context.\n", lineCount)
	}
}
