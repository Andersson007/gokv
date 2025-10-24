// Command gokv-cli: Main client entry point
package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

var Version = "0.1.0"

var (
	flagV = flag.Bool("v", false, "print version and exit")
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

func (s server) getData() string {
	// Read response
	buf := make([]byte, 1024)
	n, err := s.conn.Read(buf)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Read error:", err)
		os.Exit(1)
	}

	// Return server response as string
	return string(buf[:n])
}

func handleUserInput(scanner *bufio.Scanner, conn net.Conn) {

	fmt.Println("Enter commands (type 'exit' or press Ctrl+D to quit):")

	srv := server{conn: conn}

	for {
		fmt.Print("> ")
		if !scanner.Scan() {	// Read next line
			err := scanner.Err()
			if err == nil {
				fmt.Println("\n Ctrl+D detected (EOF)")
				// Close connection
				srv.sendData("EXIT")
			}

			fmt.Println("Input error:", err)
			os.Exit(1)
		}

		input := scanner.Text()

		if input == "exit" {
			fmt.Println("Exit")
			// Close connection
			srv.sendData("EXIT")
		}

		// Move this to a function
		if input != "" {
			fmt.Println("> You entered:", input)
			// Send data to the server
			srv.sendData(input)

			// Get response from server
			resp := srv.getData()

			// Print server response
			fmt.Println("> Server response:", resp)
		}
	}
}

func printVer() {
	fmt.Println("Version:", Version)
	os.Exit(0)
}

func main() {
	// Get CLI argument
	flag.Parse()

	if *flagV {
		printVer()
	}

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

	fmt.Println("GoKV client ver.", Version)

	handleUserInput(scanner, conn)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}
}
