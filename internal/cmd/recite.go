package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var reciteCmd = &cobra.Command{
	Use:   "recite [track-name]",
	Short: "Output the current Plan to the context window.",
	Long: `Forces the agent to 'attend' to the immediate next step, preventing drift.
Usage: cdd recite <track-name>`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		trackName := args[0]
		planFile := filepath.Join(".context/tracks", trackName, "plan.md")

		content, err := os.ReadFile(planFile)
		if err != nil {
			fmt.Println("Error: Plan not found.")
			os.Exit(1)
		}

		fmt.Printf("=== RECITATION: %s ===\n", trackName)
		fmt.Println(string(content))
		fmt.Println("\n=== INSTRUCTION ===")
		fmt.Println("1. Identify the first unchecked item ([ ]).")
		fmt.Println("2. That is your ONLY focus for the next step.")
	},
}

func init() {
	rootCmd.AddCommand(reciteCmd)
}
