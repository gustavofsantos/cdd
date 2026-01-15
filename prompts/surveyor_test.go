package prompts_test

import (
	"strings"
	"testing"

	"cdd/prompts"
)

func TestSurveyorPromptHasFrontmatter(t *testing.T) {
	if !strings.HasPrefix(prompts.Surveyor, "---") {
		t.Error("Surveyor prompt must start with YAML frontmatter (---)")
	}

	if !strings.Contains(prompts.Surveyor, "name: cdd-surveyor") {
		t.Error("Surveyor prompt frontmatter missing 'name: cdd-surveyor'")
	}

	if !strings.Contains(prompts.Surveyor, "description:") {
		t.Error("Surveyor prompt frontmatter missing 'description'")
	}

	if !strings.Contains(prompts.Surveyor, "metadata:") {
		t.Error("Surveyor prompt frontmatter missing 'metadata'")
	}

	if !strings.Contains(prompts.Surveyor, "version:") {
		t.Error("Surveyor prompt metadata missing 'version'")
	}
}

func TestSurveyorPromptHasContent(t *testing.T) {
	requiredSections := []string{
		"Role: Surveyor",
		"Objective",
	}

	for _, section := range requiredSections {
		if !strings.Contains(prompts.Surveyor, section) {
			t.Errorf("Surveyor prompt missing section: %q", section)
		}
	}
}
