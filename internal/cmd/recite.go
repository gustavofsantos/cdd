package cmd

import (
	"fmt"
	"path/filepath"

	"cdd/internal/platform"

	"github.com/spf13/cobra"
)

func NewReciteCmd(fs platform.FileSystem) *cobra.Command {
	return &cobra.Command{
		Use:   "recite [track-name]",
		Short: "Output the current Plan to the context window.",
		Long: `Output the current track's plan to focus the agent's attention.

The 'recite' command is a calibration tool used to prevent "context drift". 
By outputting the plan and reminding the agent of their immediate next task, 
it ensures that development stays aligned with the intended specification.

EXAMPLES:
  $ cdd recite user-authentication
  $ cdd recite fix-bug-123`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			trackName := args[0]
			planFile := filepath.Join(".context/tracks", trackName, "plan.md")

			content, err := fs.ReadFile(planFile)
			if err != nil {
				return fmt.Errorf("Error: Plan not found.")
			}

			cmd.Printf("=== RECITATION: %s ===\n", trackName)
			cmd.Println(string(content))
			cmd.Println("\n=== INSTRUCTION ===")
			cmd.Println("1. Identify the first unchecked item ([ ]).")
			cmd.Println("2. That is your ONLY focus for the next step.")
			return nil
		},
	}
}

func init() {
	rootCmd.AddCommand(NewReciteCmd(platform.NewRealFileSystem()))
}
