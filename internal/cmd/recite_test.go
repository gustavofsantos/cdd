package cmd_test

import (
	"bytes"
	"testing"

	"cdd/internal/cmd"
	"cdd/internal/platform"
)

func TestReciteCmd_DisplaysPlan(t *testing.T) {
	fs := platform.NewMockFileSystem()
	// Setup plan
	planContent := "# My Plan"
	_ = fs.WriteFile(".context/tracks/active-track/plan.md", []byte(planContent), 0644)

	command := cmd.NewReciteCmd(fs)
	buf := new(bytes.Buffer)
	command.SetOut(buf)
	command.SetErr(buf)

	command.SetArgs([]string{"active-track"})
	err := command.Execute()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	output := buf.String()
	expectedHeader := "=== RECITATION: active-track ===\n"
	if !contains(output, expectedHeader) {
		t.Errorf("expected output to contain '%s', got '%s'", expectedHeader, output)
	}
	if !contains(output, planContent) {
		t.Errorf("expected output to contain plan content, got '%s'", output)
	}
}

func contains(s, substr string) bool {
	return bytes.Contains([]byte(s), []byte(substr))
}
