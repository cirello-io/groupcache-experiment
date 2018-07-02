package storage

import (
	"math/rand"
	"sync"
	"time"
)

// Storage implements a simple KV manipulator.
type Storage struct {
	mu   sync.Mutex
	data map[string]string
}

// New instantiates a new KV manipulator storage.
func New() *Storage {
	return &Storage{
		data: make(map[string]string),
	}
}

// Get loads key from the storage. It uses a random sleep to artifically add
// some latency.
func (db *Storage) Get(key string) string {
	// It adds some latency, so it becomes more obvious when the result
	// is being returned by the storage.
	time.Sleep(time.Duration(rand.Intn(1500)+1500) * time.Millisecond)
	db.mu.Lock()
	defer db.mu.Unlock()
	return db.data[key]
}

// Set saves a key-value pair in the storage.
func (db *Storage) Set(key string, value string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.data[key] = value
}
