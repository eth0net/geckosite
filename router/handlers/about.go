package handlers

import (
	"html/template"
	"net/http"
)

// About returns the about us page.
func About(w http.ResponseWriter, r *http.Request) {
	pageData := struct{ Title, Path string }{"About Us", r.URL.Path}
	lp, hp := "templates/layout.gohtml", "templates/about.gohtml"
	tmpl := template.Must(template.ParseFiles(lp, hp))
	tmpl.ExecuteTemplate(w, "layout", pageData)
}
