package prompts_test

import (
	"testing"

	"cdd/prompts"
)

func TestExecutorAndPlannerPrompts(t *testing.T) {
	if prompts.Executor == "" {
		t.Error("Executor prompt is empty")
	}
	if prompts.Planner == "" {
		t.Error("Planner prompt is empty")
	}
	if prompts.Calibration == "" {
		t.Error("Calibration prompt is empty")
	}
	if prompts.Integrator == "" {
		t.Error("Integrator prompt is empty")
	}
}
