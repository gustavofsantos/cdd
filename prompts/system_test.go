package prompts_test

import (
	"strings"
	"testing"

	"cdd/prompts"
)

func TestSystemPromptIsLean(t *testing.T) {
	requiredPhrases := []string{
		"The Lean CDD Protocol",
		"strict 3-file Track structure",
		"spec.md",
		"(The Delta)",
		"plan.md",
		"(The Execution)",
		"decisions.md",
		"(The How)",
	}

	for _, phrase := range requiredPhrases {
		if !strings.Contains(prompts.System, phrase) {
			t.Errorf("System prompt missing philosophy phase: %q", phrase)
		}
	}
}

func TestSystemPromptCommandsAndOverrides(t *testing.T) {
	requiredPhrases := []string{
		"cdd recite",
		"AGENTS.local.md",
	}

	for _, phrase := range requiredPhrases {
		if !strings.Contains(prompts.System, phrase) {
			t.Errorf("System prompt missing command/override phase: %q", phrase)
		}
	}

	// Ensure we are NOT using go run in the system prompt instructions anymore
	if strings.Contains(prompts.System, "go run cmd/cdd/main.go") {
		t.Errorf("System prompt still contains 'go run cmd/cdd/main.go', it should use 'cdd' directly")
	}
}

func TestSystemPromptConstraints(t *testing.T) {
	requiredPhrases := []string{
		"Global Constraints",
		"Files:",
		"Lifecycle:",
	}

	for _, phrase := range requiredPhrases {
		if !strings.Contains(prompts.System, phrase) {
			t.Errorf("System prompt missing constraint: %q", phrase)
		}
	}
}
