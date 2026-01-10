package cmd

import (
	"bytes"
	"strings"
	"testing"

	"cdd/internal/platform"

	"github.com/spf13/cobra"
)

func TestCheckInboxSize_Suggestion(t *testing.T) {
	fs := platform.NewMockFileSystem()
	inboxFile := ".context/inbox.md"

	// Create a file with 51 lines
	largeContent := strings.Repeat("line\n", 50) + "last line"
	_ = fs.WriteFile(inboxFile, []byte(largeContent), 0644)

	cmd := &cobra.Command{}
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)

	CheckInboxSize(fs, cmd)

	output := buf.String()
	if !strings.Contains(output, "getting large") {
		t.Errorf("expected suggestion for large inbox, got: %s", output)
	}
	if !strings.Contains(output, "51 lines") {
		t.Errorf("expected count 51, got: %s", output)
	}
}

func TestCheckInboxSize_NoSuggestion(t *testing.T) {
	fs := platform.NewMockFileSystem()
	inboxFile := ".context/inbox.md"

	// Create a file with 10 lines
	smallContent := strings.Repeat("line\n", 9) + "last line"
	_ = fs.WriteFile(inboxFile, []byte(smallContent), 0644)

	cmd := &cobra.Command{}
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)

	CheckInboxSize(fs, cmd)

	output := buf.String()
	if output != "" {
		t.Errorf("expected no suggestion for small inbox, got: %s", output)
	}
}
