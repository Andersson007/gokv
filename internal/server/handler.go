// Package server provides logic for handling a single connection
package server

import (
	"fmt"
	"net"
	"os"

	"gokv/internal/logger"
	"gokv/internal/protocol"
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
			// TODO It should get an exit code from client,
			// write it to the log and exit
			return nil
		}

		dc := protocol.Parse(log, string(buf[:n]))
		fmt.Println(dc)
		fmt.Println("Client sent:", string(buf[:n]))

		//if dc.ctype == protocol.EXIT {
		//	log <- msg.New(logger.INFO, "Client closed connection")
		//	return nil
		//}

		// Respond to the client
		conn.Write([]byte("OK"))
	}
	return nil
}
