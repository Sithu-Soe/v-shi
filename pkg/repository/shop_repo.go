package repository

import (
	"context"
	"fmt"
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

	return shops, total, nil
}

func (r *shopRepository) UpdateByFields(ctx context.Context, updateFields *models.UpdateFields) error {
	return r.db.WithContext(ctx).Debug().Model(&models.Shop{}).Where(updateFields.Field, updateFields.Value).Updates(updateFields.Data).Error
}

func (r *shopRepository) DeleteMany(ctx context.Context, ids string) error {
	return r.db.WithContext(ctx).Debug().Delete(&models.Shop{}, fmt.Sprintf("id in (%s)", ids)).Error
}

// Shop Location Start
func (r *shopRepository) CreateShopLocation(ctx context.Context, shopLocation *models.ShopLocation) error {
	return r.db.WithContext(ctx).Debug().Create(&shopLocation).Error
}

func (r *shopRepository) UpdateByFieldsShopLocation(ctx context.Context, updateFields *models.UpdateFields) error {
	return r.db.WithContext(ctx).Debug().Model(&models.ShopLocation{}).Where(updateFields.Field, updateFields.Value).Updates(updateFields.Data).Error
}

func (r *shopRepository) DeleteManyShopLocations(ctx context.Context, ids string) error {
	return r.db.WithContext(ctx).Debug().Delete(&models.ShopLocation{}, fmt.Sprintf("id in (%s)", ids)).Error
}

// Shop Location End

func (r *shopRepository) FindAllShopLocations(ctx context.Context, input *model.FilterShopLocation) ([]*models.ShopLocation, int64, error) {
	tb := r.db.WithContext(ctx).Debug().Model(&models.ShopLocation{})
	if input != nil {
		r.filterToQueryShopLocation(input, tb)
	}

	var total int64
	shopLocations := make([]*models.ShopLocation, 0)
	if err := tb.Count(&total).Scopes(utils.Paginate(input.Page, input.PageSize)).Preload("Shop").Find(&shopLocations).Error; err != nil {
		return nil, 0, err
	}

	return shopLocations, total, nil
}

// utilities start
func (r *shopRepository) filterToQuery(input *model.FilterShop, tb *gorm.DB) {
	tb.Table("shops as s")
	tb.Joins("LEFT JOIN shop_owners so ON so.id = s.shop_owner_id ")
	tb.Group("s.id")
	if input.ID != nil {
		tb.Where("s.id", input.ID)
	}

	if input.Name != nil {
		tb.Where("s.name LIKE ?", "%"+*input.Name+"%")
	}

	if input.StartTime != nil && input.EndTime != nil {
		tb.Where("s.created_at BETWEEN ? AND ?", input.StartTime, input.EndTime)
	}

	if input.ShopOwnerID != nil {
		tb.Where("so.id", input.ShopOwnerID)
	}

	if input.ShopOwnerName != nil {
		tb.Where("so.name LIKE ?", "%"+*input.ShopOwnerName+"%")
	}
}

func (r *shopRepository) filterToQueryShopLocation(input *model.FilterShopLocation, tb *gorm.DB) {
	tb.Table("shop_locations as sl")
	tb.Joins("LEFT JOIN shops s ON s.id = sl.shop_id")
	tb.Group("sl.id")
	if input.ID != nil {
		tb.Where("sl.id", input.ID)
	}

	if input.Name != nil {
		tb.Where("sl.name LIKE ?", "%"+*input.Name+"%")
	}

	if input.Description != nil {
		tb.Where("sl.description LIKE ?", "%"+*input.Description+"%")
	}

	if input.StartTime != nil && input.EndTime != nil {
		tb.Where("sl.created_at BETWEEN ? AND ?", input.StartTime, input.EndTime)
	}

	if input.Lat != nil {
		tb.Where("sl.lat", input.Lat)
	}

	if input.Lng != nil {
		tb.Where("sl.lng", input.Lng)
	}

	if input.ShopID != nil {
		tb.Where("s.id", input.ShopID)
	}

	if input.ShopName != nil {
		tb.Where("s.name LIKE ?", "%"+*input.ShopName+"%")
	}
}

// utilities end
