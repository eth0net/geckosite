package router

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/raziel2244/geckosite/database"
	"github.com/raziel2244/geckosite/database/model"
	"gorm.io/gorm/clause"
)

func home(w http.ResponseWriter, r *http.Request) {
	lp, hp := "templates/layout.gohtml", "templates/home.gohtml"
	tmpl := template.Must(template.ParseFiles(lp, hp))
	tmpl.ExecuteTemplate(w, "layout", nil)
}

func about(w http.ResponseWriter, r *http.Request) {
	lp, hp := "templates/layout.gohtml", "templates/about.gohtml"
	tmpl := template.Must(template.ParseFiles(lp, hp))
	tmpl.ExecuteTemplate(w, "layout", nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
	lp, hp := "templates/layout.gohtml", "templates/contact.gohtml"
	tmpl := template.Must(template.ParseFiles(lp, hp))
	tmpl.ExecuteTemplate(w, "layout", nil)
}

func cards(w http.ResponseWriter, r *http.Request) {
	type card struct{ Title, Image, Link string }

	var cards = map[string]struct {
		Title string
		Cards []card
	}{
		"geckos": {"Geckos", []card{
			{"Crested Geckos", "/static/img/hhicon.png", "/geckos/crested"},
			{"Gargoyle Geckos", "/static/img/hhicon.png", "/geckos/gargoyle"},
			{"Leopard Geckos", "/static/img/hhicon.png", "/geckos/leopard"},
		}},
		"crested": {"Crested Geckos", []card{
			{"Pets", "/static/img/hhicon.png", "/geckos/crested/pets"},
			{"Breeders", "/static/img/hhicon.png", "/geckos/crested/breeders"},
			{"For Sale", "/static/img/hhicon.png", "/geckos/crested/sale"},
		}},
		"gargoyle": {"Gargoyle Geckos", []card{
			{"Pets", "/static/img/hhicon.png", "/geckos/gargoyle/pets"},
			{"Breeders", "/static/img/hhicon.png", "/geckos/gargoyle/breeders"},
			{"For Sale", "/static/img/hhicon.png", "/geckos/gargoyle/sale"},
		}},
		"leopard": {"Leopard Geckos", []card{
			{"Pets", "/static/img/hhicon.png", "/geckos/leopard/pets"},
			{"Breeders", "/static/img/hhicon.png", "/geckos/leopard/breeders"},
			{"For Sale", "/static/img/hhicon.png", "/geckos/leopard/sale"},
		}},
	}

	data := cards[r.URL.Path[1:]]
	if t := mux.Vars(r)["type"]; t != "" {
		data = cards[t]
	}

	lp, hp := "templates/layout.gohtml", "templates/cards.gohtml"
	tmpl := template.Must(template.ParseFiles(lp, hp))
	tmpl.ExecuteTemplate(w, "layout", data)
}

func animals(w http.ResponseWriter, r *http.Request) {
	// data := struct{}{}

}

func animal(w http.ResponseWriter, r *http.Request) {
	var data model.Animal
	database.DB.Preload(clause.Associations).First(&data, "id = ?", mux.Vars(r)["id"])

	lp, hp := "templates/layout.gohtml", "templates/animal.gohtml"
	tmpl := template.Must(template.ParseFiles(lp, hp))
	tmpl.ExecuteTemplate(w, "layout", data)
}
