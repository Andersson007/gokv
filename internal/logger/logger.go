// Package logger provides logger service
package logger

import (
	"fmt"
)

func Log(ch chan string) {
	for msg := range ch {		// Loop until chan is closed
		// TODO Just print to stdout for now
		// Replace with logging to a file later
		fmt.Println("Logger: ", msg)
	}
}
