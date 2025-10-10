// Package logger provides logger service
package logger

import (
	"fmt"
	"log"
	"os"
)

func Log(ch chan string, logFilePath string, logFileFlags int, logFmtFlags int) {
	logFile, err := os.OpenFile(logFilePath, logFileFlags, 0644)
	if err != nil {
		fmt.Fprintln(os.Stdout, "Error opening log file:", err)
		return
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	// TODO Pass these flags from main
	log.SetFlags(logFmtFlags)

	for msg := range ch {
		// TODO add different log levels
		log.Println("Logger: ", msg)
	}
}
