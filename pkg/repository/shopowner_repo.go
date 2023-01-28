package repository

import (
	"context"
	"v-shi/pkg/models"

	"gorm.io/gorm"
)

type shopOwnerRepository struct {
	db *gorm.DB
}

func newShopOwnerRepository(rConfig *RepoConfig) *shopOwnerRepository {
	return &shopOwnerRepository{
		db: rConfig.DS.DB,
	}
}

func (r *shopOwnerRepository) Create(ctx context.Context, shopOwner *models.ShopOwner) error {
	return r.db.WithContext(ctx).Debug().Create(&shopOwner).Error
}
