package data

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

var (
	ErrKeyNotFound = errors.New("key does not exist")
)

func (db *Database) Set(key, val string) {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[key] = val
}

func (db *Database) Get(key string) (string, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	val, exists := db.data[key]
	if !exists {
		return "", ErrKeyNotFound
	}

	return val, nil
}

func (db *Database) Delete(key string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	_, exists := db.data[key]
	if !exists {
		return ErrKeyNotFound
	}

	delete(db.data, key)
	return nil
}

func (db *Database) Print() string {
	db.mu.RLock()
	defer db.mu.RUnlock()

	keys := make([]string, 0, len(db.data))

	for key := range db.data {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	var builder strings.Builder

	for _, key := range keys {
		val := db.data[key]

		fmt.Fprintf(&builder, "Key: %s, Val: %s\n", key, val)
	}

	return builder.String()
}

func (db *Database) MGet(keys []string) map[string]string {
	db.mu.RLock()
	defer db.mu.RUnlock()

	results := make(map[string]string)

	for _, key := range keys {
		val, err := db.Get(key)
		if err == nil {
			results[key] = val
		} else {
			results[key] = "nil"
		}
	}

	return results
}

// Utility functions

func (db *Database) Snapshot() map[string]string {
	db.mu.RLock()
	defer db.mu.RUnlock()

	copy := make(map[string]string)

	for k, v := range db.data {
		copy[k] = v
	}

	return copy
}

func (db *Database) Restore(snapshot map[string]string) {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data = snapshot
}

// Command utility

func (cmd Command) String() string {
	return strings.Join(
		append([]string{cmd.Name}, cmd.Args...),
		" ",
	)
}
