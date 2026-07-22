package tests

import (
	"fmt"
	"net"
	"redis-lite/server"
	"sync"
	"testing"
	"time"
)

func TestServerStress(t *testing.T) {
	s, _ := server.NewServer()
	go s.Start()
	defer s.Stop()

	clients := 50
	requestsPerClient := 10000

	start := time.Now()
	var wg sync.WaitGroup

	for i := 0; i < clients; i++ {
		wg.Add(1)
		go func(clientID int) {
			defer wg.Done()
			conn, _ := net.Dial("tcp", ":6379")
			defer conn.Close()

			for j := 0; j < requestsPerClient; j++ {
				cmd := fmt.Sprintf("SET key_%d_%d val\n", clientID, j)
				conn.Write([]byte(cmd))
				buf := make([]byte, 1024)
				conn.Read(buf)
			}
		}(i)
	}

	wg.Wait()
	elapsed := time.Since(start)

	totalOps := clients * requestsPerClient
	opsPerSec := float64(totalOps) / elapsed.Seconds()
	fmt.Printf("%d clients, %d ops total, %.0f ops/sec\n",
		clients, totalOps, opsPerSec)
}
