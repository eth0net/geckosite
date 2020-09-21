package router

import (
	"html/template"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
	"github.com/raziel2244/geckosite/database"
	"github.com/raziel2244/geckosite/database/model"
	"gorm.io/gorm/clause"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	lp, hp := "templates/layout.gohtml", "templates/404.gohtml"
	tmpl := template.Must(template.ParseFiles(lp, hp))
	tmpl.ExecuteTemplate(w, "layout", nil)
}

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
				{"Our Animals", "/static/img/hhicon.png", "/geckos/crested/our-animals"},
				{"Holdbacks", "/static/img/hhicon.png", "/geckos/crested/holdbacks"},
				{"For Sale", "/static/img/hhicon.png", "/geckos/crested/for-sale"},
			}},
			"gargoyle": {"Gargoyle Geckos", []card{
				{"Our Animals", "/static/img/hhicon.png", "/geckos/gargoyle/our-animals"},
				{"Holdbacks", "/static/img/hhicon.png", "/geckos/gargoyle/holdbacks"},
				{"For Sale", "/static/img/hhicon.png", "/geckos/gargoyle/for-sale"},
			}},
			"leopard": {"Leopard Geckos", []card{
				{"Our Animals", "/static/img/hhicon.png", "/geckos/leopard/our-animals"},
				{"Holdbacks", "/static/img/hhicon.png", "/geckos/leopard/holdbacks"},
				{"For Sale", "/static/img/hhicon.png", "/geckos/leopard/for-sale"},
			}},
		},
	}

	vars := mux.Vars(r)
	o, t := vars["order"], vars["type"]

	// No pages for order, return 404.
	if _, ok := pages[o]; !ok {
		notFound(w, r)
		return
	}

	// No pages for order/type, return 404.
	if _, ok := pages[o][t]; !ok {
		notFound(w, r)
		return
	}

	lp, hp := "templates/layout.gohtml", "templates/cards.gohtml"
	tmpl := template.Must(template.ParseFiles(lp, hp))
	tmpl.ExecuteTemplate(w, "layout", pages[o][t])
}

func ourAnimals(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var species model.Species
	where1 := &model.Species{Order: vars["order"], Type: vars["type"]}
	database.DB.First(&species, where1)

	var animals []*model.Animal
	database.DB.Preload("Images").
		Where("species_id = ? AND status IN ('Non-Breeder','Breeder','Future Breeder')", species.ID).
		Order("name ASC").
		Find(&animals)

	data := struct {
		Title   string
		Animals []*model.Animal
	}{
		Title:   "Our Animals | " + species.Name + "s",
		Animals: animals,
	}

	lp, hp := "templates/layout.gohtml", "templates/animals.gohtml"
	tmpl := template.Must(template.ParseFiles(lp, hp))
	tmpl.ExecuteTemplate(w, "layout", data)
}

func holdbacks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var species model.Species
	where1 := &model.Species{Order: vars["order"], Type: vars["type"]}
	database.DB.First(&species, where1)

	var animals []*model.Animal
	where2 := &model.Animal{SpeciesID: species.ID, Status: "Holdback"}
	database.DB.Preload("Images").Find(&animals, where2)

	data := struct {
		Title   string
		Animals []*model.Animal
	}{
		Title:   "Holdbacks | " + species.Name + "s",
		Animals: animals,
	}

	lp, hp := "templates/layout.gohtml", "templates/animals.gohtml"
	tmpl := template.Must(template.ParseFiles(lp, hp))
	tmpl.ExecuteTemplate(w, "layout", data)
}

func forSale(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var species model.Species
	where1 := &model.Species{Order: vars["order"], Type: vars["type"]}
	database.DB.First(&species, where1)

	var animals []*model.Animal
	where2 := &model.Animal{SpeciesID: species.ID, Status: "For Sale"}
	database.DB.Preload("Images").Find(&animals, where2)

	data := struct {
		Title   string
		Animals []*model.Animal
	}{
		Title:   "For Sale | " + species.Name + "s",
		Animals: animals,
	}

	lp, hp := "templates/layout.gohtml", "templates/animals.gohtml"
	tmpl := template.Must(template.ParseFiles(lp, hp))
	tmpl.ExecuteTemplate(w, "layout", data)
}

func animal(w http.ResponseWriter, r *http.Request) {
	var data model.Animal
	database.DB.Model(&model.Animal{}).
		Preload(clause.Associations).
		Where("id = ?", mux.Vars(r)["id"]).
		First(&data)

	if reflect.ValueOf(data).IsZero() {
		notFound(w, r)
		return
	}

	lp, hp := "templates/layout.gohtml", "templates/animal.gohtml"
	tmpl := template.Must(template.ParseFiles(lp, hp))
	tmpl.ExecuteTemplate(w, "layout", data)
}
