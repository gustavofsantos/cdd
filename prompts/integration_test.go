package prompts_test

import (
	"testing"

	"cdd/prompts"
)

func TestCorePromptsAreLoaded(t *testing.T) {
	if prompts.System == "" {
		t.Error("System prompt is empty")
	}
	if prompts.Surveyor == "" {
		t.Error("Surveyor prompt is empty")
	}
	if prompts.Analyst == "" {
		t.Error("Analyst prompt is empty")
	}
	if prompts.Architect == "" {
		t.Error("Architect prompt is empty")
	}
	if prompts.Executor == "" {
		t.Error("Executor prompt is empty")
	}
	if prompts.Integrator == "" {
		t.Error("Integrator prompt is empty")
	}
}
