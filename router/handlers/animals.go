package handlers

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/minio/minio-go/v7"
	"github.com/raziel2244/geckosite/database"
	"github.com/raziel2244/geckosite/database/model"
	"github.com/raziel2244/geckosite/s3"
	"github.com/raziel2244/geckosite/templates"
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

	for _, animal := range animals {
		// get image for animal
		ch := s3.Client.ListObjects(
			context.Background(),
			species.Order,
			minio.ListObjectsOptions{
				Prefix:    species.Type + "/" + animal.ID.String(),
				Recursive: true,
				MaxKeys:   1,
			},
		)
		// store into the animal struct
		for object := range ch {
			path := "/s3/" + species.Order + "/" + object.Key
			animal.Images = append(animal.Images, path)
		}

		if pageData.Image == "" && len(animal.Images) > 0 {
			pageData.Image = animal.Images[0]
		}
	}

	pageData.Animals = animals

	templates.Parse("animals").ExecuteTemplate(w, "layout", pageData)
}
