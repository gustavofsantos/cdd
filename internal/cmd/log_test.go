package cmd_test

import (
	"bytes"
	"fmt"
	"testing"

	"cdd/internal/cmd"
	"cdd/internal/platform"
)

func TestLogCmd_AppendsToLog(t *testing.T) {
	fs := platform.NewMockFileSystem()
	// Setup existing log
	logPath := ".context/tracks/active-track/decisions.md"
	initialContent := "# Decision Log\n"
	_ = fs.WriteFile(logPath, []byte(initialContent), 0644)

	command := cmd.NewLogCmd(fs)
	buf := new(bytes.Buffer)
	command.SetOut(buf)
	command.SetErr(buf)

	msg := "Test Decision"
	command.SetArgs([]string{"active-track", msg})
	err := command.Execute()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Verify File Content
	content, _ := fs.ReadFile(logPath)
	strContent := string(content)

	if !contains(strContent, initialContent) {
		t.Errorf("expected initial content preserved")
	}
	if !contains(strContent, msg) {
		t.Errorf("expected new message '%s' in log", msg)
	}

	// Verify output
	if !contains(buf.String(), fmt.Sprintf("Logged to %s", logPath)) {
		t.Errorf("expected success message")
	}
}
