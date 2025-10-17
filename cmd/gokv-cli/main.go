// Command gokv-cli: Main client entry point
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	clientVer := "0.1.0"  // TODO More idiomatic way?
	// Some defaults here
	// TODO Overwrite them by using CLI args
	// port :=
	// proto :=

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
			fmt.Println("You entered:", input)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}
}
