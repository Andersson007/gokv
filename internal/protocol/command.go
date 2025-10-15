// Package protocol provides parsed client commands
package protocol

type cmdType int

const (
	GET cmdType = iota
	SET
	DELETE
)

// Command represents a parsed client command.
type Command struct {
	ctype cmdType
	key int64
	val string
}
