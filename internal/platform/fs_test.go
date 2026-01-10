package platform_test

import (
	"testing"

	"cdd/internal/platform"
)

func TestMockFileSystem_WriteFile(t *testing.T) {
	fs := platform.NewMockFileSystem()
	err := fs.WriteFile("test.txt", []byte("hello"), 0644)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	content, err := fs.ReadFile("test.txt")
	if err != nil {
		t.Fatalf("expected no error reading file, got %v", err)
	}
	if string(content) != "hello" {
		t.Errorf("expected 'hello', got '%s'", string(content))
	}
}
