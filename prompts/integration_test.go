package prompts_test

import (
	"cdd/prompts"
	"testing"
)

func TestExecutorAndPlannerPrompts(t *testing.T) {
	if prompts.Executor == "" {
		t.Error("Executor prompt is empty")
	}
	if prompts.Planner == "" {
		t.Error("Planner prompt is empty")
	}
}
