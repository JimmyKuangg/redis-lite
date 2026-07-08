package data

import (
	"errors"
	"fmt"
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

func (db *Database) Print() {
	for key, val := range db.data {
		fmt.Printf("Key: %s, Val: %s", key, val)
	}
}