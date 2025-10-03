// Command gokv-server: Main server entry point
package main

import (
	"fmt"

	"gokv/internal/server"
)

func main() {
	fmt.Println("GoKV Server")

	if err := server.Listen(); err != nil {
		// TODO replace it with calling logger
		fmt.Println("server error:", err)
	}
}
