// Package storage provides a storage engine interface
package storage

// StorageEnginer is the interface that all storage backends must implement.
type StorageEnginer interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Del(key string) error
}
