package handlers

import (
	"html/template"
	"net/http"
)

// About returns the about us page.
func About(w http.ResponseWriter, r *http.Request) {
	lp, hp := "templates/layout.gohtml", "templates/about.gohtml"
	tmpl := template.Must(template.ParseFiles(lp, hp))
	tmpl.ExecuteTemplate(w, "layout", nil)
}
