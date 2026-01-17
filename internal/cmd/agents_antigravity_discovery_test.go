package cmd

import (
	"path/filepath"
	"testing"

	"cdd/internal/platform"
)

func TestAntigravityWorkflowsAreDiscoverable(t *testing.T) {
	fs := platform.NewMockFileSystem()
	cmd := NewAgentsCmd(fs)

	// Install skills (as workflows for antigravity)
	cmd.SetArgs([]string{"--install", "--target", "antigravity"})
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("installation failed: %v", err)
	}

	// Simulate Antigravity discovery by checking that all workflow files exist and have valid content
	workflowIDs := []string{"cdd", "cdd-analyst", "cdd-architect", "cdd-executor", "cdd-integrator"}

	for _, workflowID := range workflowIDs {
		workflowFile := filepath.Join(".agent", "workflows", workflowID+".md")

		// Check file exists
		_, err := fs.Stat(workflowFile)
		if err != nil {
			t.Errorf("workflow file %s does not exist: %v", workflowFile, err)
			continue
		}

		// Check content exists
		content, err := fs.ReadFile(workflowFile)
		if err != nil {
			t.Errorf("file not found in %s: %v", workflowFile, err)
			continue
		}

		// Verify it has valid Antigravity format
		err = validateAntigravityWorkflow(string(content))
		if err != nil {
			t.Errorf("workflow %s failed Antigravity validation: %v", workflowID, err)
		}

		// Verify it contains YAML frontmatter with name and description
		contentStr := string(content)
		if !hasYAMLField(contentStr, "name") {
			t.Errorf("workflow %s missing 'name' in YAML frontmatter", workflowID)
		}
		if !hasYAMLField(contentStr, "description") {
			t.Errorf("workflow %s missing 'description' in YAML frontmatter", workflowID)
		}
	}
}

func hasYAMLField(content, field string) bool {
	// Simple check for field: in YAML frontmatter
	start := 0
	if content[0:3] == "---" {
		start = 3
	}
	end := len(content)
	if idx := findString(content[start:], "---"); idx != -1 {
		end = start + idx
	}
	frontmatter := content[start:end]
	return findString(frontmatter, field+":") != -1
}

func findString(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
