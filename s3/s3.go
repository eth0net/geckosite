package s3

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	// Client stores the minio client.
	Client *minio.Client
	once   sync.Once
)

// Init creates a minio client and connects it to the server.
func Init() *minio.Client {
	once.Do(func() {
		var (
			err       error
			buckets   = []string{"geckos"}
			endpoint  = os.Getenv("MINIO_ENDPOINT")
			accessKey = os.Getenv("MINIO_ACCESS_KEY")
			secretKey = os.Getenv("MINIO_SECRET_KEY")
		)

		Client, err = minio.New(endpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
			Region: "eu-west-1",
			Secure: false,
		})

		if err != nil {
			log.Printf("failed to connect to minio: %s\n", err)
			return
		}

		for _, bucket := range buckets {
			found, err := Client.BucketExists(context.Background(), bucket)
			if err != nil {
				log.Printf("failed to check for bucket %s: %s\n", bucket, err)
				continue
			}
			if !found {
				err = Client.MakeBucket(
					context.Background(), bucket,
					minio.MakeBucketOptions{Region: "eu-west-1"},
				)

				if err != nil {
					log.Printf("failed to make bucket %s: %s\n", bucket, err)
				}
			}
		}
	})

	return Client
}
