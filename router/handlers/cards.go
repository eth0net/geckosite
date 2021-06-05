package handlers

import (
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/eth0net/geckosite/database"
	"github.com/eth0net/geckosite/database/model"
	"github.com/eth0net/geckosite/templates"
	"github.com/gorilla/mux"
	"gorm.io/gorm/clause"
)

// Cards returns a page containing cards for child routes.
func Cards(w http.ResponseWriter, r *http.Request) {
	type card struct{ Title, Path, Image string }
	type page struct {
		Title, Path, Image string
		Cards              []card
	}
	type pageMap map[string]page

	pages := map[string]pageMap{
		"geckos": {
			"": {
				Title: "Geckos",
				Cards: []card{
					{"Crested Geckos", "/geckos/crested/", ""},
					{"Gargoyle Geckos", "/geckos/gargoyle/", ""},
					{"Leopard Geckos", "/geckos/leopard/", ""},
				}},
			"crested": {
				Title: "Crested Geckos",
				Cards: []card{
					{"Personal", "/geckos/crested/personal/", ""},
					{"Holdbacks", "/geckos/crested/holdbacks/", ""},
					{"Available", "/geckos/crested/available/", ""},
				}},
			"gargoyle": {
				Title: "Gargoyle Geckos",
				Cards: []card{
					{"Personal", "/geckos/gargoyle/personal/", ""},
					{"Holdbacks", "/geckos/gargoyle/holdbacks/", ""},
					{"Available", "/geckos/gargoyle/available/", ""},
				}},
			"leopard": {
				Title: "Leopard Geckos",
				Cards: []card{
					{"Personal", "/geckos/leopard/personal/", ""},
					{"Holdbacks", "/geckos/leopard/holdbacks/", ""},
					{"Available", "/geckos/leopard/available/", ""},
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

	pageData := pages[o][t]
	pageData.Path = r.URL.Path
	for c, card := range pageData.Cards {
		// set default card image to coming soon
		pageData.Cards[c].Image = "/static/img/coming-soon.jpg"

		splitPath := strings.Split(card.Path, "/")

		var species model.Species
		database.DB.First(&species, &model.Species{Order: o, Type: splitPath[2]})

		var animals []*model.Animal
		if len(splitPath) == 3 {
			// check available
			// check holdbacks
			// check personal

			wheres := []string{
				"status = 'For Sale'",
				"status = 'Holdback'",
				"status IN ('Non-Breeder','Breeder','Future Breeder')",
			}

			for _, where := range wheres {
				database.DB.
					Preload(clause.Associations).
					Where("species_id = ? AND "+where, species.ID).
					Find(&animals)
				if len(animals) == 0 {
					continue
				}
			}
		} else {
			// check category

			wheres := map[string]string{
				"available": "status = 'For Sale'",
				"holdbacks": "status = 'Holdback'",
				"personal":  "status IN ('Non-Breeder','Breeder','Future Breeder')",
			}

			where := "species_id = ?"
			if w, ok := wheres[splitPath[3]]; ok {
				where += " AND " + w
			}
			database.DB.
				Preload(clause.Associations).
				Where(where, species.ID).
				Find(&animals)
		}

		if len(animals) == 0 {
			continue
		}

		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(animals), func(i, j int) {
			animals[i], animals[j] = animals[j], animals[i]
		})

		for _, animal := range animals {
			images := animal.Images()

			// pick random one for card
			if len(images) > 0 {
				pageData.Cards[c].Image = images[0]
				pageData.Image = images[0]
				break
			}
		}
	}

	templates.Parse("cards").ExecuteTemplate(w, "layout", pageData)
}
