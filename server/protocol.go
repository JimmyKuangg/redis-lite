package server

import (
	"errors"
	"fmt"
	"redis-lite/data"
	"strings"
)

func ParseCommand(input string) (data.Command, error) {
	var command data.Command

	args := strings.Fields(input)

	if len(args) == 0 {
		return command, fmt.Errorf("Please enter a command\n")
	}

	return data.Command{
		Name: strings.ToUpper((args[0])),
		Args: args[1:],
	}, nil
}

func ExecuteCommand(db *data.Database, cmd data.Command) error {
	switch cmd.Name {
	case "GET":
		if len(cmd.Args) != 1 {
			return errors.New("GET expects exactly one key")
		}

		_, err := db.Get(cmd.Args[0])
		return err

	case "SET":
		if len(cmd.Args) != 2 {
			return errors.New("SET expects a key and a value")
		}

		db.Set(cmd.Args[0], cmd.Args[1])
		return nil

	case "DEL":
		if len(cmd.Args) != 1 {
			return errors.New("DEL expects exactly one key")
		}

		return db.Delete(cmd.Args[0])

	case "PRINT":
		db.Print()
	}

	return fmt.Errorf("unknown command: %s", cmd.Name)
}
