package handlers

import (
	"net/http"
	"text/template"
)

// Render the main page using a template
func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../templates/main.html")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, nil)
}
