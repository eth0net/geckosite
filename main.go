package main

import (
	"github.com/eth0net/geckosite/database"
	"github.com/eth0net/geckosite/s3"
	"github.com/eth0net/geckosite/services/web"
	"github.com/eth0net/geckosite/static"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	d := database.Init()
	s := s3.Init()

	webService := web.NewService(d, s)

	r := gin.Default()

	r.GET("/s3/:bucket/*path", webService.S3)

	r.StaticFS("/static/", http.FS(static.FS))
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/static/icon/favicon.ico")
	})
	r.GET("/sitemap.xml", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/static/sitemap.xml")
	})

	r.GET("/", webService.Home)
	r.GET("/about", webService.About)
	r.GET("/contact", webService.Contact)
	r.GET("/blog", webService.Construction)

	r.GET("/:order/", webService.Cards)
	r.GET("/:order/:type/", webService.Cards)
	r.GET("/:order/:type/:id/", webService.Animals)
	r.GET("/:order/:type/:id", webService.Animal)

	if err := r.Run(":80"); err != nil {
		log.Printf("error running server: %v\n", err)
	}
}
