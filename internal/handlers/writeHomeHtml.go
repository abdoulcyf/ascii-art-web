package handlers

import (
	"net/http"
	"text/template"
)

func WriteHTML(w http.ResponseWriter, status int) error {
	w.Header().Set("Content-Type", "text/html")
	//w.WriteHeader(status)
	tmpl := template.Must(template.ParseFiles(homeTemplateAdrr))
	return tmpl.Execute(w, nil)
}
