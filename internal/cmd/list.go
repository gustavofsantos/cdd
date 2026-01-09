package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List active tracks (default) or archived tracks.",
	Long:  `List active tracks (default) or archived tracks.`,
	Run: func(cmd *cobra.Command, args []string) {
		archived, _ := cmd.Flags().GetBool("archived")

		if archived {
			fmt.Println("ARCHIVED TRACKS:")
			printDirContents(".context/archive")
		} else {
			fmt.Println("ACTIVE TRACKS:")
			printDirContents(".context/tracks")
		}
	},
}

func printDirContents(path string) {
	entries, err := os.ReadDir(path)
	if err != nil || len(entries) == 0 {
		fmt.Println("  (None)")
		return
	}

	found := false
	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Printf("  - %s\n", entry.Name())
			found = true
		}
	}
	if !found {
		fmt.Println("  (None)")
	}
}

func init() {
	listCmd.Flags().Bool("archived", false, "List archived tracks")
	rootCmd.AddCommand(listCmd)
}
