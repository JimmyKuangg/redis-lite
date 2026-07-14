package data

import (
	"sync"
	"time"
)

type Entry struct {
	Value     string
	ExpiresAt *time.Time
}

type Database struct {
	data map[string]Entry
	mu   sync.RWMutex
}

type Command struct {
	Name string
	Args []string
}
