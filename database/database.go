package database

import (
	"log"
	"sync"

	"github.com/raziel2244/geckosite/database/model"
	"gorm.io/driver/sqlite"
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
		var err error
		DB, err = gorm.Open(sqlite.Open("store.db"), &gorm.Config{})

		if err != nil {
			log.Panic("failed to connect to database", err)
		}

		DB.AutoMigrate(
			&model.Animal{},
			&model.Species{},
			&model.Trait{},
			&model.Image{},
		)
	})

	return DB
}
