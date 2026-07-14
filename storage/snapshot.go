package storage

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"redis-lite/data"
	"strings"
	"sync"
	"time"
)

var storageMu sync.Mutex

func WriteSnapshot(snapshot map[string]data.Entry) error {
	storageMu.Lock()
	defer storageMu.Unlock()

	fmt.Println("Writing to snapshot file")
	snapshotPath := filepath.Join(storageDir, snapshotFile)
	file, err := os.OpenFile(snapshotPath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	for key, entry := range snapshot {
		if entry.ExpiresAt != nil && time.Now().After(*entry.ExpiresAt) {
			continue
		}

		_, err = file.WriteString(key + " " + entry.Value + "\n")
		if err != nil {
			return err
		}
	}

	fmt.Println("Finished writing to snapshot file")
	return nil
}

func LoadSnapshot() (map[string]data.Entry, error) {
	snapshotPath := filepath.Join(storageDir, snapshotFile)
	snapshot := make(map[string]data.Entry)

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

		snapshot[args[0]] = data.Entry{
			Value: args[1],
		}
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
