package cmd_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"cdd/internal/cmd"
	"cdd/internal/platform"
)

func TestDumpCmd_AppendsStdinToScratchpad(t *testing.T) {
	fs := platform.NewMockFileSystem()
	trackName := "current-track"
	scratchFile := ".context/tracks/" + trackName + "/scratchpad.md"
	initial := "# Scratchpad\n"
	_ = fs.WriteFile(scratchFile, []byte(initial), 0644)

	command := cmd.NewDumpCmd(fs)
	bufOut := new(bytes.Buffer)
	command.SetOut(bufOut)
	command.SetErr(bufOut)

	// Simulate Stdin
	input := "Log Data 123"
	command.SetIn(bytes.NewBufferString(input))

	command.SetArgs([]string{trackName})
	err := command.Execute()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Verify content
	content, _ := fs.ReadFile(scratchFile)
	if !strings.Contains(string(content), initial) {
		t.Errorf("expected initial content")
	}
	if !strings.Contains(string(content), input) {
		t.Errorf("expected dumped content '%s'", input)
	}

	// Verify output message
	expected := fmt.Sprintf(">> Data appended to %s", scratchFile)
	if !strings.Contains(bufOut.String(), expected) {
		t.Errorf("expected success message")
	}
}
