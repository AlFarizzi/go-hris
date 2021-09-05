package helper

import (
	"net/http"
	"text/template"
)

func AuthViewParser(w http.ResponseWriter, view string) {
	tmp := template.Must(template.ParseGlob("template/auth/*.html"))
	tmp.ExecuteTemplate(w, view, nil)
}

var partials = template.Must(template.ParseGlob("template/dashboard_partials/*.html"))

func DashboardViewParser(w http.ResponseWriter, view string, path string, data map[string]interface{}) {
	tmp := template.Must(partials.ParseGlob(path))
	tmp.ExecuteTemplate(w, view, data)
}
