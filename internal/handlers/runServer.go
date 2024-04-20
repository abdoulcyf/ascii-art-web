package handlers

import (
	"log"
	"net/http"
)

// Run starts the HTTP server.
func (s *Server) RunServer() {
	mux := http.NewServeMux()

	mux.HandleFunc(homePagePath, makeHTTPHandlerFunc(s.HomeHandler))
	mux.HandleFunc(ascciArtPagePath, makeHTTPHandlerFunc(s.assertArtWebHandler))
	err := http.ListenAndServe(s.listener, mux)
	if err != nil {
		log.Fatal("Failed to start the server")
	} else {
		log.Println("Server is running at port 8080")
	}
}
