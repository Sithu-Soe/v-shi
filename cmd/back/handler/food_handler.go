package handler

import (
	"v-shi/pkg/repository"

	"github.com/gin-gonic/gin"
)

type foodHandler struct {
	R    *gin.Engine
	repo *repository.Repository
}

func newFoodHandler(h *Handler) *foodHandler {
	return &foodHandler{
		R:    h.R,
		repo: h.repo,
	}
}

func (ctr *foodHandler) register() {

}
