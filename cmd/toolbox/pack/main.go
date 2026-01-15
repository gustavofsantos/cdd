package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	action := os.Getenv("TOOLBOX_ACTION")

	switch action {
	case "describe":
		describe()
	case "execute":
		execute()
	default:
		fmt.Fprintf(os.Stderr, "invalid TOOLBOX_ACTION: %s\n", action)
		os.Exit(1)
	}
}

func describe() {
	desc := map[string]interface{}{
		"name":        "cdd-pack",
		"description": "Extract and compress relevant specs by topic focus. Searches all specifications and returns only paragraphs matching your topic, ranked by relevance.",
		"args": map[string][2]string{
			"focus": {"string", "topic to search for (required). E.g., 'log', 'command', 'authentication'"},
			"raw":   {"boolean", "output plain text without markdown rendering (optional)"},
		},
	}

	output, _ := json.Marshal(desc)
	fmt.Println(string(output))
}

func execute() {
	// Read arguments from stdin
	input, _ := io.ReadAll(os.Stdin)

	var args map[string]interface{}
	if len(input) > 0 {
		_ = json.Unmarshal(input, &args)
	}

	// Extract arguments
	focus := ""
	if f, ok := args["focus"]; ok {
		focus = fmt.Sprint(f)
	}

	raw := false
	if r, ok := args["raw"]; ok {
		if b, ok := r.(bool); ok {
			raw = b
		} else if s, ok := r.(string); ok {
			raw = s == "true" || s == "True" || s == "TRUE" || s == "1"
		}
	}

	if focus == "" {
		fmt.Fprintf(os.Stderr, "Error: focus argument is required\n")
		os.Exit(1)
	}

	// Find the cdd binary
	cddBin := "cdd"
	if _, err := exec.LookPath(cddBin); err != nil {
		fmt.Fprintf(os.Stderr, "cdd binary not found in PATH\n")
		os.Exit(1)
	}

	// Build command: cdd pack --focus <topic> [--raw]
	cmdArgs := []string{"pack", "--focus", focus}
	if raw {
		cmdArgs = append(cmdArgs, "--raw")
	}

	cmd := exec.Command(cddBin, cmdArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			os.Exit(exitErr.ExitCode())
		}
		os.Exit(1)
	}
}
