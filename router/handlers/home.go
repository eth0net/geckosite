package handlers

import (
	"html/template"
	"net/http"
)

// Home returns the home page.
func Home(w http.ResponseWriter, r *http.Request) {
	type card struct{ Title, Image, Link string }
	type page struct {
		Cards []card
	}

	data := page{
		Cards: []card{
			{"Crested Geckos", "/static/img/hhicon.png", "/geckos/crested"},
			{"Gargoyle Geckos", "/static/img/hhicon.png", "/geckos/gargoyle"},
			{"Leopard Geckos", "/static/img/hhicon.png", "/geckos/leopard"},
		},
	}

	lp, hp := "templates/layout.gohtml", "templates/home.gohtml"
	tmpl := template.Must(template.ParseFiles(lp, hp))
	tmpl.ExecuteTemplate(w, "layout", data)
}
