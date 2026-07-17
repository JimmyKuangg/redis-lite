package commands

import (
	"fmt"
	"redis-lite/data"
	"strings"
)

func ParseCommand(input string) (data.Command, error) {
	var command data.Command

	args, err := splitArgs(input)
	if err != nil {
		return command, err
	}

	if len(args) == 0 {
		return command, fmt.Errorf("Please enter a command")
	}

	return data.Command{
		Name: strings.ToUpper((args[0])),
		Args: args[1:],
	}, nil
}

func splitArgs(input string) ([]string, error) {
	var args []string
	var current strings.Builder
	inQuotes := false
	escaped := false

	for _, r := range input {
		switch {
		case escaped:
			current.WriteRune(r)
			escaped = false

		case r == '\\':
			escaped = true

		case r == '"':
			inQuotes = !inQuotes

		case (r == ' ' || r == '\t' || r == '\n' || r == '\r') && !inQuotes:
			if current.Len() > 0 {
				args = append(args, current.String())
				current.Reset()
			}

		default:
			current.WriteRune(r)
		}
	}

	if escaped {
		return nil, fmt.Errorf("unfinished escape sequence")
	}

	if inQuotes {
		return nil, fmt.Errorf("unterminated quoted string")
	}

	if current.Len() > 0 {
		args = append(args, current.String())
	}

	return args, nil
}
