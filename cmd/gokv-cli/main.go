// Command gokv-cli: Main client entry point
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	clientVer := "0.1.0"  // TODO More idiomatic way?
	// Some defaults here
	// TODO Overwrite them by using CLI args
	host := "localhost"
	proto := "tcp"
	port := "5454"

	// TODO move this to a function
	conn, err := net.Dial(proto, host + ":" + port)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Connection error", err)
		os.Exit(1)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("GoKV client ver.", clientVer)
	fmt.Println("Enter commands (type 'exit' or press Ctrl+D to quit):")

	// TODO move this to a function
	for {
		fmt.Print("> ")
		if !scanner.Scan() {	// Read next line
			err := scanner.Err()
			// TODO Move closing connection to a separate func
			if err == nil {
				fmt.Println("\n Ctrl+D detected (EOF)")
				_, conn_err := conn.Write([]byte("EXIT"))
				if conn_err != nil {
					fmt.Fprintln(os.Stderr, "Write error:", conn_err)
					return
				}
			}
			break
		}

		input := scanner.Text()

		if input == "exit" {
			fmt.Println("Exit")
			// TODO Move closing connection to a separate func
			// Notify the server
			_, err = conn.Write([]byte("EXIT"))
			if err != nil {
				fmt.Fprintln(os.Stderr, "Write error:", err)
				return
			}
			break
		}

		if input != "" {
			fmt.Println("> You entered:", input)
			// Send data
			_, err = conn.Write([]byte(input))
			if err != nil {
				fmt.Fprintln(os.Stderr, "Write error:", err)
				return
			}

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

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}
}
