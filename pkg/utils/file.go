package utils

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
	"v-shi/pkg/miio"

	"github.com/99designs/gqlgen/graphql"
	"github.com/anthonynsimon/bild/transform"
	minio "github.com/minio/minio-go/v7"
)

func CreateGraphQLUploadedFile(file io.ReadSeeker, dst string) error {
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, fileBytes, 0644)

}

func SniffContentTypeFromFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}

	// Reset the read pointer if necessary.
	file.Seek(0, 0)

	// Always returns a valid content-type and "application/octet-stream" if no others seemed to match.
	return http.DetectContentType(buffer), nil
}

func UploadImageWithGQL(ctx context.Context, fullPath string, gFile *graphql.Upload) error {
	fileSize := gFile.Size
	if fileSize > miio.MaxImgSize {
		return fmt.Errorf("image file too large")
	}

	var img image.Image
	var err error
	switch gFile.ContentType {
	case "image/jpeg":
		img, err = jpeg.Decode(gFile.File)
		if err != nil {
			return err
		}
	case "image/png":
		img, err = png.Decode(gFile.File)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported file")
	}

	newImg := transform.Resize(img, 400, 400, transform.Gaussian)

	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, newImg, nil); err != nil {
		return err
	}

	_, err = miio.MinioClient.PutObject(ctx, miio.BucketName, fullPath, buf, int64(buf.Len()), minio.PutObjectOptions{})

	return err

}
