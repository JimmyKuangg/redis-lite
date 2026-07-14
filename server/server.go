package server

import (
	"fmt"
	"net"
	"redis-lite/commands"
	"redis-lite/data"
	"redis-lite/storage"
	"time"
)

func NewServer() (*Server, error) {
	if err := storage.Init(); err != nil {
		return nil, err
	}

	return &Server{
		db: data.NewDatabase(),
	}, nil
}

func (s *Server) Start() {
	snapshot, err := storage.LoadSnapshot()
	if err != nil {
		panic(err)
	}

	s.db.Restore(snapshot)

	err = storage.Replay(s.db)
	if err != nil {
		panic(err)
	}

	s.StartCleanupWorker()
	defer s.Stop()

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

		s.mu.Lock()

		s.writeCount++
		shouldSnapshot := s.writeCount >= 10

		if shouldSnapshot {
			s.writeCount = 0
		}

		s.mu.Unlock()

		if shouldSnapshot {
			if err := storage.TakeSnapshot(s.db); err != nil {
				return err.Error()
			}
		}
	}

	return resp
}

func (s *Server) StartCleanupWorker() {
	ticker := time.NewTicker(10 * time.Second)

	go func() {
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				s.db.CleanupExpired()

			case <-s.stop:
				return
			}
		}
	}()
}

func (s *Server) Stop() {
	s.stop <- struct{}{}
}
