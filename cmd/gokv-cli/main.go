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

	conn, err := net.Dial(proto, host + ":" + port)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Connection error", err)
		os.Exit(1)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("GoKV client ver.", clientVer)
	fmt.Println("Enter commands (type 'exit' or press Ctrl+D to quit):")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {	// Read next line
			break				// EOF (including Ctrl+D) or error
		}

		input := scanner.Text()

		if input == "exit" {
			fmt.Println("Exit")
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
