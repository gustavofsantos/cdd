package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cdd",
	Short: "CDD Tool Suite",
	Long:  `Context-Driven Development (CDD) Tool Suite.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
