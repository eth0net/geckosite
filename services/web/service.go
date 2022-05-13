package web

import (
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

// Service provides handler functions for rendering HTML.
type Service interface {
	About(c *gin.Context)
	Animal(c *gin.Context)
	Animals(c *gin.Context)
	Cards(c *gin.Context)
	Construction(c *gin.Context)
	Contact(c *gin.Context)
	Home(c *gin.Context)
	NotFound(c *gin.Context)
	S3(c *gin.Context)
}

type service struct {
	db *gorm.DB
	s3 *minio.Client
}

// NewService returns a new initialised Service.
func NewService(db *gorm.DB, s3 *minio.Client) Service {
	return &service{db: db, s3: s3}
}
