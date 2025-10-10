// Package server provides TCP listener, connection handling
package server

import (
	"fmt"
	"net"
)

// Listen for connections
func Listen(log chan string, protocol string, port int) error {
	// Write this to log as well
	portStr := fmt.Sprintf(":%v", port)

	// Start listening
	ln, err := net.Listen(protocol, portStr)
	if err != nil {
		// TODO Replace it with calling a logger
		// here and all over the file
		fmt.Println("Server error:", err)
		return err
	}
	defer func() {
		if err := ln.Close(); err != nil {
			fmt.Println("Error closing listener:", err)
		}
	}()

	msg := fmt.Sprintf("Listening on %s/%v ...\n", protocol, port)
	fmt.Printf(msg)
	log <- msg

	// Accept a single connection
	conn, err := ln.Accept()
	if err != nil {
		fmt.Println("Server error:", err)
		return err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println("Error closing connection:", err)
		}
	}()

	fmt.Println("Client connected:", conn.RemoteAddr())
	return nil
}
