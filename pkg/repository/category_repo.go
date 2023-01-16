package repository

import (
	"context"
	"fmt"
	"v-shi/cmd/back/graph/model"
	"v-shi/pkg/models"

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
	db := r.db.WithContext(ctx).Debug().Model(&model.Category{})
	db.Where(updateFields.Field, updateFields.Value)
	return db.Updates(updateFields.Data).Error
}

func (r *categoryRepository) All(ctx context.Context) ([]*model.Category, int64, error) {
	list := make([]*model.Category, 0)
	tb := r.db.WithContext(ctx).Debug().Model(&models.Category{})
	var total int64
	tb.Count(&total)
	if err := tb.Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
