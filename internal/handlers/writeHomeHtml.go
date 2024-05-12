package handlers

import (
	"net/http"
	"text/template"
)

type Data struct {
	ActionElement string
}

func WriteHTML(w http.ResponseWriter, status int) error {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(status)

	data := Data{
		ActionElement: actionPath,
	}
	tmpl := template.Must(template.ParseFiles(homeTemplateAdrr))
	return tmpl.Execute(w, data)
}
