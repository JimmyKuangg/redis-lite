package storage

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"redis-lite/commands"
	"redis-lite/data"
	"strings"
)

func Replay(db *data.Database) error {
	path := filepath.Join(storageDir, aofFile)

	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("error opening aof file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var cmdsRan int

	fmt.Println("Replaying AOF...")
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		cmd, err := commands.ParseCommand(line)
		if err != nil {
			return fmt.Errorf("error parsing command: %w", err)
		}

		_, err = commands.ExecuteCommand(db, cmd)
		if err != nil {
			return fmt.Errorf("error executing command: %w", err)
		}

		cmdsRan++
	}

	fmt.Printf("AOF replay complete. %d commands applied\n", cmdsRan)
	return scanner.Err()
}
