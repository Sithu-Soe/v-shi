package handler

import (
	"log"
	"net/http"
	"v-shi/pkg/miio"
	"v-shi/pkg/repository"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

type shopHandler struct {
	R    *gin.Engine
	repo *repository.Repository
}

func newShopHandler(h *Handler) *shopHandler {
	return &shopHandler{
		R:    h.R,
		repo: h.repo,
	}
}

func (ctr *shopHandler) register() {
	group := ctr.R.Group("/api/shops")
	group.GET("/images/:filename", ctr.getImage)

}

func (ctr *shopHandler) getImage(c *gin.Context) {
	filename := c.Param("filename")

	fullPath := "shop/images/" + filename

	// err := miio.MinioClient.FGetObject(c.Request.Context(), miio.BucketName, fullPath, "shop/images", minio.GetObjectOptions{})
	minioObject, err := miio.MinioClient.GetObject(c.Request.Context(), miio.BucketName, fullPath, minio.GetObjectOptions{})
	if err != nil {
		log.Println(err)
		return
	}

	info, err := minioObject.Stat()
	if err != nil {
		log.Println(err)
		return
	}

	// contentTypeArr := strings.Split()

	c.DataFromReader(http.StatusOK, info.Size, "", minioObject, nil)
}
