package api

import "errors"

// KeyNotFound indicates that the provided key does not exist in the Store
// It can be returned on by a Get function call
var KeyNotFound = errors.New("key not found")

// A store is a base interface for a KV store
// The only method needed is the close method which should be called when all operations are completed and the connection can be shut down

// Closer represents something that can be closed: either a database connection or a transaction
type Closer interface {
	Close() error
}

// Getter wraps the Get method
type Getter interface {
	// Get fetches the value from a given key and returns it
	// It will return a KeyNotFound error if the provided key is not found
	Get(key string) (result []byte, err error)
}

// Setter wraps the Set method
type Setter interface {
	// Set sets the value for a given key
	Set(key string, value []byte) (err error)
}

// AGetter is an async wrapper the Get method
type AGetter interface {
	// Get fetches the value from a given key and returns it
	// It will return a KeyNotFound error if the provided key is not found
	Get(key string) (result <-chan []byte, err <-chan error)
}

// Setter wraps the Set method
type ASetter interface {
	// Set sets the value for a given key
	Set(key string, value []byte) (err <-chan error)
}

type RWStore interface {
	Closer
	Getter
	Setter
}
