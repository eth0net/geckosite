package handlers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

// Cards returns a page containing cards for child routes.
func Cards(w http.ResponseWriter, r *http.Request) {
	type card struct{ Title, Image, Link string }
	type page struct {
		Title string
		Cards []card
	}
	type pageMap map[string]page

	pages := map[string]pageMap{
		"geckos": {
			"": {"Geckos", []card{
				{"Crested Geckos", "/static/img/hhicon.png", "/geckos/crested"},
				{"Gargoyle Geckos", "/static/img/hhicon.png", "/geckos/gargoyle"},
				{"Leopard Geckos", "/static/img/hhicon.png", "/geckos/leopard"},
			}},
			"crested": {"Crested Geckos", []card{
				{"Personal", "/static/img/hhicon.png", "/geckos/crested/personal"},
				{"Holdbacks", "/static/img/hhicon.png", "/geckos/crested/holdbacks"},
				{"For Sale", "/static/img/hhicon.png", "/geckos/crested/for-sale"},
			}},
			"gargoyle": {"Gargoyle Geckos", []card{
				{"Personal", "/static/img/hhicon.png", "/geckos/gargoyle/personal"},
				{"Holdbacks", "/static/img/hhicon.png", "/geckos/gargoyle/holdbacks"},
				{"For Sale", "/static/img/hhicon.png", "/geckos/gargoyle/for-sale"},
			}},
			"leopard": {"Leopard Geckos", []card{
				{"Personal", "/static/img/hhicon.png", "/geckos/leopard/personal"},
				{"Holdbacks", "/static/img/hhicon.png", "/geckos/leopard/holdbacks"},
				{"For Sale", "/static/img/hhicon.png", "/geckos/leopard/for-sale"},
			}},
		},
	}

	vars := mux.Vars(r)
	o, t := vars["order"], vars["type"]

	// No pages for order, return 404.
	if _, ok := pages[o]; !ok {
		NotFound(w, r)
		return
	}

	// No pages for order/type, return 404.
	if _, ok := pages[o][t]; !ok {
		NotFound(w, r)
		return
	}

	lp, hp := "templates/layout.gohtml", "templates/cards.gohtml"
	tmpl := template.Must(template.ParseFiles(lp, hp))
	tmpl.ExecuteTemplate(w, "layout", pages[o][t])
}
