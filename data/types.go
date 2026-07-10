package data

import "sync"

type Database struct {
	data map[string]string
	mu   sync.RWMutex
}

type Command struct {
	Name string
	Args []string
}
