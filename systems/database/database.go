package database

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/eth0net/geckosite/systems/database/model"
)

var (
	// DB stores the database connection.
	DB   *gorm.DB
	once sync.Once
)

// Init opens a database connection and performs migrations.
func Init() *gorm.DB {
	once.Do(initAndMigrate)
	return DB
}

func initAndMigrate() {
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

	err = DB.AutoMigrate(
		&model.Animal{},
		&model.AnimalParent{},
		&model.Contact{},
		&model.Measurement{},
		&model.Species{},
		&model.Trait{},
		&model.Transaction{},
	)

	if err != nil {
		log.Panic("failed to auto migrate database models")
	}

	err = DB.SetupJoinTable(&model.Animal{}, "Parents", &model.AnimalParent{})
	if err != nil {
		log.Panic("failed to setup animal parent join")
	}

	err = DB.SetupJoinTable(&model.Animal{}, "Children", &model.AnimalParent{})
	if err != nil {
		log.Panic("failed to setup animal children join")
	}
}
