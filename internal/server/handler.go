// Package server provides logic for handling a single connection
package server

import (
	"fmt"
	"net"
	"os"

	"gokv/internal/logger"
)

func HandleConn(log chan logger.LogEntry, conn net.Conn) error {
	msg := logger.LogEntry{Level: logger.INFO, Msg: "Init"}

	fmt.Println("Client connected:", conn.RemoteAddr())

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
