// Package logger provides logger service
package logger

import (
	"fmt"
	"log"
	"os"
)

func Log(ch chan string, logFilePath string) {
	logFile, err := os.OpenFile(logFilePath,
		os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintln(os.Stdout, "Error opening log file:", err)
		return
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	for msg := range ch {
		// TODO add different log levels
		log.Println("Logger: ", msg)
	}
}
