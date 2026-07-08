package server

import (
	"fmt"
	"net"
	"redis-lite/data"
)

type Server struct {
	db *data.Database
}

func Start() {
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

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			return
		}

		input := string(buffer[:n])
		_, err = ParseCommand(input)
		if err != nil {
			conn.Write([]byte(err.Error()))
			return
		}
	}
}
