package handlers

import (
	"net/http"
	"text/template"
)

// WriteHTMLResult writes an HTML response for displaying ASCII art result.
func WriteHTMLAscii(w http.ResponseWriter, status int, asciiArt string) error {
	w.Header().Add("Content-type", contentType)
	w.WriteHeader(status)

	// Parse content to html template
	tmpl := template.Must(template.ParseFiles(ascciArtTemplateAddress))

	// Excuting ascii art template
	errAscciTemp := tmpl.Execute(w, struct{ Banner string }{asciiArt})
	if errAscciTemp != nil {
		errMsg = "----<-WriteHTMLAscii------<--ParseFiles----"
		logMsg = "Error executing Asccii art template" + errAscciTemp.Error()
		logger.Error(logMsg)
	} else {
		logger.Info(" Asccii Art Template parsed successfully")
	}

	return nil
}
