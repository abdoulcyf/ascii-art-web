package servers

import (
	"net/http"
	"os"
)

func (s *server) fileServer(staticFilesAddress string) error {

	// Check if the directory exists
	if _, err := os.Stat(staticFilesAddress); os.IsNotExist(err) {
		errMsg = "=fileServer===staticFilesAddress==" + err.Error()
		logger.Error(errMsg)
	}

	fileServer := http.FileServer(http.Dir(staticFilesAddress))

	// Register the file server handler with the "/static/" URL pattern
	http.Handle(baseStaticDir, http.StripPrefix(baseStaticDir, fileServer))

	return nil
}
