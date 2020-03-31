package filemanager

import (
	"fmt"
	"io/ioutil"
	"os"
)

type FileManager struct {
}

func (f *FileManager) OpenFile(filename string, flags int, permissions os.FileMode) (*os.File, error) {
	return os.OpenFile(filename, flags, permissions)
}

func (f *FileManager) SaveFile(filename, path string, data []byte) error {
	fullPath := buildFilePath(filename, path)

	err := ioutil.WriteFile(fullPath, data, 0644)

	return err
}

func buildFilePath(filename, path string) string {
	if path == "" {
		return filename
	}
	return fmt.Sprintf("%s/%s", path, filename)
}
