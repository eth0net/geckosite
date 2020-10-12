package main

import (
	"github.com/raziel2244/geckosite/database"
	"github.com/raziel2244/geckosite/router"
	"github.com/raziel2244/geckosite/s3"
	"github.com/raziel2244/geckosite/seed"
)

func main() {
	seed.Rand()

	database.Init()
	// seed.Database()

	s3.Init()

	router.InitAndServe()
}
