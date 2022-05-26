package web

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"

	"github.com/eth0net/geckosite/systems/database/model"
	"github.com/eth0net/geckosite/systems/templates"
)

// Animals returns a list of animals with the given order, type and category.
func (s service) Animals(c *gin.Context) {
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

	pageData := pageDataMap[c.Param("id")]
	pageData.Path = c.Request.URL.Path

	var species model.Species
	s.db.
		Where(&model.Species{Order: c.Param("order"), Type: c.Param("type")}).
		First(&species)
	pageData.Species = &species

	var animals []*model.Animal
	s.db.
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

	templates.Parse("animals").ExecuteTemplate(c.Writer, "layout", pageData)
}
