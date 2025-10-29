// In-memory implementation of the engine

package storage

import (
	"fmt"
	"sync"
)

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
	s.mu.Lock()
	s.data[key] = val
	fmt.Println(s.data)  // TODO Remove this debugging
	s.mu.Unlock()
	return nil
}

func (s *InMemStorage) Get(key string) (string, error) {
	// TODO make it allowed to read w/o locking?
	s.mu.RLock()
	val, ok := s.data[key]
	fmt.Println(s.data)  // TODO Remove this debugging
	s.mu.RUnlock()
	if !ok {
		return "", fmt.Errorf("key not found: %s", key)
	}
	return val, nil
}

func (s *InMemStorage) Del(key string) error {
	s.mu.Lock()
	_, ok := s.data[key]
	if ok {
		delete(s.data, key)
		s.mu.Unlock()
		fmt.Println(s.data)  // TODO Remove this debugging
		return nil
	}
	s.mu.Unlock()
	return fmt.Errorf("key not found: %s", key)
}
