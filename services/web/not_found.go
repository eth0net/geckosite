package web

import (
	"github.com/eth0net/geckosite/templates"
	"github.com/gin-gonic/gin"
)

// NotFound returns the 404 not found page.
func (s service) NotFound(c *gin.Context) {
	c.Writer.WriteHeader(404)

	data := struct{ Title, Path, Message string }{
		"Not found",
		c.Request.URL.Path,
		"Error 404: Page not found",
	}

	templates.Parse("message").ExecuteTemplate(c.Writer, "layout", data)
}
