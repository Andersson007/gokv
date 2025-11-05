// In-memory implementation of the engine

package storage

import (
	"errors"
	"fmt"
	"sync"
)

var ErrKeyNotFound = errors.New("key not found")

// It implements the StorageEnginer interface.
type InMemStorage struct {
	mu	 sync.RWMutex
	data map[string]string
}

// Constructor
func NewInMemStorage() *InMemStorage {
	return &InMemStorage{
		data: make(map[string]string),
	}
}

func (s *InMemStorage) Set(key string, val string) error {
	// TODO Add getting and handling errors here
	// similar to the other methods
	s.mu.Lock()
	s.data[key] = val
	s.mu.Unlock()
	return nil
}

func (s *InMemStorage) Get(key string) (string, error) {
	s.mu.RLock()
	val, ok := s.data[key]
	s.mu.RUnlock()
	if !ok {
		return "", fmt.Errorf("%w: %s", ErrKeyNotFound, key)
	}
	return val, nil
}

func (s *InMemStorage) Del(key string) error {
	s.mu.Lock()
	_, ok := s.data[key]
	if ok {
		delete(s.data, key)
		s.mu.Unlock()
		return nil
	}
	s.mu.Unlock()
	return fmt.Errorf("key not found: %s", key)
}
