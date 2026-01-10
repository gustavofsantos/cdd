package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"cdd/internal/platform"

	"github.com/spf13/cobra"
)

func NewDeleteCmd(fs platform.FileSystem) *cobra.Command {
	return &cobra.Command{
		Use:   "delete [track]",
		Short: "Delete a non-archived track.",
		Long:  `Delete a non-archived track. This command is intended for user workflow facilitation and should not be used by AI.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			track := args[0]

			// Sanitize input to prevent directory traversal
			if strings.Contains(track, string(os.PathSeparator)) || strings.Contains(track, "..") {
				return fmt.Errorf("invalid track name: %s", track)
			}

			path := filepath.Join(".context", "tracks", track)

			// Check if track exists in active tracks
			_, err := fs.Stat(path)
			if os.IsNotExist(err) {
				return fmt.Errorf("track '%s' not found in active tracks", track)
			} else if err != nil {
				return err
			}

			// Delete the track
			if err := fs.RemoveAll(path); err != nil {
				return fmt.Errorf("failed to delete track '%s': %w", track, err)
			}

			cmd.Printf("Track '%s' deleted.\n", track)
			return nil
		},
	}
}

func init() {
	rootCmd.AddCommand(NewDeleteCmd(platform.NewRealFileSystem()))
}
