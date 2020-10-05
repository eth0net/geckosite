package handlers

import (
	"html/template"
	"net/http"
)

// Home returns the home page.
func Home(w http.ResponseWriter, r *http.Request) {
	lp, hp := "templates/layout.gohtml", "templates/home.gohtml"
	tmpl := template.Must(template.ParseFiles(lp, hp))
	tmpl.ExecuteTemplate(w, "layout", nil)
}
