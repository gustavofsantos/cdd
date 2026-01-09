package platform

import (
	"bytes"
	"os"
	"strings"
	"time"
)

type MockFileSystem struct {
	Files map[string][]byte
}

func NewMockFileSystem() *MockFileSystem {
	return &MockFileSystem{
		Files: make(map[string][]byte),
	}
}

func (m *MockFileSystem) Stat(name string) (os.FileInfo, error) {
	// Check if it's a file
	if val, ok := m.Files[name]; ok {
		return &MockFileInfo{name: name, size: int64(len(val)), isDir: false}, nil
	}
	// Check if it's a directory (by prefix)
	// Simple heuristic: if any key starts with name + "/", it's a dir
	prefix := name
	if !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}
	for k := range m.Files {
		if strings.HasPrefix(k, prefix) {
			return &MockFileInfo{name: name, isDir: true}, nil
		}
	}
	return nil, os.ErrNotExist
}

func (m *MockFileSystem) MkdirAll(path string, perm os.FileMode) error {
	// We verify checks usually via Stat. We can perform a "touch" of a directory placeholder
	// strictly speaking, but our prefix logic in Stat handles implicit dirs.
	return nil
}

func (m *MockFileSystem) WriteFile(name string, data []byte, perm os.FileMode) error {
	m.Files[name] = data
	return nil
}

func (m *MockFileSystem) ReadFile(name string) ([]byte, error) {
	if data, ok := m.Files[name]; ok {
		return data, nil
	}
	return nil, os.ErrNotExist
}

func (m *MockFileSystem) OpenFile(name string, flag int, perm os.FileMode) (File, error) {
	var content []byte
	if val, ok := m.Files[name]; ok {
		content = val
	} else if flag&os.O_CREATE == 0 {
		return nil, os.ErrNotExist
	}

	if flag&os.O_TRUNC != 0 {
		content = []byte{}
	}

	// In a real file, O_APPEND means writes happen at end.
	// We will simulate simply by loading content.
	return &MockFile{name: name, fs: m, Buffer: bytes.NewBuffer(content)}, nil
}

func (m *MockFileSystem) Create(name string) (File, error) {
	m.Files[name] = []byte{}
	return &MockFile{name: name, fs: m, Buffer: bytes.NewBuffer([]byte{})}, nil
}

func (m *MockFileSystem) Rename(oldpath, newpath string) error {
	// 1. Check direct file
	if data, ok := m.Files[oldpath]; ok {
		m.Files[newpath] = data
		delete(m.Files, oldpath)
		return nil
	}

	// 2. Check directory (prefix)
	found := false
	oldPrefix := oldpath
	if !strings.HasSuffix(oldPrefix, "/") {
		oldPrefix += "/"
	}
	newPrefix := newpath
	if !strings.HasSuffix(newPrefix, "/") {
		newPrefix += "/"
	}

	var keysToMove []string
	for k := range m.Files {
		if strings.HasPrefix(k, oldPrefix) {
			keysToMove = append(keysToMove, k)
		}
	}

	if len(keysToMove) > 0 {
		found = true
		for _, k := range keysToMove {
			suffix := strings.TrimPrefix(k, oldPrefix)
			newKey := newPrefix + suffix
			m.Files[newKey] = m.Files[k]
			delete(m.Files, k)
		}
	}

	if found {
		return nil
	}

	return os.ErrNotExist
}

func (m *MockFileSystem) Remove(name string) error {
	delete(m.Files, name)
	return nil
}

func (m *MockFileSystem) ReadDir(name string) ([]os.DirEntry, error) {
	entries := []os.DirEntry{}
	seen := make(map[string]bool)

	prefix := name
	if !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}

	for k := range m.Files {
		if strings.HasPrefix(k, prefix) {
			// Extract direct child
			rel := strings.TrimPrefix(k, prefix)
			parts := strings.Split(rel, "/")
			child := parts[0]
			if !seen[child] {
				seen[child] = true
				isDir := len(parts) > 1
				entries = append(entries, &MockDirEntry{name: child, isDir: isDir})
			}
		}
	}
	return entries, nil
}

// MockFileInfo
type MockFileInfo struct {
	name  string
	size  int64
	isDir bool
}

func (m *MockFileInfo) Name() string       { return m.name }
func (m *MockFileInfo) Size() int64        { return m.size }
func (m *MockFileInfo) Mode() os.FileMode  { return 0644 }
func (m *MockFileInfo) ModTime() time.Time { return time.Now() }
func (m *MockFileInfo) IsDir() bool        { return m.isDir }
func (m *MockFileInfo) Sys() any           { return nil }

// MockDirEntry
type MockDirEntry struct {
	name  string
	isDir bool
}

func (m *MockDirEntry) Name() string      { return m.name }
func (m *MockDirEntry) IsDir() bool       { return m.isDir }
func (m *MockDirEntry) Type() os.FileMode { return 0644 }
func (m *MockDirEntry) Info() (os.FileInfo, error) {
	return &MockFileInfo{name: m.name, isDir: m.isDir}, nil
}

// MockFile
type MockFile struct {
	name string
	fs   *MockFileSystem
	*bytes.Buffer
}

func (m *MockFile) Close() error {
	if m.fs != nil {
		m.fs.Files[m.name] = m.Buffer.Bytes()
	}
	return nil
}
func (m *MockFile) WriteString(s string) (n int, err error) {
	return m.Buffer.WriteString(s)
}
