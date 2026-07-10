package storage

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"redis-lite/data"
	"strings"
)

func WriteSnapshot(snapshot map[string]string) error {
	snapshotPath := filepath.Join(storageDir, snapshotFile)
	file, err := os.OpenFile(snapshotPath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	for key, val := range snapshot {
		_, err = file.WriteString(key + " " + val + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func LoadSnapshot() (map[string]string, error) {
	snapshotPath := filepath.Join(storageDir, snapshotFile)
	snapshot := make(map[string]string)

	file, err := os.Open(snapshotPath)
	if err != nil {
		return snapshot, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		args := strings.Split(scanner.Text(), " ")
		if len(args) != 2 {
			return snapshot, fmt.Errorf("invalid snapshot line: %q", scanner.Text())
		}

		snapshot[args[0]] = args[1]
	}

	if err := scanner.Err(); err != nil {
		return snapshot, err
	}

	return snapshot, nil
}

func TakeSnapshot(db *data.Database) error {
	snapshot := db.Snapshot()

	err := WriteSnapshot(snapshot)
	if err != nil {
		return err
	}

	return ResetAOF()
}
