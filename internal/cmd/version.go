package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "dev"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of cdd",
	Long: `Display the current version of the CDD Tool Suite.

Useful for debugging environment issues and ensuring you are using the 
latest recommended protocol version.

EXAMPLES:
  $ cdd version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("cdd version %s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
