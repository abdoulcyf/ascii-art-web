package handlers

import (
	"net/http"
	"text/template"
)

// WriteHTMLResult writes an HTML response for displaying ASCII art result.
func WriteHTMLResult(w http.ResponseWriter, status int, asciiArt string, homePageUrlt string) error {
	w.Header().Add("Content-type", "text/html")
	w.WriteHeader(status)
	tmpl := template.Must(template.ParseFiles(asciiTemplateAdrr))
	return tmpl.Execute(w, struct{ Banner string }{asciiArt})
}
