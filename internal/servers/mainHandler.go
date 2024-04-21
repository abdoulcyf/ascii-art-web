package servers

import (
	//"errors"
	//"log"
	"net/http"

	"github.com/ediallocyf/asciiartweb/internal/handlers"
)

// Run starts the HTTP server.
func MainHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.URL.Path == homePagePath && r.Method == http.MethodGet:
		errHomeHandler := handlers.HomeHandler(w, r)
		if errHomeHandler != nil {
			errMsg = "----<-MainHandler------<--HomeHandler----" + errHomeHandler.Error()
			logMsg = "Error in loading home page"
			logger.Error(logMsg + errMsg)
			return
		} else {
			logger.Info(" Home page loaded successfully")
		}
	case r.URL.Path == ascciArtPagePath && r.Method == http.MethodPost:
		errAsciiArt := handlers.GenerateAsciiArtHandler(w, r)
		if errAsciiArt != nil {
			errMsg = "----<-MainHandler------<--AssertArtWebHandler----" + errAsciiArt.Error()
			logMsg = "Error in loading arscii art page"
			logger.Error(logMsg + errMsg)
			return
		}
	default:
		w.WriteHeader(http.StatusNotExtended)
		logger.Info("Ascii art page loaded successfully")
	}

}
