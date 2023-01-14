package handler

import (
	"v-shi/pkg/ds"
	"v-shi/pkg/middleware"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	R *gin.Engine
}
type HConfig struct {
	R  *gin.Engine
	DS *ds.DataSource
}

func NewHandler(c *HConfig) *Handler {
	return &Handler{
		R: c.R,
	}
}

func (h *Handler) Register() {
	h.R.Use(middleware.Cors())

	// GraphQL
	graphqlHandler := newGraphqlHandler(h)
	graphqlHandler.register()

}
