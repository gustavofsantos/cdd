package cmd

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"cdd/internal/platform"
	"github.com/spf13/cobra"
)

// GetActiveTasks returns a sorted list of all active task names
func GetActiveTasks(fs platform.FileSystem) ([]string, error) {
	entries, err := fs.ReadDir(".context/tracks")
	if err != nil {
		// Return empty list if directory doesn't exist
		return []string{}, nil
	}

	var tasks []string
	for _, entry := range entries {
		if entry.IsDir() {
			tasks = append(tasks, entry.Name())
		}
	}

	sort.Strings(tasks)
	return tasks, nil
}

// GetTaskCompletionSuggestion returns a suggestion for autocompletion and the count of active tasks
// - If exactly one task is active, returns that task name and count=1
// - If zero or multiple tasks are active, returns empty string and the count
func GetTaskCompletionSuggestion(fs platform.FileSystem) (string, int, error) {
	tasks, err := GetActiveTasks(fs)
	if err != nil {
		return "", 0, err
	}

	count := len(tasks)
	if count == 1 {
		return tasks[0], count, nil
	}

	return "", count, nil
}

// getViewCompletion provides completion suggestions for the view command
func getViewCompletion(fs platform.FileSystem, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	tasks, err := GetActiveTasks(fs)
	if err != nil {
		return []string{}, cobra.ShellCompDirectiveNoFileComp
	}

	// Filter tasks that start with toComplete
	var filtered []string
	for _, task := range tasks {
		if len(toComplete) == 0 || (len(task) >= len(toComplete) && task[:len(toComplete)] == toComplete) {
			filtered = append(filtered, task)
		}
	}

	return filtered, cobra.ShellCompDirectiveNoFileComp
}

// RenderTaskSelectionMenu renders a formatted menu for selecting from multiple tasks
func RenderTaskSelectionMenu(tasks []string) string {
	if len(tasks) == 0 {
		return ""
	}

	var menu strings.Builder
	menu.WriteString("\nSelect a task:\n\n")

	for i, task := range tasks {
		menu.WriteString(fmt.Sprintf("%d) %s\n", i+1, task))
	}

	return menu.String()
}

// HandleTaskSelection processes user input and returns the selected task name
// Input should be 1-indexed (e.g., "1" for first task)
func HandleTaskSelection(tasks []string, input string) (string, error) {
	if input == "" {
		return "", fmt.Errorf("no input provided")
	}

	index, err := strconv.Atoi(input)
	if err != nil {
		return "", fmt.Errorf("invalid input: %s is not a number", input)
	}

	// Convert 1-indexed to 0-indexed
	if index <= 0 || index > len(tasks) {
		return "", fmt.Errorf("invalid selection: %d is out of range (1-%d)", index, len(tasks))
	}

	return tasks[index-1], nil
}
