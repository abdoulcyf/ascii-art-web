package servers

import (
	"fmt"
	"net/http"
	"os"
)

func (s *server) serveStaticFiles(staticFilesAddress  string) error{
	// Check if the directory exists
	if _, err := os.Stat(staticFilesAddress ); os.IsNotExist(err) {
		return  fmt.Errorf("static directory %s does not exist", staticFilesAddress )
	}
	// Create a file server to serve static files from the "/static/" directory
	fileServer := http.FileServer(http.Dir(staticFilesAddress ))

	
	// Register the file server handler with the "/static/" URL pattern
	http.Handle(baseStaticDir, http.StripPrefix(baseStaticDir, fileServer))

	// Return the ServeMux, which now handles static file requests
	return  nil
}
