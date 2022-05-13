package web

import (
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"io"
)

// S3 handles minio operations.
func (s service) S3(c *gin.Context) {
	bucket := c.Param("bucket")

	obj, err := s.s3.GetObject(
		c, bucket,
		c.Request.URL.Path[len(bucket)+5:],
		minio.GetObjectOptions{},
	)

	if err != nil {
		s.NotFound(c)
		return
	}

	io.Copy(c.Writer, obj)
}
