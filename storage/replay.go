package storage

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"redis-lite/commands"
	"redis-lite/data"
)

func Replay(db *data.Database) error {
	path := filepath.Join(storageDir, aofFile)

	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("error opening aof file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		cmd, err := commands.ParseCommand(line)
		if err != nil {
			return fmt.Errorf("error parsing command: %w", err)
		}

		_, err = commands.ExecuteCommand(db, cmd)
		if err != nil {
			return fmt.Errorf("error executing command: %w", err)
		}
	}

	return scanner.Err()
}
