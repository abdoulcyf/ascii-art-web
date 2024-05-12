package servers

import (
	"net/http"

	"github.com/ediallocyf/asciiartweb/internal/handlers"
)

// Run starts the HTTP server.
func (s *server) Run() {
	mux := http.NewServeMux()

	mux.HandleFunc(homePagePath, makeHTTPHandlerFunc(handlers.MainPageHandler))
	mux.HandleFunc(asciiArtPage, makeHTTPHandlerFunc(handlers.AsciiArtHandler))
	err := http.ListenAndServe(s.listener, mux)
	if err != nil {
		return
	}
}
