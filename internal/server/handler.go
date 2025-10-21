// Package server provides logic for handling a single connection
package server

import (
	"fmt"
	"net"
	"os"

	"gokv/internal/logger"
	"gokv/internal/protocol"
)

type HandlerErrorCode int

const (
	ErrClientClosedConn HandlerErrorCode = iota
)

type HandlerError struct {
	Code HandlerErrorCode
	Msg string
}

func (e HandlerError) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Msg)
}

func HandleConn(log chan logger.LogEntry, conn net.Conn) error {
	msg := logger.LogEntry{Level: logger.INFO, Msg: "Init"}

	addr := fmt.Sprintf("%v", conn.RemoteAddr())
	fmt.Println("Client connected:", addr)

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

		if dc.Cmd == protocol.EXIT {
			fmt.Println("Connection closed by client", addr)
			log <- msg.New(logger.INFO,
				"Connection closed by client", addr)
			return HandlerError{
				Code: ErrClientClosedConn,
				Msg: "connection closed by client",
			}
		}

		// Respond to the client
		conn.Write([]byte("OK"))
	}
	return nil
}
