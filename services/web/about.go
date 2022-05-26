package web

import (
	"github.com/gin-gonic/gin"

	"github.com/eth0net/geckosite/systems/templates"
)

// About returns the about us page.
func (s service) About(c *gin.Context) {
	pageData := struct{ Title, Path string }{"About Us", c.Request.URL.Path}
	templates.Parse("about").ExecuteTemplate(c.Writer, "layout", pageData)
}
