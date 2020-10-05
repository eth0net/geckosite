package handlers

import (
	"html/template"
	"net/http"
)

// NotFound returns the 404 not found page.
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)

	data := struct{ Title, Message string }{
		"Not found",
		"Error 404: Page not found",
	}

	lp, hp := "templates/layout.gohtml", "templates/message.gohtml"
	tmpl := template.Must(template.ParseFiles(lp, hp))
	tmpl.ExecuteTemplate(w, "layout", data)
}
