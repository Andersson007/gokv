// Command gokv-cli: Main client entry point
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type server struct {
	conn net.Conn
}

func (s server) sendData(data string) {
	_, err := s.conn.Write([]byte(data))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Write error:", err)
		os.Exit(1)
	}
}

func handleUserInput(scanner *bufio.Scanner, conn net.Conn) {

	fmt.Println("Enter commands (type 'exit' or press Ctrl+D to quit):")

	srv := server{conn: conn}

	for {
		fmt.Print("> ")
		if !scanner.Scan() {	// Read next line
			err := scanner.Err()
			// TODO Move closing connection to a separate func
			if err == nil {
				fmt.Println("\n Ctrl+D detected (EOF)")
				srv.sendData("EXIT")
			}

			fmt.Println("Input error:", err)
			os.Exit(1)
		}

		input := scanner.Text()

		if input == "exit" {
			fmt.Println("Exit")
			// TODO Move closing connection to a separate func
			// Notify the server
			srv.sendData("EXIT")
		}

		// Move this to a function
		if input != "" {
			fmt.Println("> You entered:", input)
			// Send data
			srv.sendData(input)

			// Read response
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Read error:", err)
				return
			}

			// Print server response
			fmt.Println("> Server response:", string(buf[:n]))
		}
	}
}

func main() {
	clientVer := "0.1.0"  // TODO More idiomatic way?
	// Some defaults here
	// TODO Overwrite them by using CLI args
	host := "localhost"
	proto := "tcp"
	port := "5454"

	conn, err := net.Dial(proto, host + ":" + port)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Connection error", err)
		os.Exit(1)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("GoKV client ver.", clientVer)

	handleUserInput(scanner, conn)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}
}
