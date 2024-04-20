package handlers

import (
	"net/http"
	"text/template"
)

// WriteHTMLResult writes an HTML response for displaying ASCII art result.
func WriteHTMLAscii(w http.ResponseWriter, status int, asciiArt string) error {
	w.Header().Add("Content-type", contentType)
	w.WriteHeader(status)
	tmpl := template.Must(template.ParseFiles(ascciArtTemplateAddress))
	return tmpl.Execute(w, struct{ Banner string }{asciiArt})
}
