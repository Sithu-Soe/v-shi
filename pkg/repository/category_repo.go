package repository

import (
	"context"
	"fmt"
	"v-shi/cmd/back/graph/model"
	"v-shi/pkg/models"
	"v-shi/pkg/utils"

	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func newCategoryRepository(rConfig *RepoConfig) *categoryRepository {
	return &categoryRepository{
		db: rConfig.DS.DB,
	}
}

func (r *categoryRepository) Create(ctx context.Context, category *models.Category) error {
	return r.db.WithContext(ctx).Debug().Create(&category).Error
}

func (r *categoryRepository) FindByField(ctx context.Context, field, value any) (*model.Category, error) {
	db := r.db.WithContext(ctx).Debug().Model(&model.Category{})
	category := model.Category{}
	err := db.First(&category, fmt.Sprintf("BINARY %s = ?", field), value).Error
	return &category, err
}

func (r *categoryRepository) UpdateByFields(ctx context.Context, updateFields *models.UpdateFields) error {
	return r.db.WithContext(ctx).Debug().Model(&models.Category{}).Where(updateFields.Field, updateFields.Value).Updates(updateFields.Data).Error
}

func (r *categoryRepository) DeleteMany(ctx context.Context, ids string) error {
	return r.db.WithContext(ctx).Debug().Delete(&models.Category{}, fmt.Sprintf("id in (%s)", ids)).Error
}

func (r *categoryRepository) FindAll(ctx context.Context, input *model.FilterCategory) ([]*models.Category, int64, error) {
	tb := r.db.WithContext(ctx).Debug().Model(&models.Category{})
	if input != nil {
		r.filterToQuery(input, tb)
	}

	var total int64
	categories := make([]*models.Category, 0)
	if err := tb.Count(&total).Scopes(utils.Paginate(input.Page, input.PageSize)).Find(&categories).Error; err != nil {
		return nil, 0, err
	}

	return categories, total, nil
}

func (r *categoryRepository) Count(ctx context.Context, ids string) (int64, error) {
	var total int64
	return total, r.db.WithContext(ctx).Debug().Model(&models.Category{}).Where(fmt.Sprintf("id IN (%s)", ids)).Count(&total).Error
}

// utilities start
func (r *categoryRepository) filterToQuery(input *model.FilterCategory, tb *gorm.DB) {
	tb.Table("categories as c")
	tb.Joins("LEFT JOIN categories_foods cf ON cf.category_id = c.id")
	tb.Joins("LEFT JOIN foods f ON cf.food_id = f.id")
	tb.Group("c.id")
	if input.ID != nil {
		tb.Where("s.id", input.ID)
	}

	if input.Name != nil {
		tb.Where("s.name LIKE ?", "%"+*input.Name+"%")
	}

	if input.StartTime != nil && input.EndTime != nil {
		tb.Where("s.created_at BETWEEN ? AND ?", input.StartTime, input.EndTime)
	}

	if input.FoodID != nil {
		tb.Where("f.id", input.FoodID)
	}

	if input.FoodName != nil {
		tb.Where("f.name LIKE ?", "%"+*input.FoodName+"%")
	}
}

// utilities end
