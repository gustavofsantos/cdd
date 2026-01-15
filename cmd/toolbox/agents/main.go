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
		"name":        "cdd-agents",
		"description": "Manage AI agent integration and skills",
		"args": map[string][2]string{
			"install": {"bool", "install CDD skills"},
			"target":  {"string", "installation target (agent, antigravity, claude, agents, cursor)"},
			"all":     {"bool", "install all skills"},
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

	// Build command: cdd agents [flags]
	cmdArgs := []string{"agents"}

	// Check for install flag
	if install, ok := args["install"]; ok {
		if bval, ok := install.(bool); ok && bval {
			cmdArgs = append(cmdArgs, "--install")
		}
	}

	// Check for target flag
	if target, ok := args["target"]; ok && target != nil && target != "" {
		cmdArgs = append(cmdArgs, "--target", fmt.Sprint(target))
	}

	// Check for all flag
	if all, ok := args["all"]; ok {
		if bval, ok := all.(bool); ok && bval {
			cmdArgs = append(cmdArgs, "--all")
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
