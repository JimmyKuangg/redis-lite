package storage

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	storageDir = ".redislite"
	aofFile    = "appendonly.aof"
)

func Init() error {
	err := os.MkdirAll(storageDir, 0755)
	if err != nil {
		return fmt.Errorf("error creating redislite directory: %w", err)
	}

	aofPath := filepath.Join(storageDir, aofFile)
	file, err := os.OpenFile(
		aofPath,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644,
	)
	if err != nil {
		return fmt.Errorf("error with log file: %w", err)
	}
	defer file.Close()

	return nil
}
