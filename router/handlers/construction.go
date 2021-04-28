package handlers

import (
	"net/http"

	"github.com/raziel2244/geckosite/templates"
)

// Construction returns the under construction page.
func Construction(w http.ResponseWriter, r *http.Request) {
	data := struct{ Title, Path, Message string }{
		"Under Construction",
		r.URL.Path,
		"Page is currently under construction",
	}

	templates.Parse("message").ExecuteTemplate(w, "layout", data)
}
