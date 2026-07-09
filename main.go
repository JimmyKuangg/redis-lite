package main

import (
	"redis-lite/server"
)

func main() {
	s := server.NewServer()
	s.Start()
}
