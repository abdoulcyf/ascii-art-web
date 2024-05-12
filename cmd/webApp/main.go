package main

import (
	"fmt"

	"github.com/ediallocyf/asciiartweb/internal/servers"
)

const (
	portNumber = ":8080"
)

func main() {
	server := servers.NewAPIServer(portNumber)
	fmt.Println("Server is running at port: 8080")
	server.Run()
}
