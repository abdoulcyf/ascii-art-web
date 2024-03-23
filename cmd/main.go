package main

import (
	handlers "main/handlers"
	templates "main/templates"
	"net/http"
)

func main() {
	handlers.StandardHandler()
	handlers.ShadowHandler()
	handlers.ThinkerToyHandler()

	//Define HTTP routes
	http.HandleFunc("/", templates.MainPageHandler)
	http.HandleFunc("/ascii-art", templates.AsciiArtHandler)

	// Start the HTTP server
	http.ListenAndServe(":8080", nil)
}
