package commands

import "redis-lite/data"

func IsWriteCommand(cmd data.Command) bool {
	switch cmd.Name {
	case "SET", "MSET", "DEL":
		return true
	}

	return false
}
