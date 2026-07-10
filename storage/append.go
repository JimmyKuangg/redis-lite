package storage

import (
	"os"
	"path/filepath"
)

func Append(command string) error {
	aofPath := filepath.Join(storageDir, aofFile)

	file, err := os.OpenFile(
		aofPath,
		os.O_APPEND|os.O_WRONLY,
		0644,
	)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(command + "\n")
	return err
}
