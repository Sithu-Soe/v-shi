package handler

import (
	"v-shi/pkg/ds"

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
}
