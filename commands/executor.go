package commands

import (
	"errors"
	"fmt"
	"redis-lite/data"
	"strings"
)

func ExecuteCommand(db *data.Database, cmd data.Command) (string, error) {
	switch cmd.Name {
	case "GET":
		if len(cmd.Args) != 1 {
			return "", errors.New("GET expects exactly one key")
		}

		resp, err := db.Get(cmd.Args[0])
		return resp, err

	case "MGET":
		values := db.MGet(cmd.Args)

		var builder strings.Builder

		for key, val := range values {
			fmt.Fprintf(&builder, "%s: %s\n", key, val)
		}

		return builder.String(), nil

	case "SET":
		if len(cmd.Args) != 2 {
			return "", errors.New("SET expects a key and a value")
		}

		db.Set(cmd.Args[0], cmd.Args[1])

		return "OK", nil

	case "MSET":
		if len(cmd.Args) == 0 || len(cmd.Args)%2 != 0 {
			return "", errors.New("invalid format for SET")
		}

		for i := 0; i < len(cmd.Args); i += 2 {
			db.Set(cmd.Args[i], cmd.Args[i+1])
		}

		return "OK", nil

	case "DEL":
		if len(cmd.Args) < 1 {
			return "", errors.New("DEL expects at least one key")
		}

		for _, arg := range cmd.Args {
			_, err := db.Get(arg)
			if err != nil {
				return "", fmt.Errorf("error in DEL, %w", err)
			}
		}

		for _, arg := range cmd.Args {
			err := db.Delete(arg)
			if err != nil {
				return "", fmt.Errorf("error in DEL, %w", err)
			}
		}

		return "", nil
	case "PRINT":
		return db.Print(), nil

	case "PING":
		return "PONG", nil
	}

	return "", fmt.Errorf("unknown command: %s", cmd.Name)
}
