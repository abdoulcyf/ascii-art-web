package main

import "github.com/ediallocyf/asciiartweb/api"

func main() {
	server := api.NewAPIServer(":2030")
	server.Run()
}
