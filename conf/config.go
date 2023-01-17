package conf

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

var (
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey

	FHostName string
)

func init() {
	//getting working directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error on getting directory : ", err.Error())
	}

	//load the env file
	err = godotenv.Load(dir + "/conf/.env")
	if err != nil {
		log.Fatalf("error on reading env file : %v\n", err)
	}

	// Load rsa [private]
	privateBytes, err := os.ReadFile(os.Getenv("RSA_PRIVATE"))
	if err != nil {
		// logger.Sugar.Error("Error on loading private key: ", err)
		log.Fatalf("error on reading private key : %v\n", err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatalf("error on parsing private key : %v\n", err)
	}
	PrivateKey = privateKey

	// Load rsa [public]
	publicBytes, err := os.ReadFile(os.Getenv("RSA_PUBLIC"))
	if err != nil {
		log.Fatalf("error on loading public key : %v\n", err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatalf("error on parsing public key : %v\n", err)
	}
	PublicKey = publicKey

	path := "storage"
	for _, p := range []string{"shop/images", "shop/videos", "food/images", "food/videos"} {
		fullPath := path + "/" + p
		if _, err := os.Stat(fullPath); err != nil {
			if errors.Is(err, os.ErrNotExist) || strings.Contains(err.Error(), "no such file or directory") {
				err := os.MkdirAll(fullPath, os.ModePerm)
				if err != nil {
					fmt.Println(err)
				}

			}
		}

	}

	FHostName = os.Getenv("FRONT_HOSTNAME")

}
