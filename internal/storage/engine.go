// Package storage provides a storage engine interface
package storage

// Engine is the interface that all storage backends must implement.
type Engine interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}
