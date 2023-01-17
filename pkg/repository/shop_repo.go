package repository

import (
	"context"
	"v-shi/pkg/models"

	"gorm.io/gorm"
)

type shopRepository struct {
	db *gorm.DB
}

func newShopRepository(rConfig *RepoConfig) *shopRepository {
	return &shopRepository{
		db: rConfig.DS.DB,
	}
}

func (r *shopRepository) Create(ctx context.Context, shop *models.Shop) error {
	return r.db.WithContext(ctx).Debug().Create(&shop).Error
}
