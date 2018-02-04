package filesystem

import (
	"io/ioutil"
	"os"
)

type FileSystem interface {
	ReadCompleteFileFromDisk(path string) ([]byte, error)
	WriteCompleteFileToDisk(path string, data []byte, permissions os.FileMode) error
	DeleteFileFromDisk(path string) error
}

type fileSystem struct {
}

func NewFileSystem() FileSystem {
	return &fileSystem{}
}

func (f *fileSystem) ReadCompleteFileFromDisk(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func (f *fileSystem) WriteCompleteFileToDisk(path string, data []byte, permissions os.FileMode) error {
	return ioutil.WriteFile(path, data, permissions)
}

func (f *fileSystem) DeleteFileFromDisk(path string) error {
	return os.Remove(path)
}
