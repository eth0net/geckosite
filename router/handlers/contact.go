package handlers

import (
	"html/template"
	"net/http"
)

// Contact returns the contact us page.
func Contact(w http.ResponseWriter, r *http.Request) {
	lp, hp := "templates/layout.gohtml", "templates/contact.gohtml"
	tmpl := template.Must(template.ParseFiles(lp, hp))
	tmpl.ExecuteTemplate(w, "layout", nil)
}
