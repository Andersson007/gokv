// Package protocol provides parsed client commands
package protocol

type cmdType int

const (
	GET cmdType = iota
	SET
	DEL
	EXIT
)

// Command represents a parsed client command
type DataCmd struct {
	Cmd cmdType
	Key string
	Val string
}
