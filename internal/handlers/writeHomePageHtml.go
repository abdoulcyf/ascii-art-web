package handlers

import (
	"net/http"
	"text/template"
)

// WriteHTML writes an HTML response with the provided status code and content.
func writeHomePageContent(w http.ResponseWriter, status int) error {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(status)
	tmpl := template.Must(template.ParseFiles(hopePageTemplateAdrress))
	return tmpl.Execute(w, nil)
}
