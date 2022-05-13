package web

import (
	"github.com/eth0net/geckosite/templates"
	"github.com/gin-gonic/gin"
)

// Construction returns the under construction page.
func (s service) Construction(c *gin.Context) {
	data := struct{ Title, Path, Message string }{
		"Under Construction",
		c.Request.URL.Path,
		"Page is currently under construction",
	}

	templates.Parse("message").ExecuteTemplate(c.Writer, "layout", data)
}
