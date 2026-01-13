package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Bootstrap the CDD environment.",
	Long:  `Creates the persistent memory structure (.context/) and starts the 'setup' track.`,
	Run: func(cmd *cobra.Command, args []string) {

		cmd.Println("Initializing Context-Driven Environment...")
		dirs := []string{
			".context/tracks",
			".context/archive",
			".context/specs",
		}
		for _, d := range dirs {
			if err := os.MkdirAll(d, 0755); err != nil {
				fmt.Printf("Error creating directory %s: %v\n", d, err)
				os.Exit(1)
			}
		}

		data := trackData{
			TrackName: "setup",
			CreatedAt: time.Now().Format("Mon Jan 2 15:04:05 MST 2006"),
		}

		// Create default files
		files := []struct {
			Destination string
			Template    string
		}{
			{".context/product.md", "init_product.md"},
			{".context/tech-stack.md", "init_tech-stack.md"},
			{".context/architecture.md", "init_architecture.md"},
		}

		for _, f := range files {
			if _, err := os.Stat(f.Destination); os.IsNotExist(err) {
				content, err := renderTrackTemplate(filepath.Base(f.Destination), f.Template, data)
				if err != nil {
					fmt.Printf("Error rendering %s: %v\n", f.Destination, err)
					continue
				}
				if err := os.WriteFile(f.Destination, content, 0644); err != nil {
					fmt.Printf("Error writing %s: %v\n", f.Destination, err)
				} else {
					cmd.Printf("Created '%s'.\n", f.Destination)
				}
			}
		}

		// Auto-start setup track
		setupDir := ".context/tracks/setup"
		if _, err := os.Stat(setupDir); os.IsNotExist(err) {
			if err := os.MkdirAll(setupDir, 0755); err != nil {
				fmt.Printf("Error creating setup track: %v\n", err)
				os.Exit(1)
			}

			// Spec
			specContent, err := renderTrackTemplate("spec.md", "setup_spec.md", data)
			if err != nil {
				fmt.Printf("Error rendering spec.md: %v\n", err)
			} else {
				if err := os.WriteFile(filepath.Join(setupDir, "spec.md"), specContent, 0644); err != nil {
					fmt.Printf("Error writing spec.md: %v\n", err)
				}
			}

			// Plan
			planContent, err := renderTrackTemplate("plan.md", "setup_plan.md", data)
			if err != nil {
				fmt.Printf("Error rendering plan.md: %v\n", err)
			} else {
				if err := os.WriteFile(filepath.Join(setupDir, "plan.md"), planContent, 0644); err != nil {
					fmt.Printf("Error writing plan.md: %v\n", err)
				}
			}

			// Decisions
			decisionsContent, err := renderTrackTemplate("decisions.md", "decisions.md", data)
			if err != nil {
				fmt.Printf("Error rendering decisions.md: %v\n", err)
			} else {
				if err := os.WriteFile(filepath.Join(setupDir, "decisions.md"), decisionsContent, 0644); err != nil {
					fmt.Printf("Error writing decisions.md: %v\n", err)
				}
			}

			cmd.Println("Created 'setup' track.")
		}

		cmd.Println("CDD Initialized.")
		cmd.Println("👉 Run 'cdd prompts --bootstrap' to get the prompt for the next step.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
