package servers

import (
	//"fmt"
	"fmt"
	"net/http"

	"github.com/ediallocyf/asciiartweb/internal/handlers"
)

// Run starts the HTTP server.
func (s *server) Run() error {

	//Call serveStaticFiles function to get the static files handler
	errFileServer := s.serveStaticFiles(staticFilesAddress)

	if errFileServer != nil {
		fmt.Printf("file directory does not exit %s", errFileServer)
		return errFileServer
	}

	//----------------------------
	http.HandleFunc(homePagePath, makeHTTPHandlerFunc(handlers.MainPageHandler))
	http.HandleFunc(asciiArtPage, makeHTTPHandlerFunc(handlers.AsciiArtHandler))
	err := http.ListenAndServe(s.listener, nil)
	if err != nil {
		return err
	}
	return nil
}
