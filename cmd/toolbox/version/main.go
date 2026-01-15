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
		"name":        "cdd-version",
		"description": "Display the version of the cdd CLI",
		"args":        map[string][2]string{},
	}

	output, _ := json.Marshal(desc)
	fmt.Println(string(output))
}

func execute() {
	// Read stdin (not used for version, but consume it)
	_, _ = io.ReadAll(os.Stdin)

	// Find the cdd binary
	cddBin := "cdd"
	if _, err := exec.LookPath(cddBin); err != nil {
		fmt.Fprintf(os.Stderr, "cdd binary not found in PATH\n")
		os.Exit(1)
	}

	// Execute: cdd version
	cmd := exec.Command(cddBin, "version")
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
