package handlers

import (
	"net/http"
	"text/template"
)

func ErrorPage(w http.ResponseWriter, errMsg string) {
	funcName := "ErrorPage"
	// Define the template
	t := template.Must(template.ParseFiles(errorTemplAdrr))
	data := struct {
		Msg string
		Url string
	}{
		Msg: errMsg,
		Url: homePagePath,
	}
	err := t.Execute(w, data)
	opName := "Execute-ErrorPage"
	opDes := "Execute ErrorPage"
	if err != nil {
		logger.Error(funcName, opName, err, opDes)
		return
	}
}
