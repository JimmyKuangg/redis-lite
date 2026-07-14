package server

import (
	"redis-lite/data"
	"sync"
)

type Server struct {
	db         *data.Database
	writeCount int
	mu         sync.Mutex
	stop       chan struct{}
}
