package repository

import (
	"context"
	"fmt"
	"log"
	"v-shi/cmd/back/graph/model"
	"v-shi/pkg/models"
	"v-shi/pkg/utils"

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

func (r *shopRepository) FindAll(ctx context.Context, input *model.FilterShop) ([]*models.Shop, int64, error) {
	tb := r.db.WithContext(ctx).Debug().Model(&models.Shop{})
	if input != nil {
		r.filterToQuery(input, tb)
	}

	var total int64
	shops := make([]*models.Shop, 0)
	if err := tb.Count(&total).Scopes(utils.Paginate(input.Page, input.PageSize)).Preload("ShopOwner").Find(&shops).Error; err != nil {
		return nil, 0, err
	}

	log.Println(shops[0].ShopOwner, "SHOPS")

	return shops, total, nil
}

func (r *shopRepository) UpdateByFields(ctx context.Context, updateFields *models.UpdateFields) error {
	return r.db.WithContext(ctx).Debug().Model(&models.Shop{}).Where(updateFields.Field, updateFields.Value).Updates(updateFields.Data).Error
}

func (r *shopRepository) DeleteMany(ctx context.Context, ids string) error {
	return r.db.WithContext(ctx).Debug().Delete(&models.Shop{}, fmt.Sprintf("id in (%s)", ids)).Error
}

func (r *shopRepository) filterToQuery(input *model.FilterShop, tb *gorm.DB) {
	tb.Table("shops as s")
	tb.Joins("LEFT JOIN shop_owners so ON so.id = s.shop_owner_id ")
	if input.ID != nil {
		tb.Where("s.id", input.ID)
	}

	if input.Name != nil {
		tb.Where("s.name", input.Name)
	}

	if input.StartTime != nil && input.EndTime != nil {
		tb.Where("s.created_at BETWEEN ? AND ?", input.StartTime, input.EndTime)
	}

	if input.ShopOwnerID != nil {
		tb.Where("so.id", input.ShopOwnerID)
	}

	if input.ShopOwnerName != nil {
		tb.Where("so.name", input.ShopOwnerName)
	}
}
