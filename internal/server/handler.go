// Package server provides logic for handling a single connection
package server

import (
	"fmt"
	"net"
	"os"

	"gokv/internal/logger"
	"gokv/internal/protocol"
	"gokv/internal/storage"
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

func HandleConn(
	log chan logger.LogEntry,
	conn net.Conn,
	storage storage.StorageEnginer,
) error {

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

		fmt.Println("Client sent:", string(buf[:n]))
		dc := protocol.Parse(log, string(buf[:n]))
		fmt.Println(dc)
		// TODO Move to the switch to a function
		switch dc.Cmd {
		case protocol.GET:
			storage.Get(dc.Key)
		case protocol.SET:
			storage.Set(dc.Key, dc.Val)
		case protocol.DEL:
			storage.Del(dc.Key)
		}

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
