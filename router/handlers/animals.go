package handlers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/raziel2244/geckosite/database"
	"github.com/raziel2244/geckosite/database/model"
)

// Animals returns a list of animals with the given order, type and category.
func Animals(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	pageDataMap := map[string]struct {
		Title, Where, Path, Image string
		Species                   *model.Species
		Animals                   []*model.Animal
	}{
		"for-sale": {
			Title: "For Sale",
			Where: "status = 'For Sale'",
		},
		"holdbacks": {
			Title: "Holdbacks",
			Where: "status = 'Holdback'",
		},
		"personal": {
			Title: "Our Animals",
			Where: "status IN ('Non-Breeder','Breeder','Future Breeder')",
		},
	}

	if _, found := pageDataMap[vars["id"]]; !found {
		Animal(w, r)
		return
	}

	pageData := pageDataMap[vars["id"]]
	pageData.Path = r.URL.Path

	var species model.Species
	database.DB.
		Where(&model.Species{Order: vars["order"], Type: vars["type"]}).
		First(&species)
	pageData.Species = &species

	var animals []*model.Animal
	database.DB.
		Where("species_id = ? AND "+pageData.Where, species.ID).
		Order("name").
		Find(&animals)

	// Load images for all animals on the page.
	for _, animal := range animals {
		images := animal.LoadImages()

		// Set page image to first image found.
		if pageData.Image == "" && len(images) > 0 {
			pageData.Image = images[0]
		}
	}

	pageData.Animals = animals

	lp, hp := "templates/layout.gohtml", "templates/animals.gohtml"
	tmpl := template.Must(template.ParseFiles(lp, hp))
	tmpl.ExecuteTemplate(w, "layout", pageData)
}
