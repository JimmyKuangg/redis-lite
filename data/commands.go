package data

import (
	"errors"
	"fmt"
	"maps"
	"sort"
	"strings"
	"time"
)

var (
	ErrKeyNotFound = errors.New("key does not exist")
)

func (db *Database) Set(key, val string) {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[key] = Entry{
		Value: val,
	}
}

func (db *Database) Get(key string) (string, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	return db.get(key)
}

func (db *Database) MGet(keys []string) map[string]string {
	db.mu.RLock()
	defer db.mu.RUnlock()

	results := make(map[string]string)

	for _, key := range keys {
		val, err := db.get(key)

		if err == nil {
			results[key] = val
		} else {
			results[key] = "nil"
		}
	}

	return results
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

func (db *Database) Expire(key string, duration int) {
	db.mu.Lock()
	defer db.mu.Unlock()

	expiry := time.Now().Add(time.Duration(duration) * time.Second)

	entry, exists := db.data[key]
	if !exists {
		return
	}

	entry.ExpiresAt = &expiry
	db.data[key] = entry
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

	for key, entry := range db.data {
		fmt.Fprintf(
			&builder,
			"%s => %s, TTL: %v\n",
			key,
			entry.Value,
			entry.ExpiresAt,
		)
	}

	return builder.String()
}

// Unsafes

func (db *Database) get(key string) (string, error) {
	entry, exists := db.data[key]

	if !exists {
		return "", ErrKeyNotFound
	}

	if entry.ExpiresAt != nil && time.Now().After(*entry.ExpiresAt) {
		delete(db.data, key)
		return "", ErrKeyNotFound
	}

	return entry.Value, nil
}

// Utility functions

func (db *Database) Snapshot() map[string]Entry {
	db.mu.RLock()
	defer db.mu.RUnlock()

	snapshot := make(map[string]Entry)
	maps.Copy(snapshot, db.data)

	return snapshot
}

func (db *Database) Restore(snapshot map[string]Entry) {
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
