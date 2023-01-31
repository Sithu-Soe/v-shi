package handler

import (
	"v-shi/pkg/ds"
	"v-shi/pkg/middleware"
	"v-shi/pkg/repository"
	"v-shi/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	R    *gin.Engine
	repo *repository.Repository

	// sevice
	mediaSvc *service.MediaService
}
type HConfig struct {
	R  *gin.Engine
	DS *ds.DataSource
}

func NewHandler(c *HConfig) *Handler {
	repo := repository.NewRepository(&repository.RepoConfig{
		DS: c.DS,
	})
	return &Handler{
		R:        c.R,
		repo:     repo,
		mediaSvc: service.NewMediaServer(),
	}
}

func (h *Handler) Register() {
	h.R.Use(middleware.Cors())

	// GraphQL
	graphqlHandler := newGraphqlHandler(h)
	graphqlHandler.register()

	// shop
	shopHandler := newShopHandler(h)
	shopHandler.register()

	// food
	foodHandler := newFoodHandler(h)
	foodHandler.register()
}
