package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cdd",
	Short: "CDD Tool Suite",
	Long: `Context-Driven Development (CDD) Tool Suite.

CDD is a methodology that puts the "Context" (Specification, Plan, and Decisions) 
at the center of the development lifecycle. This tool suite automates the 
management of ephemeral "Tracks" (workspaces) and the integration of their 
results into the global project context.

CORE PRINCIPLES:
1. Tracks are Ephemeral: Work happens in isolated folders under .context/tracks/.
2. Specs are Eternal: The source of truth is maintained in .context/specs/.


WORKFLOW:
1. start: Create a new track for a specific feature or bugfix.
2. analyze: (Developer/Agent) Fill in the spec.md and plan.md.
3. execute: (Developer/Agent) Implement changes following the plan.
4. archive: Move the track to history.

EXAMPLES:
  $ cdd init                 # Initialize a new project with CDD
  $ cdd start user-auth      # Start a new track for "user-auth"

  $ cdd archive user-auth    # Complete the track and move to history`,
}

func init() {
	rootCmd.SetOut(os.Stdout)
	rootCmd.SetErr(os.Stderr)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
