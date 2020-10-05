package handlers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/raziel2244/geckosite/database"
	"github.com/raziel2244/geckosite/database/model"
)

// Personal returns a list of personal animals with the given order/type.
func Personal(w http.ResponseWriter, r *http.Request) {
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
