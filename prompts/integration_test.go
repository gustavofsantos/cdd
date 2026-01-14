package prompts_test

import (
	"testing"

	"cdd/prompts"
)

func TestCorePromptsAreLoaded(t *testing.T) {
	if prompts.System == "" {
		t.Error("System prompt is empty")
	}
	if prompts.Analyst == "" {
		t.Error("Analyst prompt is empty")
	}
}
