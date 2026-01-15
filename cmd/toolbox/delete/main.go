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
		"name":        "cdd-delete",
		"description": "Destructively remove an active track",
		"args": map[string][2]string{
			"track_name": {"string", "name of the track to delete (required)"},
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

	// Extract track name
	trackName := ""
	if name, ok := args["track_name"]; ok {
		trackName = fmt.Sprint(name)
	}

	if trackName == "" {
		fmt.Fprintf(os.Stderr, "Error: track_name is required\n")
		os.Exit(1)
	}

	// Find the cdd binary
	cddBin := "cdd"
	if _, err := exec.LookPath(cddBin); err != nil {
		fmt.Fprintf(os.Stderr, "cdd binary not found in PATH\n")
		os.Exit(1)
	}

	// Execute: cdd delete <track-name>
	cmd := exec.Command(cddBin, "delete", trackName)
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
