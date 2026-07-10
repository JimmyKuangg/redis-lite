package server

import (
	"fmt"
	"net"
	"redis-lite/commands"
	"redis-lite/data"
	"redis-lite/storage"
)

type Server struct {
	db         *data.Database
	writeCount int
}

func NewServer() (*Server, error) {
	if err := storage.Init(); err != nil {
		return nil, err
	}

	return &Server{
		db: data.NewDatabase(),
	}, nil
}

func (s *Server) Start() {
	err := storage.Replay(s.db)
	if err != nil {
		panic(err)
	}

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

		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			return
		}

		input := string(buffer[:n])
		response := s.handleRequest(input)

		if err := WriteResponse(conn, response); err != nil {
			return
		}
	}
}

func (s *Server) handleRequest(input string) string {
	cmd, err := commands.ParseCommand(input)
	if err != nil {
		return err.Error()
	}

	resp, err := commands.ExecuteCommand(s.db, cmd)
	if err != nil {
		return err.Error()
	}

	if commands.IsWriteCommand(cmd) {
		if err := storage.Append(cmd.String()); err != nil {
			return err.Error()
		}

		s.writeCount++
	}

	if s.writeCount >= 10 {
		err := storage.TakeSnapshot(s.db)
		if err != nil {
			return err.Error()
		}

		s.writeCount = 0
	}

	return resp
}
