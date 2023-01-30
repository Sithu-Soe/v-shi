package miio

import (
	"context"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	ShopImagePath = "shop/images/"
	FoodImagePath = "food/images/"
)

var (
	MinioClient *minio.Client
	BucketName  string
	MaxImgSize  = int64(1024 * 1024 * 5) // 5MB
)

func init() {
	BucketName = os.Getenv("MINIO_BUCKET")
	if BucketName == "" {
		log.Panic("there's no bucket")
	}
	// endpoint := "play.min.io"
	// accessKeyID := "Q3AM3UQ867SPQQA43P2F"
	// secretAccessKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY")
	secretAccessKey := os.Getenv("MINIO_SECRET_KEY")
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	MinioClient = minioClient

	exists, err := minioClient.BucketExists(context.Background(), BucketName)

	if err != nil {
		log.Panic(err)
	}

	if !exists {
		if err := minioClient.MakeBucket(context.TODO(), BucketName, minio.MakeBucketOptions{}); err != nil {
			log.Panic(err)
		}
	}
}
