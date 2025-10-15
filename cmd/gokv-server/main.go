// Command gokv-server: Main server entry point
package main

import (
	"fmt"
	"log"
	"os"

	"gokv/internal/logger"
	"gokv/internal/server"
)

func main() {
	// Some defaults here
	// TODO Use a config file later to overwrite them
	protocol := "tcp"
	port := 5454
	loggerChanBufferLen := 100
	logFilePath := "/tmp/gokv-server.log"
	logFileFlags := os.O_CREATE | os.O_APPEND | os.O_WRONLY
	logFmtFlags := log.Ldate | log.Ltime
	msg := logger.LogEntry{Level: logger.INFO, Msg: "Init"}

	// Start the logger service
	log := make(chan logger.LogEntry, loggerChanBufferLen)
	go logger.Log(log, logFilePath, logFileFlags, logFmtFlags)

	log <- msg.With(logger.INFO, "GoKV Server starting")

	if err := server.Listen(log, protocol, port); err != nil {
		fmt.Fprintln(os.Stderr, "server error:", err)
		log <- msg.With(logger.INFO, fmt.Sprintln("Server error:", err))
	}
}
