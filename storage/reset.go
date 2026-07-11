package storage

import (
	"os"
	"path/filepath"
)

func ResetAOF() error {
	storageMu.Lock()
	defer storageMu.Unlock()

	aofPath := filepath.Join(storageDir, aofFile)
	file, err := os.OpenFile(aofPath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	return nil
}
