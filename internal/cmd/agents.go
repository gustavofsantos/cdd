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
	installAgentSkill bool
)

func NewAgentsCmd(fs platform.FileSystem) *cobra.Command {
	agentsCmd := &cobra.Command{
		Use:   "agents",
		Short: "Manage Agent Skills.",
		Long: `Manage Agent Skills to orchestrate AI following the CDD protocol.

Agent Skills are the primary way to extend the AI's capabilities and ensure 
it follows the Context-Driven Development methodology.

FLAGS:
  --install      Install the CDD System Prompt as an Agent Skill (.agent/skills/cdd/SKILL.md).

EXAMPLES:
  $ cdd agents --install`,
		Run: func(cmd *cobra.Command, args []string) {
			if installAgentSkill {
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
						re := regexp.MustCompile(`version:\s*["']?(\d+)["']?`)
						match := re.FindSubmatch(existing)
						installedVersion := 0
						if len(match) > 1 {
							installedVersion, _ = strconv.Atoi(string(match[1]))
						}

						if installedVersion >= currentVersion {
							cmd.Printf("Agent Skill 'cdd' is up to date (v%d)\n", installedVersion)
							return
						}

						// Migrate
						backupFile := skillFile + ".bak"
						if err := fs.Rename(skillFile, backupFile); err != nil {
							cmd.PrintErrf("Error backing up legacy skill: %v\n", err)
							return
						}
						cmd.Printf("Migrated legacy Agent Skill to v%d. Backup saved to %s\n", currentVersion, backupFile)
					}
				}

				frontmatter := fmt.Sprintf("---\nname: cdd\ndescription: Protocol for implementing software features using the Context-Driven Development methodology.\nmetadata:\n  version: \"%d\"\n---\n\n", currentVersion)
				content := frontmatter + prompts.System

				if err := fs.WriteFile(skillFile, []byte(content), 0644); err != nil {
					cmd.PrintErrf("Error writing skill file: %v\n", err)
					return
				}

				cmd.Printf("Agent Skill 'cdd' installed at %s\n", skillFile)
				return
			}

			// If no flag provided, show help
			_ = cmd.Help()
		},
	}

	agentsCmd.Flags().BoolVar(&installAgentSkill, "install", false, "Install the CDD System Prompt as an Agent Skill.")

	return agentsCmd
}

func init() {
	rootCmd.AddCommand(NewAgentsCmd(platform.NewRealFileSystem()))
}
