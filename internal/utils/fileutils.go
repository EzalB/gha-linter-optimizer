package utils

import (
	"os"
	"path/filepath"
)

func GetWorkflowFiles(basePath string) []string {
	var files []string

	_ = filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if filepath.Ext(path) == ".yml" || filepath.Ext(path) == ".yaml" {
			files = append(files, path)
		}
		return nil
	})

	return files
}
