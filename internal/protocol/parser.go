package protocol

import (
	"fmt"
	"strings"

	"gokv/internal/logger"
)


// Implements command parser logic
func Parse(log chan logger.LogEntry, rawCmd string) DataCmd {
	parts := strings.Fields(rawCmd)	// e.g. ["SET", "key", "TO", "val"]
	msg := logger.LogEntry{Level: logger.INFO, Msg: "Init"}

	if len(parts) < 2 {
		fmt.Println("Wrong command:", rawCmd)
		log <- msg.New(logger.ERROR, "Wrong command:", rawCmd)
		return DataCmd{}
		// TODO Error msg to client
		// TODO Maybe to pass it as a part of DataCmd?
	}

	switch parts[0] {
	case "SET":
		return DataCmd{
			ctype: SET,
			key: parts[1],
			val: parts[3],
		}

	case "GET":
		return DataCmd{
			ctype: GET,
			key: parts[1],
		}

	case "DEL":
		return DataCmd{
			ctype: DEL,
			key: parts[1],
		}

	case "EXIT":
		return DataCmd{
			ctype: EXIT,
		}

	default:
		fmt.Println("Unsupported operation:", parts[0])
		log <- msg.New(logger.ERROR, "Unsupported operation:", parts[0])
		return DataCmd{}
		// TODO Error msg to client
		// TODO Maybe to pass it as a part of DataCmd?
	}
}
