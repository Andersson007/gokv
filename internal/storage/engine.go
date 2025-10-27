// Package storage provides a storage engine interface
package storage

// StorageEngine is the interface that all storage backends must implement.
type StorageEngine interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Del(key string) error
}
