package prompts_test

import (
	"strings"
	"testing"

	"cdd/prompts"
)

func TestSystemPromptIsLean(t *testing.T) {
	requiredPhrases := []string{
		"Tests are Truth",
		"Treat this file as documentation",
		"Spec Cleanup",
		"## Test Reference",
	}

	for _, phrase := range requiredPhrases {
		if !strings.Contains(prompts.System, phrase) {
			t.Errorf("System prompt missing lean phase: %q", phrase)
		}
	}
}

func TestSystemPromptCommandsAndOverrides(t *testing.T) {
	requiredPhrases := []string{
		"cdd recite",
		"cdd start",
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

func TestSystemPromptHasCommitOften(t *testing.T) {

	requiredPhrases := []string{

		"Commit Often",

		"include the commit hash",
	}

	for _, phrase := range requiredPhrases {

		if !strings.Contains(prompts.System, phrase) {

			t.Errorf("System prompt missing commit often mandate: %q", phrase)

		}

	}

}
