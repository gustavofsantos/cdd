package cmd

import (
	"fmt"
	"path/filepath"
	"regexp"

	"github.com/spf13/cobra"

	"cdd/internal/platform"
	"cdd/prompts"
)

var (
	installAgentSkill bool
	installTarget     string
)

type skill struct {
	id          string
	name        string
	description string
	content     string
}

func (s *skill) getVersion() string {
	re := regexp.MustCompile(`version:\s*["']?([^"'\s]+)["']?`)
	match := re.FindStringSubmatch(s.content)
	if len(match) > 1 {
		return match[1]
	}
	return "0.0.0"
}

func installSkill(cmd *cobra.Command, fs platform.FileSystem, s skill, baseDir string) error {
	skillDir := filepath.Join(baseDir, "skills", s.id)
	if err := fs.MkdirAll(skillDir, 0755); err != nil {
		return fmt.Errorf("error creating skill directory %s: %w", skillDir, err)
	}

	skillFile := filepath.Join(skillDir, "SKILL.md")
	currentVersion := s.getVersion()

	// Check existing
	if info, err := fs.Stat(skillFile); err == nil && !info.IsDir() {
		existing, err := fs.ReadFile(skillFile)
		if err == nil {
			// Check version
			re := regexp.MustCompile(`version:\s*["']?([^"'\s]+)["']?`)
			match := re.FindStringSubmatch(string(existing))
			installedVersion := "0.0.0"
			if len(match) > 1 {
				installedVersion = match[1]
			}

			if installedVersion == currentVersion {
				cmd.Printf("Agent Skill '%s' is up to date (v%s)\n", s.id, installedVersion)
				return nil
			}

			// Migrate / Upgrade (Simple string equality check for now, can be improved to semver)
			// If different, we backup and overwrite.
			backupFile := skillFile + ".bak"
			if err := fs.Rename(skillFile, backupFile); err != nil {
				return fmt.Errorf("error backing up legacy skill %s: %w", s.id, err)
			}
			cmd.Printf("Updated Agent Skill '%s' to v%s. Backup saved to %s\n", s.id, currentVersion, backupFile)
		}
	}

	if err := fs.WriteFile(skillFile, []byte(s.content), 0644); err != nil {
		return fmt.Errorf("error writing skill file %s: %w", s.id, err)
	}

	cmd.Printf("Agent Skill '%s' installed at %s\n", s.id, skillFile)
	return nil
}

func NewAgentsCmd(fs platform.FileSystem) *cobra.Command {
	agentsCmd := &cobra.Command{
		Use:   "agents",
		Short: "Manage Agent Skills.",
		Long: `Manage Agent Skills to orchestrate AI following the CDD protocol.

Agent Skills are the primary way to extend the AI's capabilities and ensure 
it follows the Context-Driven Development methodology.

FLAGS:
  --install      Install all CDD Agent Skills (Orchestrator, Analyst, Architect, Executor, Integrator).
  --target       Target directory for installation (agent, agents, claude). Defaults to agent.

EXAMPLES:
  $ cdd agents --install
  $ cdd agents --install --target claude`,
		Run: func(cmd *cobra.Command, args []string) {
			if installAgentSkill {
				baseDir := ".agent"
				switch installTarget {
				case "claude":
					baseDir = ".claude"
				case "agents":
					baseDir = ".agents"
				case "agent":
					baseDir = ".agent"
				default:
					if installTarget != "" {
						cmd.PrintErrf("Warning: unknown target '%s', defaulting to '.agent'\n", installTarget)
					}
				}

				skills := []skill{
					{id: "cdd", name: "cdd", description: "Orchestrator", content: prompts.System},
					{id: "cdd-analyst", name: "cdd-analyst", description: "Analyst", content: prompts.Analyst},
					{id: "cdd-architect", name: "cdd-architect", description: "Architect", content: prompts.Architect},
					{id: "cdd-executor", name: "cdd-executor", description: "Executor", content: prompts.Executor},
					{id: "cdd-integrator", name: "cdd-integrator", description: "Integrator", content: prompts.Integrator},
				}

				for _, s := range skills {
					if err := installSkill(cmd, fs, s, baseDir); err != nil {
						cmd.PrintErrf("%v\n", err)
					}
				}
				return
			}

			// If no flag provided, show help
			_ = cmd.Help()
		},
	}

	agentsCmd.Flags().BoolVar(&installAgentSkill, "install", false, "Install the CDD System Prompt as an Agent Skill.")
	agentsCmd.Flags().StringVar(&installTarget, "target", "agent", "Target directory for installation (agent, agents, claude).")

	return agentsCmd
}

func init() {
	rootCmd.AddCommand(NewAgentsCmd(platform.NewRealFileSystem()))
}
