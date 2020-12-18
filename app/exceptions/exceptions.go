package exceptions

import (
	"cms/pkg/exception"
	"fmt"
	"html/template"
	"net/http"
)

func baseHttpException(w http.ResponseWriter, wantJson bool, statusCode int, message string)  {
	w.WriteHeader(statusCode)
	if wantJson {
		w.Header().Set("Content-type", "application/json")
		fmt.Fprint(w, exception.GetJsonBody(statusCode, message))
	} else {
		file := "./resources/views/public/error.gohtml"
		tmpl, _ := template.ParseFiles(file)
		tmpl.Execute(w, message)
	}
}
