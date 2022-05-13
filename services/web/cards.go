package web

import (
	"math/rand"
	"strings"
	"time"

	"github.com/eth0net/geckosite/database/model"
	"github.com/eth0net/geckosite/templates"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

// Cards returns a page containing cards for child routes.
func (s service) Cards(c *gin.Context) {
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

	paramOrder, paramType := c.Param("order"), c.Param("type")

	// No pages for order, return 404.
	if _, ok := pages[paramOrder]; !ok {
		s.NotFound(c)
		return
	}

	// No pages for order/type, return 404.
	if _, ok := pages[paramOrder][paramType]; !ok {
		s.NotFound(c)
		return
	}

	pageData := pages[paramOrder][paramType]
	pageData.Path = c.Request.URL.Path
	for c, card := range pageData.Cards {
		// set default card image to coming soon
		pageData.Cards[c].Image = "/static/img/coming-soon.jpg"

		splitPath := strings.Split(card.Path, "/")

		var species model.Species
		s.db.First(&species, &model.Species{Order: paramOrder, Type: splitPath[2]})

		var animals []*model.Animal
		if len(splitPath) == 3 {
			// check available
			// check holdbacks
			// check personal

			wheres := []string{
				"status = 'For Sale'",
				"status = 'Holdback'",
				"status IN ('Breeder','Future Breeder')",
			}

			for _, where := range wheres {
				s.db.
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
				"personal":  "status IN ('Breeder','Future Breeder')",
			}

			where := "species_id = ?"
			if w, ok := wheres[splitPath[3]]; ok {
				where += " AND " + w
			}
			s.db.
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

	templates.Parse("cards").ExecuteTemplate(c.Writer, "layout", pageData)
}
