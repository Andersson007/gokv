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

func respondToClient(
	log chan logger.LogEntry,
	conn net.Conn,
	store storage.StorageEnginer,
	dc protocol.DataCmd,
	addr string,
	msg logger.LogEntry,
) error {
	switch dc.Cmd {
	case protocol.GET:
		val, err := store.Get(dc.Key)
		if err != nil {
			if errors.Is(err, storage.ErrKeyNotFound) {
				conn.Write([]byte("The key doesn't exist: " + dc.Key))
   			} else {
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

		// TODO This is discarded now. Think of creating a channel
		// to return errors from handler goroutines if needed
		// TODO UPDATE: I know the case. When connection is closed
		// by the client, server Listen() reports the error
		// "Error reading from connection: EOF". The solution could
		// be 1. Creating a channel 2. Assigning ID to each
		// handler goroutine 3. There'll be a check in the reading
		// from conn loop checking if the gorouting has reported
		// that connection was closed by the client.
		return HandlerError{
			Code: ErrClientClosedConn,
			Msg: "connection closed by client",
		}
	}
	return nil
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

		dc := protocol.Parse(log, string(buf[:n]))

		// TODO This function returns an error.
		// Think if you need it. TODO see comments in the function
		respondToClient(log, conn, store, dc, addr, msg)
	}
	return nil
}
