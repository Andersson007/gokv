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

	fmt.Println("Client connected:", conn.RemoteAddr())

	// TODO Move this function to handler
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading from connection:", err)
			log <- msg.New(logger.ERROR,
				"Error reading from connection:", err)
			// TODO When it return from here
			// server must continue listening
			// TODO Handle client disconnect gracefully
			return nil
		}
		fmt.Println("Client sent:", string(buf[:n]))

		// Respond to the client
		conn.Write([]byte("OK"))
	}
	return nil
}
