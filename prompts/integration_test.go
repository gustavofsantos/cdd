package prompts_test

import (
	"testing"

	"cdd/prompts"
)

func TestExecutorAndPlannerPrompts(t *testing.T) {
	if prompts.Calibration == "" {
		t.Error("Calibration prompt is empty")
	}
}
