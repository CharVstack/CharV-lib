package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func IsPoolExists(path string) bool {
	f, _ := os.Stat(path)
	return f.IsDir()
}

func GetPoolFiles(storageDir string) ([]string, error) {
	entries, err := os.ReadDir(storageDir)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, entry := range entries {
		files = append(files, entry.Name())
	}
	return files, nil
}

func GetPoolInfo(file string, storageDir string) (*PoolInfo, error) {
	path, _ := filepath.Abs(filepath.Join(storageDir, file))
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var storagePool PoolInfo
	err = json.Unmarshal(bytes, &storagePool)
	if err != nil {
		return nil, err
	}
	return &storagePool, nil
}
