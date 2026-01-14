package cmd_test

import (
	"bytes"
	"fmt"
	"strings"
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

func TestLogCmd_InputsFromStdin_WithTrackName(t *testing.T) {
	fs := platform.NewMockFileSystem()
	logPath := ".context/tracks/active-track/decisions.md"
	_ = fs.MkdirAll(".context/tracks/active-track", 0755)

	command := cmd.NewLogCmd(fs)
	buf := new(bytes.Buffer)
	inBuf := bytes.NewBufferString("Log message from stdin")
	command.SetOut(buf)
	command.SetErr(buf)
	command.SetIn(inBuf)

	command.SetArgs([]string{"active-track"})
	err := command.Execute()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	content, _ := fs.ReadFile(logPath)
	if !strings.Contains(string(content), "Log message from stdin") {
		t.Errorf("expected new message 'Log message from stdin' in log")
	}
}

func TestLogCmd_InputsFromStdin_InferTrackName(t *testing.T) {
	fs := platform.NewMockFileSystem()
	_ = fs.MkdirAll(".context/tracks/inferred-track", 0755)
	// MockFS requires a file to exist to list the directory
	_ = fs.WriteFile(".context/tracks/inferred-track/.keep", []byte{}, 0644)
	logPath := ".context/tracks/inferred-track/decisions.md"

	command := cmd.NewLogCmd(fs)
	buf := new(bytes.Buffer)
	inBuf := bytes.NewBufferString("Inferred message")
	command.SetOut(buf)
	command.SetErr(buf)
	command.SetIn(inBuf)

	command.SetArgs([]string{}) // 0 args
	err := command.Execute()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	content, _ := fs.ReadFile(logPath)
	if !strings.Contains(string(content), "Inferred message") {
		t.Errorf("expected logs to contain inferred message")
	}
}

func TestLogCmd_InputsFromStdin_InferTrackName_MultipleTracksError(t *testing.T) {
	fs := platform.NewMockFileSystem()
	_ = fs.MkdirAll(".context/tracks/track-1", 0755)
	_ = fs.WriteFile(".context/tracks/track-1/.keep", []byte{}, 0644)
	_ = fs.MkdirAll(".context/tracks/track-2", 0755)
	_ = fs.WriteFile(".context/tracks/track-2/.keep", []byte{}, 0644)

	command := cmd.NewLogCmd(fs)
	inBuf := bytes.NewBufferString("Message")
	command.SetIn(inBuf)
	command.SetArgs([]string{})

	err := command.Execute()
	if err == nil {
		t.Fatal("expected error due to multiple tracks, got nil")
	}
	if !strings.Contains(err.Error(), "multiple active tracks found") {
		t.Errorf("expected multiple tracks error, got %v", err)
	}
}
