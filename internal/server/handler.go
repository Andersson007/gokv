// Package server provides logic for handling a single connection
package server

import (
	"errors"
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
	store storage.StorageEnginer,
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
			return nil
		}

		fmt.Println("Client sent:", string(buf[:n]))
		dc := protocol.Parse(log, string(buf[:n]))
		fmt.Println(dc)
		// TODO Move to the switch to a function
		switch dc.Cmd {
		case protocol.GET:
			val, err := store.Get(dc.Key)
			if err != nil {
				if errors.Is(err, storage.ErrKeyNotFound) {
					// TODO Should I just remove this line or log this?
					fmt.Println("The key doesn't exist:", dc.Key)
					conn.Write([]byte("The key doesn't exist: " + dc.Key))
    			} else {
					// TODO Should I just remove this line or log this?
        			fmt.Println("Error:", err)
					conn.Write([]byte("Error: " + dc.Key))
    			}
			} else {
				// If not errors, send the val to the client
				conn.Write([]byte(val))
			}
		case protocol.SET:
			store.Set(dc.Key, dc.Val)
			conn.Write([]byte("SET"))
		case protocol.DEL:
			store.Del(dc.Key)
			conn.Write([]byte("DELETED"))
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
	}
	return nil
}
