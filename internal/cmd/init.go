package cmd

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	"cdd/prompts"

	"github.com/spf13/cobra"
)

var systemPrompt = prompts.System
var bootstrapPrompt = prompts.Bootstrap
var inboxPrompt = prompts.Inbox

var (
	showSystemPrompt    bool
	showBootstrapPrompt bool
	showInboxPrompt     bool
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Bootstrap the CDD environment.",
	Long: `Creates the persistent memory structure (.context/) and starts the 'setup' track.

Flags:
  --system-prompt     Output the CDD System Prompt for your AI agent.
  --bootstrap-prompt  Output the Architect Prompt for initial setup.
  --inbox-prompt      Output the Context Gardener Prompt for processing the inbox.`,
	Run: func(cmd *cobra.Command, args []string) {
		if showSystemPrompt {
			fmt.Println(systemPrompt)
			return
		}
		if showBootstrapPrompt {
			fmt.Println(bootstrapPrompt)
			return
		}
		if showInboxPrompt {
			fmt.Println(inboxPrompt)
			return
		}

		fmt.Println("Initializing Context-Driven Environment...")
		dirs := []string{
			".context/tracks",
			".context/archive",
			".context/features",
		}
		for _, d := range dirs {
			if err := os.MkdirAll(d, 0755); err != nil {
				fmt.Printf("Error creating directory %s: %v\n", d, err)
				os.Exit(1)
			}
		}

		files := []string{
			".context/product.md",
			".context/tech-stack.md",
			".context/workflow.md",
			".context/patterns.md",
			".context/inbox.md",
		}
		for _, f := range files {
			if _, err := os.Stat(f); os.IsNotExist(err) {
				file, err := os.Create(f)
				if err != nil {
					fmt.Printf("Error creating file %s: %v\n", f, err)
				} else {
					file.Close()
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

			specContent := "# Spec: Project Initialization\nThis track is dedicated to establishing the global project context.\n"
			if err := os.WriteFile(filepath.Join(setupDir, "spec.md"), []byte(specContent), 0644); err != nil {
				fmt.Printf("Error writing spec.md: %v\n", err)
			}

			planContent := `# Plan for setup
- [ ] Phase 1: Automated Archeology (Scan file structure & configs)
- [ ] Phase 2: Draft Context (Create provisional product.md/tech-stack.md)
- [ ] Phase 3: Focus Alignment (Ask user for their specific Domain of Interest)
- [ ] Phase 4: Deep Dive (Scan the specific Domain of Interest)
- [ ] Phase 5: Final Review & Archive
`
			if err := os.WriteFile(filepath.Join(setupDir, "plan.md"), []byte(planContent), 0644); err != nil {
				fmt.Printf("Error writing plan.md: %v\n", err)
			}

			if err := os.WriteFile(filepath.Join(setupDir, "decisions.md"), []byte("# Decision Log\n"), 0644); err != nil {
				fmt.Printf("Error writing decisions.md: %v\n", err)
			}
			if err := os.WriteFile(filepath.Join(setupDir, "scratchpad.md"), []byte("# Scratchpad\n"), 0644); err != nil {
				fmt.Printf("Error writing scratchpad.md: %v\n", err)
			}

			fmt.Println("Created 'setup' track.")
		}

		fmt.Println("CDD Initialized.")
		fmt.Println("👉 Run 'cdd init --bootstrap-prompt' to get the prompt for the next step.")
	},
}

func init() {
	initCmd.Flags().BoolVar(&showSystemPrompt, "system-prompt", false, "Output the CDD System Prompt for your AI agent.")
	initCmd.Flags().BoolVar(&showBootstrapPrompt, "bootstrap-prompt", false, "Output the Architect Prompt for initial setup.")
	initCmd.Flags().BoolVar(&showInboxPrompt, "inbox-prompt", false, "Output the Context Gardener Prompt for processing the inbox.")
	rootCmd.AddCommand(initCmd)
}
