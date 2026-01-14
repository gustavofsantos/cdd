package cmd

import (
	"strings"
	"testing"

	"cdd/internal/platform"
)

func TestGetActiveTasks(t *testing.T) {
	fs := platform.NewMockFileSystem()

	// Setup mock filesystem with multiple active tracks
	// Note: Must write files for directories to appear in ReadDir
	_ = fs.WriteFile(".context/tracks/task-1/plan.md", []byte(""), 0644)
	_ = fs.WriteFile(".context/tracks/task-2/plan.md", []byte(""), 0644)
	_ = fs.WriteFile(".context/tracks/task-3/plan.md", []byte(""), 0644)

	t.Run("returns all active tasks", func(t *testing.T) {
		tasks, err := GetActiveTasks(fs)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(tasks) != 3 {
			t.Errorf("expected 3 tasks, got %d", len(tasks))
		}

		expected := map[string]bool{"task-1": true, "task-2": true, "task-3": true}
		for _, task := range tasks {
			if !expected[task] {
				t.Errorf("unexpected task: %s", task)
			}
		}
	})

	t.Run("returns empty list when no tasks exist", func(t *testing.T) {
		fs2 := platform.NewMockFileSystem()
		tasks, err := GetActiveTasks(fs2)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(tasks) != 0 {
			t.Errorf("expected 0 tasks, got %d", len(tasks))
		}
	})

	t.Run("returns single task", func(t *testing.T) {
		fs3 := platform.NewMockFileSystem()
		_ = fs3.WriteFile(".context/tracks/single-task/plan.md", []byte(""), 0644)

		tasks, err := GetActiveTasks(fs3)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(tasks) != 1 {
			t.Errorf("expected 1 task, got %d", len(tasks))
		}

		if tasks[0] != "single-task" {
			t.Errorf("expected 'single-task', got %s", tasks[0])
		}
	})
}

func TestGetTaskCompletionSuggestion(t *testing.T) {
	t.Run("returns single task name when exactly one task is active", func(t *testing.T) {
		fs := platform.NewMockFileSystem()
		_ = fs.WriteFile(".context/tracks/my-task/plan.md", []byte(""), 0644)

		suggestion, count, err := GetTaskCompletionSuggestion(fs)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if suggestion != "my-task" {
			t.Errorf("expected 'my-task', got %s", suggestion)
		}

		if count != 1 {
			t.Errorf("expected count 1, got %d", count)
		}
	})

	t.Run("returns empty string when multiple tasks are active", func(t *testing.T) {
		fs := platform.NewMockFileSystem()
		_ = fs.WriteFile(".context/tracks/task-1/plan.md", []byte(""), 0644)
		_ = fs.WriteFile(".context/tracks/task-2/plan.md", []byte(""), 0644)

		suggestion, count, err := GetTaskCompletionSuggestion(fs)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if suggestion != "" {
			t.Errorf("expected empty string, got %s", suggestion)
		}

		if count != 2 {
			t.Errorf("expected count 2, got %d", count)
		}
	})

	t.Run("returns empty string when no tasks are active", func(t *testing.T) {
		fs := platform.NewMockFileSystem()

		suggestion, count, err := GetTaskCompletionSuggestion(fs)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if suggestion != "" {
			t.Errorf("expected empty string, got %s", suggestion)
		}

		if count != 0 {
			t.Errorf("expected count 0, got %d", count)
		}
	})
}

func TestRenderTaskSelectionMenu(t *testing.T) {
	t.Run("renders menu for multiple tasks", func(t *testing.T) {
		tasks := []string{"task-1", "task-2", "task-3"}
		menu := RenderTaskSelectionMenu(tasks)

		// Verify menu contains all task names
		for _, task := range tasks {
			if !strings.Contains(menu, task) {
				t.Errorf("expected menu to contain task '%s'", task)
			}
		}

		// Verify menu is not empty
		if menu == "" {
			t.Errorf("expected non-empty menu")
		}
	})

	t.Run("formats menu with indices", func(t *testing.T) {
		tasks := []string{"task-1", "task-2"}
		menu := RenderTaskSelectionMenu(tasks)

		// Should contain "1)" and "2)" for numbering
		if !strings.Contains(menu, "1)") {
			t.Errorf("expected menu to contain '1)' for numbering")
		}
		if !strings.Contains(menu, "2)") {
			t.Errorf("expected menu to contain '2)' for numbering")
		}
	})

	t.Run("handles single task in menu", func(t *testing.T) {
		tasks := []string{"single-task"}
		menu := RenderTaskSelectionMenu(tasks)

		if !strings.Contains(menu, "single-task") {
			t.Errorf("expected menu to contain 'single-task'")
		}
	})
}

func TestHandleTaskSelection(t *testing.T) {
	t.Run("returns task at valid index", func(t *testing.T) {
		tasks := []string{"task-1", "task-2", "task-3"}
		selected, err := HandleTaskSelection(tasks, "1")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if selected != "task-1" {
			t.Errorf("expected 'task-1', got '%s'", selected)
		}
	})

	t.Run("returns task at second index", func(t *testing.T) {
		tasks := []string{"task-1", "task-2", "task-3"}
		selected, err := HandleTaskSelection(tasks, "2")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if selected != "task-2" {
			t.Errorf("expected 'task-2', got '%s'", selected)
		}
	})

	t.Run("returns error for invalid index", func(t *testing.T) {
		tasks := []string{"task-1", "task-2"}
		_, err := HandleTaskSelection(tasks, "5")
		if err == nil {
			t.Errorf("expected error for invalid index")
		}
	})

	t.Run("returns error for zero index", func(t *testing.T) {
		tasks := []string{"task-1", "task-2"}
		_, err := HandleTaskSelection(tasks, "0")
		if err == nil {
			t.Errorf("expected error for zero index")
		}
	})

	t.Run("returns error for non-numeric input", func(t *testing.T) {
		tasks := []string{"task-1", "task-2"}
		_, err := HandleTaskSelection(tasks, "abc")
		if err == nil {
			t.Errorf("expected error for non-numeric input")
		}
	})

	t.Run("returns error for empty input", func(t *testing.T) {
		tasks := []string{"task-1", "task-2"}
		_, err := HandleTaskSelection(tasks, "")
		if err == nil {
			t.Errorf("expected error for empty input")
		}
	})
}
