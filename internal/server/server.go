// Package server provides TCP listener, connection handling
package server

import (
	"fmt"
	"net"
	"os"

	"gokv/internal/logger"
)

// Listen for connections
func Listen(log chan logger.LogEntry, protocol string, port int) error {
	// Write this to log as well
	portStr := fmt.Sprintf(":%v", port)
	msg := logger.LogEntry{Level: logger.INFO, Msg: "Init"}

	// Start listening
	ln, err := net.Listen(protocol, portStr)
	if err != nil {
		fmt.Println("Server error:", err)
		log <- msg.New(logger.ERROR, "Server error", err)
		return err
	}
	defer func() {
		if err := ln.Close(); err != nil {
			fmt.Println("Error closing listener:", err)
		}
	}()

	welcomeMsg := fmt.Sprintf("Listening on %s/%v ...\n", protocol, port)
	fmt.Printf(welcomeMsg)
	log <- msg.New(logger.INFO, welcomeMsg)

	// Accept a single connection
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Server error:", err)
			return err
		}
		defer func() {
			if err := conn.Close(); err != nil {
				fmt.Fprintln(os.Stderr, "Error closing connection:", err)
				log <- msg.New(logger.ERROR, "Error closing connection:", err)
			}
		}()

		// Launch a conn handler
		go HandleConn(log, conn)
	}
	return nil
}
