package cmd

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/spf13/cobra"

	"cdd/internal/platform"
	"cdd/prompts"
)

var (
	showSysPrompt   bool
	showBootPrompt  bool
	showCalibPrompt bool
	installSkill    bool
)

func NewPromptsCmd(fs platform.FileSystem) *cobra.Command {
	promptsCmd := &cobra.Command{
		Use:   "prompts",
		Short: "Output the various CDD prompts.",
		Long: `Output the essential CDD prompts to be shared with AI agents.

CDD relies on specific system instructions to guide AI agents through the 
protocol. Use this command to retrieve the core prompts.

FLAGS:
  --system       The primary system instructions for the CDD Engine.
  --bootstrap    Instructions for the initial state analysis (Phase 0).
  --calibration  A concise set of rules for continuous alignment.
  --install      Install the CDD System Prompt as an Agent Skill (.agent/skills/cdd/SKILL.md).

EXAMPLES:
  $ cdd prompts --system > .context/SYSTEM_PROMPT.md
  $ cdd prompts --bootstrap
  $ cdd prompts --install`,
		Run: func(cmd *cobra.Command, args []string) {
			if installSkill {
				skillDir := ".agent/skills/cdd"
				if err := fs.MkdirAll(skillDir, 0755); err != nil {
					cmd.PrintErrf("Error creating skill directory: %v\n", err)
					return
				}

				skillFile := filepath.Join(skillDir, "SKILL.md")
				currentVersion := 2

				// Check existing
				if info, err := fs.Stat(skillFile); err == nil && !info.IsDir() {
					existing, err := fs.ReadFile(skillFile)
					if err == nil {
						// Check version
						re := regexp.MustCompile(`version:\s*(\d+)`)
						match := re.FindSubmatch(existing)
						installedVersion := 0
						if len(match) > 1 {
							installedVersion, _ = strconv.Atoi(string(match[1]))
						}

						if installedVersion >= currentVersion {
							cmd.Printf("Skill 'cdd' is up to date (v%d)\n", installedVersion)
							return
						}

						// Migrate
						backupFile := skillFile + ".bak"
						if err := fs.Rename(skillFile, backupFile); err != nil {
							cmd.PrintErrf("Error backing up legacy skill: %v\n", err)
							return
						}
						cmd.Printf("Migrated legacy skill to v%d. Backup saved to %s\n", currentVersion, backupFile)
					}
				}

				frontmatter := fmt.Sprintf("---\nname: cdd\nversion: %d\ndescription: Protocol for implementing software features using the Context-Driven Development methodology.\n---\n\n", currentVersion)
				content := frontmatter + prompts.System

				if err := fs.WriteFile(skillFile, []byte(content), 0644); err != nil {
					cmd.PrintErrf("Error writing skill file: %v\n", err)
					return
				}

				cmd.Printf("Skill 'cdd' installed at %s\n", skillFile)
				return
			}
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

	promptsCmd.Flags().BoolVar(&showSysPrompt, "system", false, "Output the CDD System Prompt.")
	promptsCmd.Flags().BoolVar(&showBootPrompt, "bootstrap", false, "Output the Architect Prompt for initial setup.")
	promptsCmd.Flags().BoolVar(&showCalibPrompt, "calibration", false, "Output the Calibration Prompt.")
	promptsCmd.Flags().BoolVar(&installSkill, "install", false, "Install the CDD System Prompt as an Agent Skill.")

	return promptsCmd
}

func init() {
	rootCmd.AddCommand(NewPromptsCmd(platform.NewRealFileSystem()))
}
