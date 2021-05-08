package database

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/raziel2244/geckosite/database/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// DB stores the database connection.
	DB   *gorm.DB
	once sync.Once
)

// Init opens a database connection and performs migrations.
func Init() *gorm.DB {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%v dbname=%v user=%v password=%v",
			os.Getenv("DB_HOST"), os.Getenv("DB_NAME"),
			os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		)

		var err error
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Panic("failed to connect to database")
		}

		DB.AutoMigrate(
			&model.Animal{},
			&model.AnimalParent{},
			&model.Contact{},
			&model.Measurement{},
			&model.Species{},
			&model.Trait{},
		)

		DB.SetupJoinTable(&model.Animal{}, "Parents", &model.AnimalParent{})
		DB.SetupJoinTable(&model.Animal{}, "Children", &model.AnimalParent{})
	})

	return DB
}
