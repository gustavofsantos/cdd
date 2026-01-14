package cmd

import (
	"cdd/prompts"

	"github.com/spf13/cobra"
)

var (
	showSysPrompt   bool
	showBootPrompt  bool
	showCalibPrompt bool
)

var promptsCmd = &cobra.Command{
	Use:   "prompts",
	Short: "Output the various CDD prompts.",
	Long: `Output the essential CDD prompts to be shared with AI agents.

CDD relies on specific system instructions to guide AI agents through the 
protocol. Use this command to retrieve the core prompts.

FLAGS:
  --system       The primary system instructions for the CDD Engine.
  --bootstrap    Instructions for the initial state analysis (Phase 0).
  --calibration  A concise set of rules for continuous alignment.

EXAMPLES:
  $ cdd prompts --system > .context/SYSTEM_PROMPT.md
  $ cdd prompts --bootstrap`,
	Run: func(cmd *cobra.Command, args []string) {
		if showSysPrompt {
			cmd.Println(prompts.System)
			return
		}
		if showBootPrompt {
			cmd.Println(prompts.Bootstrap)
			return
		}

		if showCalibPrompt {
			cmd.Println(prompts.Calibration)
			return
		}

		// If no flag provided, show help
		_ = cmd.Help()
	},
}

func init() {
	promptsCmd.Flags().BoolVar(&showSysPrompt, "system", false, "Output the CDD System Prompt.")
	promptsCmd.Flags().BoolVar(&showBootPrompt, "bootstrap", false, "Output the Architect Prompt for initial setup.")
	promptsCmd.Flags().BoolVar(&showCalibPrompt, "calibration", false, "Output the Calibration Prompt.")

	rootCmd.AddCommand(promptsCmd)
}
