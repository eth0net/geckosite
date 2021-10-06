package handlers

import (
	"net/http"

	"github.com/eth0net/geckosite/database"
	"github.com/eth0net/geckosite/database/model"
	"github.com/eth0net/geckosite/templates"
	"github.com/gorilla/mux"
	"gorm.io/gorm/clause"
)

// Animals returns a list of animals with the given order, type and category.
func Animals(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	pageDataMap := map[string]struct {
		Title, Where, Path, Image string
		Species                   *model.Species
		Animals                   []*model.Animal
	}{
		"available": {
			Title: "Available",
			Where: "status = 'For Sale'",
		},
		"holdbacks": {
			Title: "Holdbacks",
			Where: "status = 'Holdback'",
		},
		"personal": {
			Title: "Our Animals",
			Where: "status IN ('Breeder','Future Breeder')",
		},
	}

	pageData := pageDataMap[vars["status"]]
	pageData.Path = r.URL.Path

	var species model.Species
	database.DB.
		Where(&model.Species{Order: vars["order"], Type: vars["type"]}).
		First(&species)
	pageData.Species = &species

	var animals []*model.Animal
	database.DB.
		Preload(clause.Associations).
		Where("species_id = ? AND "+pageData.Where, species.ID).
		Order("name").
		Find(&animals)

	for _, animal := range animals {
		// get image for animal
		images := animal.Images()

		if len(images) == 0 {
			continue
		}

		pageData.Animals = append(pageData.Animals, animal)

		if pageData.Image == "" {
			pageData.Image = images[0]
		}
	}

	templates.Parse("animals").ExecuteTemplate(w, "layout", pageData)
}
