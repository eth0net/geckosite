package handlers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/raziel2244/geckosite/database"
	"github.com/raziel2244/geckosite/database/model"
)

// Holdbacks returns a list of holdbacks for the given order/type.
func Holdbacks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var species model.Species
	where1 := &model.Species{Order: vars["order"], Type: vars["type"]}
	database.DB.First(&species, where1)

	var animals []*model.Animal
	where2 := &model.Animal{SpeciesID: species.ID, Status: "Holdback"}
	database.DB.Find(&animals, where2)

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
