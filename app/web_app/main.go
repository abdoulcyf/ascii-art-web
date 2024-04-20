package main

import (
	"github.com/ediallocyf/asciiartweb/internal/handlers"
)

const (
	portAddressNumber = ":8080"
)

func main() {
	server := handlers.NewAPIServer(portAddressNumber)
	server.RunServer()
}
