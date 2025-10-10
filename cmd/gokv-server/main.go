// Command gokv-server: Main server entry point
package main

import (
	"fmt"
	"os"

	"gokv/internal/logger"
	"gokv/internal/server"
)

func main() {
	// Default protocol and port
	// TODO Use a config file later to overwrite them
	protocol := "tcp"
	port := 5454
	loggerChanBufferLen := 100

	// Start the logger service
	log := make(chan string, loggerChanBufferLen)
	go logger.Log(log)

	fmt.Println("GoKV Server")

	if err := server.Listen(log, protocol, port); err != nil {
		fmt.Fprintln(os.Stderr, "server error:", err)
		log <- fmt.Sprintln("Server error:", err)
	}
}
