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
		"name":        "cdd-recite",
		"description": "Output the current Plan to the context window",
		"args": map[string][2]string{
			"track_name": {"string", "name of the track (optional, inferred if only one exists)"},
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

	// Extract track name (optional)
	trackName := ""
	if name, ok := args["track_name"]; ok && name != nil && name != "" {
		trackName = fmt.Sprint(name)
	}

	// Find the cdd binary
	cddBin := "cdd"
	if _, err := exec.LookPath(cddBin); err != nil {
		fmt.Fprintf(os.Stderr, "cdd binary not found in PATH\n")
		os.Exit(1)
	}

	// Build command: cdd recite [track-name]
	cmdArgs := []string{"recite"}
	if trackName != "" {
		cmdArgs = append(cmdArgs, trackName)
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
