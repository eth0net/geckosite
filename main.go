package main

import (
	"github.com/raziel2244/geckosite/database"
	"github.com/raziel2244/geckosite/router"
	"github.com/raziel2244/geckosite/seed"
)

func main() {
	seed.Rand()

	database.Init()
	seed.Database()

	router.InitAndServe()
}
