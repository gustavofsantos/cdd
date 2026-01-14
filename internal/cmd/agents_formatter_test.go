package cmd

import (
	"testing"

	"cdd/prompts"
)

func TestValidateAntigravitySkill(t *testing.T) {
	tests := []struct {
		name      string
		content   string
		wantError bool
		wantMsg   string
	}{
		{
			name: "valid skill with all required fields",
			content: `---
name: cdd-system
description: The Orchestrator that analyzes the plan and delegates to the appropriate Agent Skill.
metadata:
  version: 1.0.0
---
# Role: Orchestrator

Content here`,
			wantError: false,
		},
		{
			name: "valid skill with minimal fields",
			content: `---
name: cdd-analyst
description: Analyst skill
---
# Content`,
			wantError: false,
		},
		{
			name: "missing name field",
			content: `---
description: Analyst skill
metadata:
  version: 1.0.0
---
Content`,
			wantError: true,
			wantMsg:   "name",
		},
		{
			name: "missing description field",
			content: `---
name: cdd-analyst
metadata:
  version: 1.0.0
---
Content`,
			wantError: true,
			wantMsg:   "description",
		},
		{
			name: "missing frontmatter",
			content: `# Content without frontmatter`,
			wantError: true,
			wantMsg:   "frontmatter",
		},
		{
			name: "empty content",
			content: "",
			wantError: true,
			wantMsg:   "frontmatter",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateAntigravitySkill(tt.content)
			if (err != nil) != tt.wantError {
				t.Errorf("validateAntigravitySkill() error = %v, wantError %v", err, tt.wantError)
				return
			}
			if tt.wantError && tt.wantMsg != "" && err != nil {
				if !contains(err.Error(), tt.wantMsg) {
					t.Errorf("validateAntigravitySkill() error %q, want %q in message", err.Error(), tt.wantMsg)
				}
			}
		})
	}
}

func contains(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 && (s == substr || len(s) > len(substr) && (s[:len(substr)] == substr || s[len(s)-len(substr):] == substr || findSubstring(s, substr)))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func TestValidateAllCDDSkills(t *testing.T) {
	skillMap := map[string]string{
		"cdd-system":     prompts.System,
		"cdd-analyst":    prompts.Analyst,
		"cdd-architect":  prompts.Architect,
		"cdd-executor":   prompts.Executor,
		"cdd-integrator": prompts.Integrator,
	}

	for skillName, content := range skillMap {
		t.Run(skillName, func(t *testing.T) {
			err := validateAntigravitySkill(content)
			if err != nil {
				t.Errorf("CDD skill %q failed validation: %v", skillName, err)
			}
		})
	}
}
