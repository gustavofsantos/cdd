package cmd

import (
	"path/filepath"
	"testing"

	"cdd/internal/platform"
)

func TestDiscoverSpecsMultipleFiles(t *testing.T) {
	// Test discovering multiple spec files
	fs := platform.NewRealFileSystem()
	specsPath := filepath.Join("..", "..", ".context", "specs")

	specs, err := DiscoverSpecs(fs, specsPath)
	if err != nil {
		t.Fatalf("DiscoverSpecs() returned error: %v", err)
	}

	// Should have at least 3 spec files
	if len(specs) < 3 {
		t.Errorf("Expected at least 3 specs, got %d", len(specs))
	}

	// Verify each spec has required fields
	for i, spec := range specs {
		if spec.Name == "" {
			t.Errorf("Spec %d has empty Name", i)
		}
		if spec.Path == "" {
			t.Errorf("Spec %d has empty Path", i)
		}
		if len(spec.Content) == 0 {
			t.Errorf("Spec %d (%s) has empty Content", i, spec.Name)
		}
		if !platform.EndsWithString(spec.Name, ".md") {
			t.Errorf("Spec %d (%s) is not a markdown file", i, spec.Name)
		}
	}
}

func TestDiscoverSpecsFileOrder(t *testing.T) {
	// Test that all discovered specs are consistent
	fs := platform.NewRealFileSystem()
	specsPath := filepath.Join("..", "..", ".context", "specs")

	specs1, err := DiscoverSpecs(fs, specsPath)
	if err != nil {
		t.Fatalf("First DiscoverSpecs() returned error: %v", err)
	}

	specs2, err := DiscoverSpecs(fs, specsPath)
	if err != nil {
		t.Fatalf("Second DiscoverSpecs() returned error: %v", err)
	}

	// Same count
	if len(specs1) != len(specs2) {
		t.Errorf("Inconsistent spec count: %d vs %d", len(specs1), len(specs2))
	}

	// All names should match (order may differ on different filesystems)
	names1 := make(map[string]bool)
	for _, spec := range specs1 {
		names1[spec.Name] = true
	}

	for _, spec := range specs2 {
		if !names1[spec.Name] {
			t.Errorf("Spec %s missing in second discovery", spec.Name)
		}
	}
}

func TestDiscoverSpecsContentLoaded(t *testing.T) {
	// Test that all spec contents are loaded and non-empty
	fs := platform.NewRealFileSystem()
	specsPath := filepath.Join("..", "..", ".context", "specs")

	specs, err := DiscoverSpecs(fs, specsPath)
	if err != nil {
		t.Fatalf("DiscoverSpecs() returned error: %v", err)
	}

	for _, spec := range specs {
		if len(spec.Content) == 0 {
			t.Errorf("Spec %s has no content", spec.Name)
		}

		// Verify content is valid markdown (at least starts with some text)
		content := string(spec.Content)
		if len(content) < 10 {
			t.Errorf("Spec %s content too short: %d bytes", spec.Name, len(content))
		}

		// Should be readable as UTF-8
		if !isValidUTF8(spec.Content) {
			t.Errorf("Spec %s content is not valid UTF-8", spec.Name)
		}
	}
}

func TestDiscoverSpecsDistinctNames(t *testing.T) {
	// Test that all discovered specs have unique names
	fs := platform.NewRealFileSystem()
	specsPath := filepath.Join("..", "..", ".context", "specs")

	specs, err := DiscoverSpecs(fs, specsPath)
	if err != nil {
		t.Fatalf("DiscoverSpecs() returned error: %v", err)
	}

	names := make(map[string]int)
	for _, spec := range specs {
		names[spec.Name]++
	}

	for name, count := range names {
		if count > 1 {
			t.Errorf("Spec file %s discovered %d times", name, count)
		}
	}
}

func TestDiscoverSpecsPathConsistency(t *testing.T) {
	// Test that spec paths are properly constructed
	fs := platform.NewRealFileSystem()
	specsPath := filepath.Join("..", "..", ".context", "specs")

	specs, err := DiscoverSpecs(fs, specsPath)
	if err != nil {
		t.Fatalf("DiscoverSpecs() returned error: %v", err)
	}

	for _, spec := range specs {
		// Path should contain the spec name
		if !platform.EndsWithString(spec.Path, spec.Name) {
			t.Errorf("Path %s doesn't end with name %s", spec.Path, spec.Name)
		}

		// Path should be readable
		content, err := fs.ReadFile(spec.Path)
		if err != nil {
			t.Errorf("Cannot read spec file at path %s: %v", spec.Path, err)
		}

		// Content should match what was returned
		if string(content) != string(spec.Content) {
			t.Errorf("Spec %s content mismatch between discovery and file read", spec.Name)
		}
	}
}

// Helper function to check if bytes are valid UTF-8
func isValidUTF8(data []byte) bool {
	for i := 0; i < len(data); {
		r, size := utf8DecodeRune(data[i:])
		if r == utf8RuneError && size == 1 {
			return false
		}
		i += size
	}
	return true
}

const (
	utf8RuneError = '\uFFFD'
)

// Simplified UTF-8 rune decoder
func utf8DecodeRune(p []byte) (rune, int) {
	if len(p) == 0 {
		return utf8RuneError, 0
	}
	b := p[0]
	if b < 0x80 {
		return rune(b), 1
	}
	if b < 0xC0 {
		return utf8RuneError, 1
	}
	if b < 0xE0 {
		if len(p) < 2 {
			return utf8RuneError, 1
		}
		if p[1]&0xC0 != 0x80 {
			return utf8RuneError, 1
		}
		return rune(b&0x1F)<<6 | rune(p[1]&0x3F), 2
	}
	if b < 0xF0 {
		if len(p) < 3 {
			return utf8RuneError, 1
		}
		if p[1]&0xC0 != 0x80 || p[2]&0xC0 != 0x80 {
			return utf8RuneError, 1
		}
		return rune(b&0x0F)<<12 | rune(p[1]&0x3F)<<6 | rune(p[2]&0x3F), 3
	}
	if len(p) < 4 {
		return utf8RuneError, 1
	}
	if p[1]&0xC0 != 0x80 || p[2]&0xC0 != 0x80 || p[3]&0xC0 != 0x80 {
		return utf8RuneError, 1
	}
	return rune(b&0x07)<<18 | rune(p[1]&0x3F)<<12 | rune(p[2]&0x3F)<<6 | rune(p[3]&0x3F), 4
}
