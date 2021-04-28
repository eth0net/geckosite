package handlers

import (
	"net/http"

	"github.com/raziel2244/geckosite/templates"
)

// NotFound returns the 404 not found page.
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)

	data := struct{ Title, Path, Message string }{
		"Not found",
		r.URL.Path,
		"Error 404: Page not found",
	}

	templates.Parse("message").ExecuteTemplate(w, "layout", data)
}
