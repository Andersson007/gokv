// Package server provides TCP listener, connection handling
package server

import (
	"fmt"
	"net"
	"os"

	"gokv/internal/logger"
	"gokv/internal/storage"
)

// Listen for connections
func Listen(log chan logger.LogEntry, protocol string, port int) error {
	// Write this to log as well
	portStr := fmt.Sprintf(":%v", port)
	msg := logger.LogEntry{Level: logger.INFO, Msg: "Init"}

	// Initialize storage
	// TODO introduce other storage types
	storage := storage.NewInMemStorage()

	// Start listening
	ln, err := net.Listen(protocol, portStr)
	if err != nil {
		log <- msg.New(logger.ERROR, "Server error", err)
		return err
	}
	defer func() {
		if err := ln.Close(); err != nil {
			log <- msg.New(logger.ERROR, "Error closing listener:", err)
			os.Exit(1)
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

		// Launch a conn handler
		// TODO assign a unique ID to every
		// goroutine to be able to distinguish
		// which handler the log messages come from
		go HandleConn(log, conn, storage)
	}
	return nil
}
