package seed

import (
	"log"
	"math/rand"

	"github.com/raziel2244/geckosite/database"
	"github.com/raziel2244/geckosite/database/model"
)

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

	image := &model.Image{
		FileName: "hhicon",
		FileType: "png",
		FilePath: "/static/img/",
	}
	database.DB.Create(image)

	species := &model.Species{
		Name:        "Gargoyle Gecko",
		LatinName:   "Rhacodactylus auriculatus",
		Description: `Carnivorous gecko, around 12cm long with bumped head.`,
	}
	database.DB.Create(species)

	a1 := &model.Animal{
		Reference: "2020/GG1",
		Name:      "GG1",
		Sex:       "Male",
		Status:    "Breeder",
		Species:   species,
		Images:    []*model.Image{image},
	}
	database.DB.Create(a1)

	a2 := &model.Animal{
		Name:    "GG2",
		Sex:     "Female",
		Status:  "Breeder",
		Species: species,
		Images:  []*model.Image{image},
	}
	database.DB.Create(a2)

	a3 := &model.Animal{
		Name:    "GG3",
		Species: species,
		Father:  a1,
		Mother:  a2,
		Images:  []*model.Image{image},
	}
	database.DB.Create(a3)
}
