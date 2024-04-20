package main

import (
	"github.com/ediallocyf/asciiartweb/internal/handlers"
)

func main() {
	server := handlers.NewAPIServer(":8080")
	server.RunServer()
}
