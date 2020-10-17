package handlers

import (
	"context"
	"html/template"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/minio/minio-go/v7"
	"github.com/raziel2244/geckosite/database"
	"github.com/raziel2244/geckosite/database/model"
	"github.com/raziel2244/geckosite/s3"
)

// Cards returns a page containing cards for child routes.
func Cards(w http.ResponseWriter, r *http.Request) {
	type card struct{ Title, Path, Image string }
	type page struct {
		Title string
		Cards []card
	}
	type pageMap map[string]page

	pages := map[string]pageMap{
		"geckos": {
			"": {"Geckos", []card{
				{"Crested Geckos", "/geckos/crested", ""},
				{"Gargoyle Geckos", "/geckos/gargoyle", ""},
				{"Leopard Geckos", "/geckos/leopard", ""},
			}},
			"crested": {"Crested Geckos", []card{
				{"Personal", "/geckos/crested/personal", ""},
				{"Holdbacks", "/geckos/crested/holdbacks", ""},
				{"For Sale", "/geckos/crested/for-sale", ""},
			}},
			"gargoyle": {"Gargoyle Geckos", []card{
				{"Personal", "/geckos/gargoyle/personal", ""},
				{"Holdbacks", "/geckos/gargoyle/holdbacks", ""},
				{"For Sale", "/geckos/gargoyle/for-sale", ""},
			}},
			"leopard": {"Leopard Geckos", []card{
				{"Personal", "/geckos/leopard/personal", ""},
				{"Holdbacks", "/geckos/leopard/holdbacks", ""},
				{"For Sale", "/geckos/leopard/for-sale", ""},
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

	rand.Seed(time.Now().Unix())

	pageData := pages[o][t]
	for c, card := range pageData.Cards {
		splitPath := strings.Split(card.Path, "/")

		var species model.Species
		database.DB.First(&species, &model.Species{Order: o, Type: splitPath[2]})

		var animals []*model.Animal
		if len(splitPath) == 3 {
			// check for sale
			// check holdbacks
			// check personal

			wheres := []string{
				"status = 'For Sale'",
				"status = 'Holdback'",
				"status IN ('Non-Breeder','Breeder','Future Breeder')",
			}

			for _, where := range wheres {
				database.DB.Where("species_id = ? AND "+where, species.ID.String()).Find(&animals)
				if len(animals) > 0 {
					break
				}
			}
		} else {
			// check category

			wheres := map[string]string{
				"for-sale":  "status = 'For Sale'",
				"holdbacks": "status = 'Holdback'",
				"personal":  "status IN ('Non-Breeder','Breeder','Future Breeder')",
			}

			where := wheres[splitPath[3]]
			database.DB.Find(&animals, "species_id = ? AND "+where, species.ID)
		}

		if len(animals) == 0 {
			pageData.Cards[c].Image = "/static/img/coming-soon.jpg"
			continue
		}

		for _, animal := range animals {
			// get images for animal
			var images []string
			ch := s3.Client.ListObjects(
				context.Background(),
				o,
				minio.ListObjectsOptions{
					Prefix:    splitPath[2] + "/" + animal.ID.String(),
					Recursive: true,
				},
			)
			for object := range ch {
				path := "/s3/" + o + "/" + object.Key
				images = append(images, path)
			}

			// pick random one for card
			if len(images) > 0 {
				pageData.Cards[c].Image = images[rand.Intn(len(images))]
				break
			}
		}
	}

	lp, hp := "templates/layout.gohtml", "templates/cards.gohtml"
	tmpl := template.Must(template.ParseFiles(lp, hp))
	tmpl.ExecuteTemplate(w, "layout", pages[o][t])
}
