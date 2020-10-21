package main

import (
	"github.com/raziel2244/geckosite/database"
	"github.com/raziel2244/geckosite/mail"
	"github.com/raziel2244/geckosite/router"
	"github.com/raziel2244/geckosite/s3"
)

func main() {
	database.Init()
	s3.Init()
	mail.Init()
	router.InitAndServe()
}
