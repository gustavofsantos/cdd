package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"cdd/internal/platform"

	"github.com/spf13/cobra"
)

func NewDumpCmd(fs platform.FileSystem) *cobra.Command {
	return &cobra.Command{
		Use:   "dump [track-name]",
		Short: "Pipe large output to scratchpad.",
		Long: `Pipe large output to scratchpad to keep chat context clean.
Usage: command | cdd dump <track-name>`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			trackName := args[0]
			scratchFile := filepath.Join(".context/tracks", trackName, "scratchpad.md")

			f, err := fs.OpenFile(scratchFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return fmt.Errorf("Error opening scratchpad: %v", err)
			}
			defer func() { _ = f.Close() }()

			if _, err := io.Copy(f, cmd.InOrStdin()); err != nil {
				return fmt.Errorf("Error writing to scratchpad: %v", err)
			}

			cmd.Printf(">> Data appended to %s\n", scratchFile)
			cmd.Println(">> INSTRUCTION: Do not read the whole file. Use grep/head/tail to find what you need.")
			return nil
		},
	}
}

func init() {
	rootCmd.AddCommand(NewDumpCmd(platform.NewRealFileSystem()))
}
