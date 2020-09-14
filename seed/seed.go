package seed

import (
	"log"
	"math/rand"
	"time"

	"github.com/raziel2244/geckosite/database"
	"github.com/raziel2244/geckosite/database/model"
)

const dateFormat = "2006-01-02"

// Rand seeds math/rand for repeatable results.
func Rand() {
	rand.Seed(1)
}

// Database seeds the database.
func Database() {
	var err error
	err = database.DB.Migrator().DropTable(
		&model.Animal{},
		&model.Species{},
		&model.Trait{},
		&model.Image{},
		"animal_images",
	)
	err = database.DB.AutoMigrate(
		&model.Animal{},
		&model.Species{},
		&model.Trait{},
		&model.Image{},
	)
	if err != nil {
		log.Panicln("Failed to seed DB", err)
	}

	species := &model.Species{
		Name:        "Gargoyle Gecko",
		LatinName:   "Rhacodactylus auriculatus",
		Description: `Carnivorous gecko, around 12cm long with bumped head.`,
		Order:       "geckos",
		Type:        "gargoyle",
	}
	database.DB.Create(species)

	d1, _ := time.Parse(dateFormat, "2020-03-20")
	a1 := &model.Animal{
		Name:        "Flick",
		Description: "Red and black",
		Image:       "/static/img/contact.jpg",
		Species:     species,
		Sex:         "Male",
		DateBought:  &d1,
	}
	database.DB.Create(a1)

	a2 := &model.Animal{
		Name:        "Echo",
		Description: "White and black with orange patches",
		Image:       "/static/img/contact.jpg",
		Species:     species,
		Sex:         "Female",
		DateBought:  &d1,
	}
	database.DB.Create(a2)

	d2, _ := time.Parse(dateFormat, "2020-06-13")
	d3, _ := time.Parse(dateFormat, "2020-07-21")
	a3 := &model.Animal{
		Reference:   "2020/GG1",
		Name:        "Wyvern",
		Description: "Black and orange",
		Image:       "/static/img/contact.jpg",
		Species:     species,
		Status:      "Hold",
		DateLaid:    &d2,
		DateHatched: &d3,
		Father:      a1,
		Mother:      a2,
	}
	database.DB.Create(a3)
}
