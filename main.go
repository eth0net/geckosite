package main

import (
	"github.com/eth0net/geckosite/database"
	"github.com/eth0net/geckosite/router"
	"github.com/eth0net/geckosite/s3"
)

func main() {
	database.Init()
	s3.Init()
	router.InitAndServe()
}
