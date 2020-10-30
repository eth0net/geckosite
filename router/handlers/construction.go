package handlers

import (
	"html/template"
	"net/http"
)

// Construction returns the under construction page.
func Construction(w http.ResponseWriter, r *http.Request) {
	data := struct{ Title, Path, Message string }{
		"Under Construction",
		r.URL.Path,
		"Page is currently under construction",
	}

	lp, hp := "templates/layout.gohtml", "templates/message.gohtml"
	tmpl := template.Must(template.ParseFiles(lp, hp))
	tmpl.ExecuteTemplate(w, "layout", data)
}
