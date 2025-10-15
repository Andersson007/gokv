// Command gokv-server: Main server entry point
package main

import (
	"fmt"
	"log"
	"os"

	"gokv/internal/logger"
	"gokv/internal/protocol"	// TODO Needs removal from here
	"gokv/internal/server"
)

func main() {
	// Some defaults here
	// TODO Use a config file later to overwrite them
	proto := "tcp"
	port := 5454
	loggerChanBufferLen := 100
	logFilePath := "/tmp/gokv-server.log"
	logFileFlags := os.O_CREATE | os.O_APPEND | os.O_WRONLY
	logFmtFlags := log.Ldate | log.Ltime
	msg := logger.LogEntry{Level: logger.INFO, Msg: "Init"}

	// Start the logger service
	log := make(chan logger.LogEntry, loggerChanBufferLen)
	go logger.Log(log, logFilePath, logFileFlags, logFmtFlags)

	log <- msg.New(logger.INFO, "GoKV Server starting")

	// TODO Must be removed after debugging
	// TODO How to set log only once?
	dc1 := protocol.Parse(log, "SET key TO val")
	dc2 := protocol.Parse(log, "GET key")
	dc3 := protocol.Parse(log, "DEL key")
	dc4 := protocol.Parse(log, "too-short")
	dc5 := protocol.Parse(log, "UNSUPPORTED key")
	fmt.Println("SET key TO val", dc1)
	fmt.Println("GET key", dc2)
	fmt.Println("DEL key", dc3)
	fmt.Println("too-short", dc4)
	fmt.Println("UNSUPPORTED key", dc5)

	if err := server.Listen(log, proto, port); err != nil {
		fmt.Fprintln(os.Stderr, "server error:", err)
		log <- msg.New(logger.INFO, fmt.Sprintln("Server error:", err))
	}
}
