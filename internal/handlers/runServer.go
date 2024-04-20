package handlers

import (
	"net/http"
)

// Run starts the HTTP server.
func (s *Server) RunServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", makeHTTPHandlerFunc(s.MainPageHandler))
	mux.HandleFunc("/ascii-art", makeHTTPHandlerFunc(s.assertArtWebHandler))
	err := http.ListenAndServe(s.listener, mux)
	if err != nil {
		return
	}
}
