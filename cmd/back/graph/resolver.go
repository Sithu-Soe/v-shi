package graph

import (
	"v-shi/pkg/repository"
	"v-shi/pkg/service"

	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Repo     *repository.Repository
	DB       *gorm.DB
	MediaSvc *service.MediaService
}
