package utils

import (
	"io"
	"net/http"
	"os"
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
