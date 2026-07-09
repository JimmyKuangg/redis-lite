package main

import (
	"fmt"
	"redis-lite/server"
)

func main() {
	s, err := server.NewServer()
	if err != nil {
		fmt.Println(err)
	}

	s.Start()
}
