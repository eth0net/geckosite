package handlers

import (
	"html/template"
	"net/http"
)

// Construction returns the under contruction page.
func Construction(w http.ResponseWriter, r *http.Request) {
	data := struct{ Title, Message string }{
		"Under Construction",
		"Page is currently under construction",
	}

	lp, hp := "templates/layout.gohtml", "templates/message.gohtml"
	tmpl := template.Must(template.ParseFiles(lp, hp))
	tmpl.ExecuteTemplate(w, "layout", data)
}
