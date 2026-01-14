package prompts_test

import (
	"strings"
	"testing"

	"cdd/prompts"
)

func TestSystemPromptIsLean(t *testing.T) {
	requiredPhrases := []string{
		"CDD Engine",
		".context/tracks",
		"spec.md",
		"plan.md",
		"decisions.md",
		"THE STATE MACHINE",
		"PHASE INSTRUCTIONS",
	}

	for _, phrase := range requiredPhrases {
		if !strings.Contains(prompts.System, phrase) {
			t.Errorf("System prompt missing phrase: %q", phrase)
		}
	}
}

func TestSystemPromptCommandsAndOverrides(t *testing.T) {
	requiredPhrases := []string{
		"Agent Skills",
	}

	for _, phrase := range requiredPhrases {
		if !strings.Contains(prompts.System, phrase) {
			t.Errorf("System prompt missing phrase: %q", phrase)
		}
	}

	forbiddenPhrases := []string{
		"AGENTS.local.md",
		"GEMINI.md",
		"CLAUDE.md",
	}

	for _, phrase := range forbiddenPhrases {
		if strings.Contains(prompts.System, phrase) {
			t.Errorf("System prompt contains forbidden external reference: %q", phrase)
		}
	}

	// Ensure we are NOT using go run in the system prompt instructions anymore
	if strings.Contains(prompts.System, "go run cmd/cdd/main.go") {
		t.Errorf("System prompt still contains 'go run cmd/cdd/main.go', it should use 'cdd' directly")
	}
}

func TestSystemPromptConstraints(t *testing.T) {
	requiredPhrases := []string{
		"GLOBAL CONSTRAINTS",
		"Atomic Steps:",
	}

	for _, phrase := range requiredPhrases {
		if !strings.Contains(prompts.System, phrase) {
			t.Errorf("System prompt missing constraint: %q", phrase)
		}
	}
}
