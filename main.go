package main

import "github.com/ediallocyf/asciiartweb/api"

func main() {
	server := api.NewAPIServer(":8080")
	server.Run()
}
