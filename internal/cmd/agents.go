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

// validateAntigravityWorkflow checks that the workflow content has required Antigravity frontmatter
func validateAntigravityWorkflow(content string) error {
	if !strings.HasPrefix(content, "---") {
		return fmt.Errorf("workflow content must start with YAML frontmatter (---)")
	}

	// Extract frontmatter
	endIdx := strings.Index(content[3:], "---")
	if endIdx == -1 {
		return fmt.Errorf("workflow content must have closing YAML frontmatter (---)")
	}

	frontmatter := content[3 : 3+endIdx]

	// Check for required fields: name and description
	// Although Antigravity Workflows documentation might specify different required fields,
	// for now we enforce name and description as they are present in our prompts.
	if !strings.Contains(frontmatter, "name:") {
		return fmt.Errorf("workflow frontmatter missing required field: name")
	}
	if !strings.Contains(frontmatter, "description:") {
		return fmt.Errorf("workflow frontmatter missing required field: description")
	}

	return nil
}

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

func printTargetRequiredError(cmd *cobra.Command) {
	cmd.PrintErrf("Error: you must specify a target with --target or use --all to install for all platforms\n\n")
	cmd.PrintErrf("Available targets: agent, agents, claude, cursor, antigravity\n\n")
	cmd.PrintErrf("Examples:\n")
	cmd.PrintErrf("  Install to a specific target:\n")
	cmd.PrintErrf("    cdd agents --install --target agent\n")
	cmd.PrintErrf("    cdd agents --install --target claude\n\n")
	cmd.PrintErrf("  Install to all platforms:\n")
	cmd.PrintErrf("    cdd agents --install --all\n")
}

func installSkillsForAllPlatforms(cmd *cobra.Command, fs platform.FileSystem, skills []skill) error {
	// Install for all directory-based platforms
	platforms := []struct {
		name    string
		baseDir string
	}{
		{"agent", ".agent"},
		{"claude", ".claude"},
		{"agents", ".agents"},
	}

	for _, p := range platforms {
		for _, s := range skills {
			if err := installSkill(cmd, fs, s, p.baseDir); err != nil {
				cmd.PrintErrf("%v\n", err)
			}
		}
	}

	// Also install for cursor and antigravity
	if err := installCursorRules(cmd, fs, skills); err != nil {
		cmd.PrintErrf("%v\n", err)
	}

	for _, s := range skills {
		if err := installAntigravityWorkflow(cmd, fs, s); err != nil {
			cmd.PrintErrf("%v\n", err)
		}
	}

	return nil
}

func installCursorRules(cmd *cobra.Command, fs platform.FileSystem, skills []skill) error {
	rulesDir := filepath.Join(".cursor", "rules")

	// Create .cursor/rules directory
	if err := fs.MkdirAll(rulesDir, 0755); err != nil {
		return fmt.Errorf("error creating cursor rules directory %s: %w", rulesDir, err)
	}

	// Install each skill as individual rule file
	for _, s := range skills {
		ruleFile := filepath.Join(rulesDir, s.id+".mdc")
		currentVersion := extractVersion(s.content)

		// Check if file exists
		if info, err := fs.Stat(ruleFile); err == nil && !info.IsDir() {
			existing, err := fs.ReadFile(ruleFile)
			if err == nil {
				installedVersion := extractVersion(string(existing))

				if installedVersion == currentVersion {
					cmd.Printf("Cursor rule '%s' is up to date (v%s)\n", s.id, installedVersion)
					continue
				}

				// Upgrade: backup and overwrite
				backupFile := ruleFile + ".bak"
				if err := fs.Rename(ruleFile, backupFile); err != nil {
					return fmt.Errorf("error backing up legacy cursor rule %s: %w", s.id, err)
				}
				cmd.Printf("Updated cursor rule '%s' to v%s. Backup saved to %s\n", s.id, currentVersion, backupFile)
			}
		}

		if err := fs.WriteFile(ruleFile, []byte(s.content), 0644); err != nil {
			return fmt.Errorf("error writing cursor rule file %s: %w", ruleFile, err)
		}

		cmd.Printf("Cursor rule '%s' installed at %s\n", s.id, ruleFile)
	}

	// Migrate legacy .cursorrules file if it exists
	legacyFile := ".cursorrules"
	if info, err := fs.Stat(legacyFile); err == nil && !info.IsDir() {
		backupFile := filepath.Join(rulesDir, ".cursorrules.legacy.bak")
		if err := fs.Rename(legacyFile, backupFile); err != nil {
			cmd.PrintErrf("Warning: could not migrate legacy .cursorrules file: %v\n", err)
		} else {
			cmd.Printf("Migrated legacy .cursorrules to %s\n", backupFile)
		}
	}

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

func installAntigravityWorkflow(cmd *cobra.Command, fs platform.FileSystem, s skill) error {
	// Validate workflow content before installation
	if err := validateAntigravityWorkflow(s.content); err != nil {
		return fmt.Errorf("invalid workflow content: %w", err)
	}

	workflowDir := filepath.Join(".agent", "workflows")
	if err := fs.MkdirAll(workflowDir, 0755); err != nil {
		return fmt.Errorf("error creating workflow directory %s: %w", workflowDir, err)
	}

	workflowFile := filepath.Join(workflowDir, s.id+".md")
	currentVersion := s.getVersion()

	// Check existing
	if info, err := fs.Stat(workflowFile); err == nil && !info.IsDir() {
		existing, err := fs.ReadFile(workflowFile)
		if err == nil {
			installedVersion := extractVersion(string(existing))

			if installedVersion == currentVersion {
				if cmd != nil {
					cmd.Printf("Antigravity Workflow '%s' is up to date (v%s)\n", s.id, installedVersion)
				}
				return nil
			}

			// Upgrade: backup and overwrite
			backupFile := workflowFile + ".bak"
			if err := fs.Rename(workflowFile, backupFile); err != nil {
				return fmt.Errorf("error backing up workflow %s: %w", s.id, err)
			}
			if cmd != nil {
				cmd.Printf("Updated Antigravity Workflow '%s' to v%s. Backup saved to %s\n", s.id, currentVersion, backupFile)
			}
		}
	}

	if err := fs.WriteFile(workflowFile, []byte(s.content), 0644); err != nil {
		return fmt.Errorf("error writing workflow file %s: %w", s.id, err)
	}

	if cmd != nil {
		cmd.Printf("Antigravity Workflow '%s' installed at %s\n", s.id, workflowFile)
	}
	return nil
}

func NewAgentsCmd(fs platform.FileSystem) *cobra.Command {
	var (
		installAgentSkill   bool
		installTarget       string
		installAllPlatforms bool
	)

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
				// Validate that either --all or --target is provided
				if !installAllPlatforms && installTarget == "" {
					printTargetRequiredError(cmd)
					return
				}

				skills := []skill{
					{id: "cdd", name: "cdd", description: "Orchestrator", content: prompts.System},
					{id: "cdd-surveyor", name: "cdd-surveyor", description: "Surveyor", content: prompts.Surveyor},
					{id: "cdd-analyst", name: "cdd-analyst", description: "Analyst", content: prompts.Analyst},
					{id: "cdd-architect", name: "cdd-architect", description: "Architect", content: prompts.Architect},
					{id: "cdd-executor", name: "cdd-executor", description: "Executor", content: prompts.Executor},
					{id: "cdd-integrator", name: "cdd-integrator", description: "Integrator", content: prompts.Integrator},
				}

				// Handle --all flag for all platforms
				if installAllPlatforms {
					if err := installSkillsForAllPlatforms(cmd, fs, skills); err != nil {
						cmd.PrintErrf("%v\n", err)
					}
					return
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
						if err := installAntigravityWorkflow(cmd, fs, s); err != nil {
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
	agentsCmd.Flags().StringVar(&installTarget, "target", "", "Target directory for installation (agent, agents, claude, cursor, antigravity).")
	agentsCmd.Flags().BoolVar(&installAllPlatforms, "all", false, "Install skills for all supported platforms.")

	return agentsCmd
}

func init() {
	rootCmd.AddCommand(NewAgentsCmd(platform.NewRealFileSystem()))
}
