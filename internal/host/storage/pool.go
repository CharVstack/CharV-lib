package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func IsPoolExists(path string) (bool, error) {
	f, err := os.Stat(path)
	return err == nil && f.IsDir(), err
}

func GetPoolFiles(storageDir string) []string {
	entries, err := os.ReadDir(storageDir)
	if err != nil {
		fmt.Println(err)
	}

	var files []string
	for _, entry := range entries {
		files = append(files, entry.Name())
	}
	return files
}

func GetPoolInfo(file string, storageDir string) *PoolInfo {
	path, _ := filepath.Abs(filepath.Join(storageDir, file))
	bytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	var storagePool PoolInfo
	err = json.Unmarshal(bytes, &storagePool)
	if err != nil {
		fmt.Println(err)
	}
	return &storagePool
}
