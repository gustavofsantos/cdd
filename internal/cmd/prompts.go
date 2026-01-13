package cmd

import (
	"cdd/prompts"

	"github.com/spf13/cobra"
)

var (
	showSysPrompt   bool
	showBootPrompt  bool
	showInPrompt    bool
	showExecPrompt  bool
	showPlanPrompt  bool
	showCalibPrompt bool
)

var promptsCmd = &cobra.Command{
	Use:   "prompts",
	Short: "Output the various CDD prompts.",
	Long:  `Retrieve the CDD System, Bootstrap, Inbox, Executor, Planner, or Calibration prompts.`,
	Run: func(cmd *cobra.Command, args []string) {
		if showSysPrompt {
			cmd.Println(prompts.System)
			return
		}
		if showBootPrompt {
			cmd.Println(prompts.Bootstrap)
			return
		}
		if showInPrompt {
			cmd.Println(prompts.Inbox)
			return
		}
		if showExecPrompt {
			cmd.Println(prompts.Executor)
			return
		}
		if showPlanPrompt {
			cmd.Println(prompts.Planner)
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
	promptsCmd.Flags().BoolVar(&showInPrompt, "inbox", false, "Output the Context Gardener Prompt.")
	promptsCmd.Flags().BoolVar(&showExecPrompt, "executor", false, "Output the Executor Prompt.")
	promptsCmd.Flags().BoolVar(&showPlanPrompt, "planner", false, "Output the Planner Prompt.")
	promptsCmd.Flags().BoolVar(&showCalibPrompt, "calibration", false, "Output the Calibration Prompt.")

	rootCmd.AddCommand(promptsCmd)
}
