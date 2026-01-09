package cmd

import (
	"cdd/internal/platform"

	"github.com/spf13/cobra"
)

func NewListCmd(fs platform.FileSystem) *cobra.Command {
	c := &cobra.Command{
		Use:   "list",
		Short: "List active tracks (default) or archived tracks.",
		Long:  `List active tracks (default) or archived tracks.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			archived, _ := cmd.Flags().GetBool("archived")

			if archived {
				cmd.Println("ARCHIVED TRACKS:")
				printDirContents(fs, cmd, ".context/archive")
			} else {
				cmd.Println("ACTIVE TRACKS:")
				printDirContents(fs, cmd, ".context/tracks")
			}
			return nil
		},
	}
	c.Flags().Bool("archived", false, "List archived tracks")
	return c
}

func printDirContents(fs platform.FileSystem, cmd *cobra.Command, path string) {
	entries, err := fs.ReadDir(path)
	if err != nil || len(entries) == 0 {
		cmd.Println("  (None)")
		return
	}

	found := false
	for _, entry := range entries {
		if entry.IsDir() {
			cmd.Printf("  - %s\n", entry.Name())
			found = true
		}
	}
	if !found {
		cmd.Println("  (None)")
	}
}

func init() {
	rootCmd.AddCommand(NewListCmd(platform.NewRealFileSystem()))
}
