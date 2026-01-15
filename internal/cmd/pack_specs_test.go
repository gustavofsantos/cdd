package cmd

import (
	"path/filepath"
	"testing"

	"cdd/internal/platform"
)

func TestDiscoverSpecs(t *testing.T) {
	// Use real file system to test discovery
	fs := platform.NewRealFileSystem()

	// Adjust path for running from test location
	specsPath := filepath.Join("..", "..", ".context", "specs")
	specs, err := DiscoverSpecs(fs, specsPath)
	if err != nil {
		t.Fatalf("DiscoverSpecs() returned error: %v", err)
	}

	if len(specs) == 0 {
		t.Fatalf("DiscoverSpecs() returned no specs, expected at least one")
	}

	// Verify we got markdown files
	for _, spec := range specs {
		if !platform.EndsWithString(spec.Path, ".md") {
			t.Errorf("Expected .md file, got %s", spec.Path)
		}
		if spec.Name == "" {
			t.Errorf("Spec name should not be empty")
		}
		if spec.Content == nil || len(spec.Content) == 0 {
			t.Errorf("Spec content should not be empty for %s", spec.Name)
		}
	}
}

func TestDiscoverSpecsWithSpecificFiles(t *testing.T) {
	// Verify specific files are discovered
	fs := platform.NewRealFileSystem()

	specsPath := filepath.Join("..", "..", ".context", "specs")
	specs, err := DiscoverSpecs(fs, specsPath)
	if err != nil {
		t.Fatalf("DiscoverSpecs() returned error: %v", err)
	}

	specNames := make(map[string]bool)
	for _, spec := range specs {
		specNames[spec.Name] = true
	}

	// These files should exist based on the territory mapping
	expectedFiles := []string{"log.md", "view.md", "help.md"}
	for _, expectedFile := range expectedFiles {
		if !specNames[expectedFile] {
			t.Errorf("Expected file %s not found in specs", expectedFile)
		}
	}
}

func TestDiscoverSpecsFileNotFound(t *testing.T) {
	fs := platform.NewRealFileSystem()

	specs, err := DiscoverSpecs(fs, ".context/nonexistent")
	if err == nil {
		t.Errorf("DiscoverSpecs() should return error for nonexistent directory")
	}
	if specs != nil && len(specs) > 0 {
		t.Errorf("DiscoverSpecs() should return empty specs for nonexistent directory")
	}
}

func TestSpecStruct(t *testing.T) {
	// Test that Spec struct is properly formed
	spec := Spec{
		Name:    "test.md",
		Path:    filepath.Join(".context", "specs", "test.md"),
		Content: []byte("# Test Content"),
	}

	if spec.Name != "test.md" {
		t.Errorf("Spec.Name should be 'test.md', got %s", spec.Name)
	}
	if len(spec.Content) == 0 {
		t.Errorf("Spec.Content should not be empty")
	}
}
