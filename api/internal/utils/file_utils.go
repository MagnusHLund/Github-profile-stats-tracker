package utils

import (
	"os"
	"path/filepath"
)

type FileUtils struct{}

func NewFileUtils() *FileUtils {
	return &FileUtils{}
}

func (fu *FileUtils) ReadFile(filePathFromRoot string) ([]byte, error) {
	workingDirectory, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	projectRoot := filepath.Join(workingDirectory, "..", "..")

	// Construct the absolute file path
	filePath := filepath.Join(projectRoot, filePathFromRoot)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return data, nil
}
