package cmd

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"

	"cdd/internal/platform"
	"cdd/prompts"
)

// validateAntigravitySkill checks that the skill content has required Antigravity SKILL.md fields
func validateAntigravitySkill(content string) error {
	if !strings.HasPrefix(content, "---") {
		return fmt.Errorf("skill content must start with YAML frontmatter (---)")
	}

	// Extract frontmatter
	endIdx := strings.Index(content[3:], "---")
	if endIdx == -1 {
		return fmt.Errorf("skill content must have closing YAML frontmatter (---)")
	}

	frontmatter := content[3 : 3+endIdx]

	// Check for required fields: name and description
	if !strings.Contains(frontmatter, "name:") {
		return fmt.Errorf("skill frontmatter missing required field: name")
	}
	if !strings.Contains(frontmatter, "description:") {
		return fmt.Errorf("skill frontmatter missing required field: description")
	}

	return nil
}

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

func extractVersion(content string) string {
	re := regexp.MustCompile(`version:\s*["']?([^"'\s]+)["']?`)
	match := re.FindStringSubmatch(content)
	if len(match) > 1 {
		return match[1]
	}
	return "0.0.0"
}

func (s *skill) getVersion() string {
	return extractVersion(s.content)
}

func buildCursorRulesContent(skills []skill) string {
	var buf strings.Builder
	
	// Write version metadata header
	version := "1.0.0"
	if len(skills) > 0 {
		version = skills[0].getVersion()
	}
	buf.WriteString("---\n")
	buf.WriteString(fmt.Sprintf("version: %s\n", version))
	buf.WriteString("---\n\n")
	
	// Concatenate all skills
	for i, s := range skills {
		if i > 0 {
			buf.WriteString("\n---\n\n")
		}
		buf.WriteString(s.content)
	}
	
	return buf.String()
}

func installCursorRules(cmd *cobra.Command, fs platform.FileSystem, skills []skill) error {
	cursorRulesFile := ".cursorrules"
	content := buildCursorRulesContent(skills)
	currentVersion := extractVersion(content)

	// Check if file exists
	if info, err := fs.Stat(cursorRulesFile); err == nil && !info.IsDir() {
		existing, err := fs.ReadFile(cursorRulesFile)
		if err == nil {
			installedVersion := extractVersion(string(existing))

			if installedVersion == currentVersion {
				cmd.Printf("Cursor rules file '%s' is up to date (v%s)\n", cursorRulesFile, installedVersion)
				return nil
			}

			// Upgrade: backup and overwrite
			backupFile := cursorRulesFile + ".bak"
			if err := fs.Rename(cursorRulesFile, backupFile); err != nil {
				return fmt.Errorf("error backing up legacy cursor rules %s: %w", cursorRulesFile, err)
			}
			cmd.Printf("Updated cursor rules to v%s. Backup saved to %s\n", currentVersion, backupFile)
		}
	}

	if err := fs.WriteFile(cursorRulesFile, []byte(content), 0644); err != nil {
		return fmt.Errorf("error writing cursor rules file %s: %w", cursorRulesFile, err)
	}

	cmd.Printf("Cursor rules installed at %s\n", cursorRulesFile)
	return nil
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
			installedVersion := extractVersion(string(existing))

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

func installAntigravitySkill(cmd *cobra.Command, fs platform.FileSystem, s skill) error {
	// Validate skill content before installation
	if err := validateAntigravitySkill(s.content); err != nil {
		return fmt.Errorf("invalid skill content: %w", err)
	}

	skillDir := filepath.Join(".agent", "skills", s.id)
	if err := fs.MkdirAll(skillDir, 0755); err != nil {
		return fmt.Errorf("error creating skill directory %s: %w", skillDir, err)
	}

	skillFile := filepath.Join(skillDir, "SKILL.md")
	currentVersion := s.getVersion()

	// Check existing
	if info, err := fs.Stat(skillFile); err == nil && !info.IsDir() {
		existing, err := fs.ReadFile(skillFile)
		if err == nil {
			installedVersion := extractVersion(string(existing))

			if installedVersion == currentVersion {
				if cmd != nil {
					cmd.Printf("Agent Skill '%s' is up to date (v%s)\n", s.id, installedVersion)
				}
				return nil
			}

			// Upgrade: backup and overwrite
			backupFile := skillFile + ".bak"
			if err := fs.Rename(skillFile, backupFile); err != nil {
				return fmt.Errorf("error backing up skill %s: %w", s.id, err)
			}
			if cmd != nil {
				cmd.Printf("Updated Agent Skill '%s' to v%s. Backup saved to %s\n", s.id, currentVersion, backupFile)
			}
		}
	}

	if err := fs.WriteFile(skillFile, []byte(s.content), 0644); err != nil {
		return fmt.Errorf("error writing skill file %s: %w", s.id, err)
	}

	if cmd != nil {
		cmd.Printf("Agent Skill '%s' installed at %s\n", s.id, skillFile)
	}
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
  --target       Target directory for installation (agent, agents, claude, cursor, antigravity). Defaults to agent.

EXAMPLES:
  $ cdd agents --install
  $ cdd agents --install --target claude
  $ cdd agents --install --target cursor
  $ cdd agents --install --target antigravity`,
		Run: func(cmd *cobra.Command, args []string) {
			if installAgentSkill {
				skills := []skill{
					{id: "cdd", name: "cdd", description: "Orchestrator", content: prompts.System},
					{id: "cdd-analyst", name: "cdd-analyst", description: "Analyst", content: prompts.Analyst},
					{id: "cdd-architect", name: "cdd-architect", description: "Architect", content: prompts.Architect},
					{id: "cdd-executor", name: "cdd-executor", description: "Executor", content: prompts.Executor},
					{id: "cdd-integrator", name: "cdd-integrator", description: "Integrator", content: prompts.Integrator},
				}

				// Handle cursor target separately
				if installTarget == "cursor" {
					if err := installCursorRules(cmd, fs, skills); err != nil {
						cmd.PrintErrf("%v\n", err)
					}
					return
				}

				// Handle Antigravity target separately
				if installTarget == "antigravity" {
					for _, s := range skills {
						if err := installAntigravitySkill(cmd, fs, s); err != nil {
							cmd.PrintErrf("%v\n", err)
						}
					}
					return
				}

				// Handle directory-based targets
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
	agentsCmd.Flags().StringVar(&installTarget, "target", "agent", "Target directory for installation (agent, agents, claude, cursor, antigravity).")

	return agentsCmd
}

func init() {
	rootCmd.AddCommand(NewAgentsCmd(platform.NewRealFileSystem()))
}
