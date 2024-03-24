package main

import (
	handlers "main/handlers"
	"net/http"
)

func main() {
	handlers.StandardHandler()
	handlers.ShadowHandler()
	handlers.ThinkerToyHandler()

	//Define HTTP routes
	http.HandleFunc("/", handlers.MainPageHandler)
	http.HandleFunc("/ascii-art", handlers.AsciiArtHandler)

	// Start the HTTP server
	http.ListenAndServe(":8080", nil)
}
