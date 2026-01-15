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

If only one track exists, it will be used automatically. When multiple tracks 
are active, you must specify which one to recite.

EXAMPLES:
   $ cdd recite user-authentication
   $ cdd recite fix-bug-123
   $ cdd recite (uses the track automatically if only one exists)`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var trackName string

			// If track name is provided, use it
			if len(args) > 0 {
				trackName = args[0]
			} else {
				// Infer track name if not provided
				tasks, err := GetActiveTasks(fs)
				if err != nil {
					return fmt.Errorf("Error: Could not list tracks.")
				}

				if len(tasks) == 0 {
					return fmt.Errorf("Error: No active tracks found. Create one with 'cdd start <track-name>'.")
				}

				if len(tasks) > 1 {
					cmd.PrintErrf("Error: multiple active tracks found (%d), please specify one\n\n", len(tasks))
					cmd.Println(RenderTaskSelectionMenu(tasks))
					return fmt.Errorf("track name required")
				}

				trackName = tasks[0]
			}

			planFile := filepath.Join(".context/tracks", trackName, "plan.md")

			content, err := fs.ReadFile(planFile)
			if err != nil {
				return fmt.Errorf("Error: Plan not found for track '%s'.", trackName)
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
