package commands

import (
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
