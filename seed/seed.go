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

	date1, _ := time.Parse(dateFormat, "2020-03-20")
	animal1 := &model.Animal{
		Name:        "Flick",
		Description: "Red and black",
		Images: []*model.Image{
			{
				FileName: "Flick-1150064.jpg",
				FilePath: "/static/img/geckos/gargoyle/Flick",
			},
			{
				FileName: "Flick-1150056.jpg",
				FilePath: "/static/img/geckos/gargoyle/Flick",
			},
			{
				FileName: "Flick-1150005.jpg",
				FilePath: "/static/img/geckos/gargoyle/Flick",
			},
			{
				FileName: "Flick-1150039.jpg",
				FilePath: "/static/img/geckos/gargoyle/Flick",
			},
		},
		Species:    species,
		Sex:        "Male",
		Status:     "Breeder",
		DateBought: &date1,
	}
	database.DB.Create(animal1)

	animal2 := &model.Animal{
		Name:        "Echo",
		Description: "White and black with orange patches",
		Images: []*model.Image{
			{
				FilePath: "/static/img/geckos/gargoyle/Echo",
				FileName: "P1140751.jpg",
			},
			{
				FilePath: "/static/img/geckos/gargoyle/Echo",
				FileName: "P1140764.jpg",
			},
			{
				FilePath: "/static/img/geckos/gargoyle/Echo",
				FileName: "P1140764.jpg",
			},
			{
				FilePath: "/static/img/geckos/gargoyle/Echo",
				FileName: "P1140764.jpg",
			},
		},
		Species:    species,
		Sex:        "Female",
		Status:     "Breeder",
		DateBought: &date1,
	}
	database.DB.Create(animal2)

	date2, _ := time.Parse(dateFormat, "2020-06-13")
	date3, _ := time.Parse(dateFormat, "2020-07-21")
	animal3 := &model.Animal{
		Reference:   "2020/GG1",
		Name:        "Wyvern",
		Description: "Black and orange",
		Species:     species,
		Status:      "Holdback",
		DateLaid:    &date2,
		DateHatched: &date3,
		Father:      animal1,
		Mother:      animal2,
	}
	database.DB.Create(animal3)
}
