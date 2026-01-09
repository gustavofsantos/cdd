package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var dumpCmd = &cobra.Command{
	Use:   "dump [track-name]",
	Short: "Pipe large output to scratchpad.",
	Long: `Pipe large output to scratchpad to keep chat context clean.
Usage: command | cdd dump <track-name>`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		trackName := args[0]
		scratchFile := filepath.Join(".context/tracks", trackName, "scratchpad.md")

		f, err := os.OpenFile(scratchFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Error opening scratchpad: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()

		if _, err := io.Copy(f, os.Stdin); err != nil {
			fmt.Printf("Error writing to scratchpad: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf(">> Data appended to %s\n", scratchFile)
		fmt.Println(">> INSTRUCTION: Do not read the whole file. Use grep/head/tail to find what you need.")
	},
}

func init() {
	rootCmd.AddCommand(dumpCmd)
}
