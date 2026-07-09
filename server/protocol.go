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
		return command, fmt.Errorf("Please enter a command")
	}

	return data.Command{
		Name: strings.ToUpper((args[0])),
		Args: args[1:],
	}, nil
}

func ExecuteCommand(db *data.Database, cmd data.Command) (string, error) {
	switch cmd.Name {
	case "GET":
		if len(cmd.Args) != 1 {
			return "", errors.New("GET expects exactly one key")
		}

		resp, err := db.Get(cmd.Args[0])
		return resp, err

	case "SET":
		if len(cmd.Args) != 2 {
			return "", errors.New("SET expects a key and a value")
		}

		db.Set(cmd.Args[0], cmd.Args[1])
		return "success", nil

	case "MSET":
		if len(cmd.Args) == 0 || len(cmd.Args)%2 != 0 {
			return "", errors.New("invalid format for SET")
		}

		for i := 0; i < len(cmd.Args); i += 2 {
			db.Set(cmd.Args[i], cmd.Args[i+1])
		}

		return "success", nil

	case "DEL":
		if len(cmd.Args) != 1 {
			return "", errors.New("DEL expects exactly one key")
		}

		for _, arg := range cmd.Args {
			_, err := db.Get(arg)
			if err != nil {
				return "", fmt.Errorf("error in DELETE, %w", err)
			}
		}

		for _, arg := range cmd.Args {
			err := db.Delete(arg)
			if err != nil {
				return "", fmt.Errorf("error in DELETE, %w", err)
			}
		}

	case "PRINT":
		return db.Print(), nil
	}

	return "", fmt.Errorf("unknown command: %s", cmd.Name)
}
