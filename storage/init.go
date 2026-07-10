package storage

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	storageDir   = ".redislite"
	aofFile      = "appendonly.aof"
	snapshotFile = "snapshot.rdb"
)

func Init() error {
	err := os.MkdirAll(storageDir, 0755)
	if err != nil {
		return fmt.Errorf("error creating redislite directory: %w", err)
	}

	aofPath := filepath.Join(storageDir, aofFile)
	snapshotPath := filepath.Join(storageDir, snapshotFile)

	err = ensureFile(aofPath)
	if err != nil {
		return fmt.Errorf("creating  AOF: %w", err)
	}

	err = ensureFile(snapshotPath)
	if err != nil {
		return fmt.Errorf("creating snapshot: %w", err)
	}

	return nil
}

func ensureFile(path string) error {
	file, err := os.OpenFile(
		path,
		os.O_CREATE|os.O_RDWR,
		0644,
	)
	if err != nil {
		return fmt.Errorf("error during init: %w", err)
	}
	defer file.Close()

	return nil
}
