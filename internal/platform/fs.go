package platform

import (
	"io"
	"os"
)

// File represents an open file.
type File interface {
	io.Reader
	io.Writer
	io.Closer
	WriteString(s string) (n int, err error)
}

// FileSystem abstracts OS file operations.
type FileSystem interface {
	Stat(name string) (os.FileInfo, error)
	MkdirAll(path string, perm os.FileMode) error
	WriteFile(name string, data []byte, perm os.FileMode) error
	ReadFile(name string) ([]byte, error)
	OpenFile(name string, flag int, perm os.FileMode) (File, error)
	Rename(oldpath, newpath string) error
	ReadDir(name string) ([]os.DirEntry, error)
	Create(name string) (File, error)
	Remove(name string) error
	RemoveAll(path string) error
}

// RealFileSystem implements FileSystem using the os package.
type RealFileSystem struct{}

func NewRealFileSystem() *RealFileSystem {
	return &RealFileSystem{}
}

func (fs *RealFileSystem) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

func (fs *RealFileSystem) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func (fs *RealFileSystem) WriteFile(name string, data []byte, perm os.FileMode) error {
	return os.WriteFile(name, data, perm)
}

func (fs *RealFileSystem) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func (fs *RealFileSystem) OpenFile(name string, flag int, perm os.FileMode) (File, error) {
	return os.OpenFile(name, flag, perm)
}

func (fs *RealFileSystem) Rename(oldpath, newpath string) error {
	return os.Rename(oldpath, newpath)
}

func (fs *RealFileSystem) ReadDir(name string) ([]os.DirEntry, error) {
	return os.ReadDir(name)
}

func (fs *RealFileSystem) Create(name string) (File, error) {
	return os.Create(name)
}

func (fs *RealFileSystem) Remove(name string) error {
	return os.Remove(name)
}

func (fs *RealFileSystem) RemoveAll(path string) error {
	return os.RemoveAll(path)
}
