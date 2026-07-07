package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		panic(err)
	}

	fmt.Println("RedisLite listening on port 6379")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		fmt.Println("Client connected")

		conn.Close()
	}
}