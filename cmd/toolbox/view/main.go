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
		"name":        "cdd-view",
		"description": "Display track details or dashboard",
		"args": map[string][2]string{
			"track_name": {"string", "name of the track (optional)"},
			"spec":       {"bool", "show track specification"},
			"plan":       {"bool", "show track plan"},
			"log":        {"bool", "show decision log"},
			"raw":        {"bool", "output raw text (pipe-friendly)"},
			"archived":   {"bool", "view archived tracks"},
			"inbox":      {"bool", "view inbox"},
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

	// Build command: cdd view [track-name] [flags]
	cmdArgs := []string{"view"}

	// Add positional argument if present
	if trackName, ok := args["track_name"]; ok && trackName != nil && trackName != "" {
		cmdArgs = append(cmdArgs, fmt.Sprint(trackName))
	}

	// Add flags
	flags := []string{"spec", "plan", "log", "raw", "archived", "inbox"}
	for _, flag := range flags {
		if val, ok := args[flag]; ok {
			if bval, ok := val.(bool); ok && bval {
				cmdArgs = append(cmdArgs, "--"+flag)
			}
		}
	}

	// Find the cdd binary
	cddBin := "cdd"
	if _, err := exec.LookPath(cddBin); err != nil {
		fmt.Fprintf(os.Stderr, "cdd binary not found in PATH\n")
		os.Exit(1)
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
