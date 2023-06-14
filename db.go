package asinodb

import "errors"

var (
	// ErrNothing error for not existing key
	ErrNothing = errors.New("asin: key not found")
)

// New instantiates a new key-value database
func New() *Database {
	return &Database{
		storage: newStorage(),
	}
}

// Database is a disk-backed key-value database.
type Database struct {
	storage *Storage
}

// Get retrieves a value for a given key.
func (d *Database) Get(key string) (interface{}, error) {
	return d.storage.get(key)
}

// Set stores a value for a given key.
func (d *Database) Set(key string, value interface{}) error {
	return d.storage.set(key, value)
}
