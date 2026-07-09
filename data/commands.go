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
	db.data[key] = val
}

func (db *Database) Get(key string) (string, error) {
	val, exists := db.data[key]
	if !exists {
		return "", ErrKeyNotFound
	}

	return val, nil
}

func (db *Database) Delete(key string) error {
	_, exists := db.data[key]
	if !exists {
		return ErrKeyNotFound
	}

	delete(db.data, key)
	return nil
}

func (db *Database) Print() string {
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
