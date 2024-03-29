package handlers

import (
	"context"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/eth0net/geckosite/database"
	"github.com/eth0net/geckosite/database/model"
	"github.com/eth0net/geckosite/s3"
	"github.com/eth0net/geckosite/templates"
	"github.com/minio/minio-go/v7"
)

// Home returns the home page.
func Home(w http.ResponseWriter, r *http.Request) {
	type card struct{ Title, Path, Image string }
	type page struct {
		Title, Path string
		Cards       []card
		Count       int64
	}

	data := page{
		Title: "Home",
		Path:  r.URL.Path,
		Cards: []card{
			{"Crested Geckos", "/geckos/crested/", ""},
			{"Gargoyle Geckos", "/geckos/gargoyle/", ""},
			{"Leopard Geckos", "/geckos/leopard/", ""},
		},
	}

	rand.Seed(time.Now().Unix())

	for c, card := range data.Cards {
		// set default card image to coming soon
		data.Cards[c].Image = "/static/img/coming-soon.jpg"

		splitPath := strings.Split(card.Path, "/")

		var species = &model.Species{}
		database.DB.
			Where(&model.Species{Order: splitPath[1], Type: splitPath[2]}).
			First(species)

		if species == nil {
			continue
		}

		var animals []*model.Animal
		// check for sale
		// check holdbacks
		// check personal

		wheres := []string{
			"status = 'For Sale'",
			"status = 'Holdback'",
			"status IN ('Breeder','Future Breeder')",
		}

		for _, where := range wheres {
			database.DB.
				Where("species_id = ? AND "+where, species.ID).
				Find(&animals)
			if len(animals) == 0 {
				continue
			}

			rand.Seed(time.Now().UnixNano())
			rand.Shuffle(len(animals), func(i, j int) {
				animals[i], animals[j] = animals[j], animals[i]
			})

			var image string
			for _, animal := range animals {
				// get images for animal
				var images []string
				ch := s3.Client.ListObjects(
					context.Background(),
					splitPath[1],
					minio.ListObjectsOptions{
						Prefix:    splitPath[2] + "/" + animal.ID.String(),
						Recursive: true,
					},
				)
				for object := range ch {
					path := "/s3/" + splitPath[1] + "/" + object.Key
					images = append(images, path)
				}

				// pick random one for card
				if len(images) > 0 {
					image = images[0]
					break
				}
			}

			if len(image) > 0 {
				data.Cards[c].Image = image
				break
			}
		}
	}

	database.DB.Table("animals").Count(&data.Count)

	templates.Parse("home").ExecuteTemplate(w, "layout", data)
}
