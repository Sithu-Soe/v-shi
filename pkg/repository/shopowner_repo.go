package repository

import (
	"context"
	"fmt"
	"v-shi/cmd/back/graph/model"
	"v-shi/pkg/models"
	"v-shi/pkg/utils"

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

func (r *shopOwnerRepository) FindAll(ctx context.Context, input *model.FilterShopOwner) ([]*models.ShopOwner, int64, error) {
	tb := r.db.WithContext(ctx).Debug().Model(&models.ShopOwner{})
	if input != nil {
		r.filterToQuery(input, tb)
	}

	var total int64
	shopOwners := make([]*models.ShopOwner, 0)
	if err := tb.Count(&total).Scopes(utils.Paginate(input.Page, input.PageSize)).Find(&shopOwners).Error; err != nil {
		return nil, 0, err
	}

	return shopOwners, total, nil
}

func (r *shopOwnerRepository) UpdateByFields(ctx context.Context, updateFields *models.UpdateFields) error {
	return r.db.WithContext(ctx).Debug().Model(&models.ShopOwner{}).Where(updateFields.Field, updateFields.Value).Updates(updateFields.Data).Error
}

func (r *shopOwnerRepository) DeleteMany(ctx context.Context, ids string) error {
	return r.db.WithContext(ctx).Debug().Delete(&models.ShopOwner{}, fmt.Sprintf("id in (%s)", ids)).Error
}

func (r *shopOwnerRepository) filterToQuery(input *model.FilterShopOwner, tb *gorm.DB) {
	tb.Table("shop_owners as so")
	tb.Joins("LEFT JOIN shops s ON so.id = s.shop_owner_id ")
	tb.Group("so.id")
	if input.ID != nil {
		tb.Where("so.id", input.ID)
	}

	if input.Email != nil {
		tb.Where("so.email", input.Email)
	}

	if input.Name != nil {
		tb.Where("so.name LIKE ?", "%"+*input.Name+"%")
	}

	if input.ShopID != nil {
		tb.Where("s.id", input.ShopID)
	}

	if input.ShopName != nil {
		tb.Where("s.name LIKE ?", "%"+*input.ShopName+"%")
	}
}
