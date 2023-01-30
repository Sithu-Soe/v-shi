package repository

import (
	"context"
	"fmt"
	"log"
	"v-shi/cmd/back/graph/model"
	"v-shi/pkg/models"
	"v-shi/pkg/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type foodRepository struct {
	db *gorm.DB
}

func newFoodRepository(rConfig *RepoConfig) *foodRepository {
	return &foodRepository{
		db: rConfig.DS.DB,
	}
}

func (r *foodRepository) Create(ctx context.Context, food *models.Food) error {
	return r.db.WithContext(ctx).Debug().Create(&food).Error
}

func (r *foodRepository) CreateWithCategories(ctx context.Context, food *models.Food, categories []*models.Category) error {
	tx := r.db.WithContext(ctx).Debug().Begin()
	defer tx.Commit()

	tb := tx.Model(&models.Food{})
	if err := tb.Create(food).Error; err != nil {
		tx.Rollback()
		return err
	}

	if len(categories) > 0 {
		if err := tb.Model(food).FirstOrInit(food).Association("Categories").Append(categories); err != nil {
			tx.Rollback()
			log.Println(err)
			return err
		}

	}

	return nil
}

func (r *foodRepository) FindAll(ctx context.Context, input *model.FilterFood) ([]*models.Food, int64, error) {
	tb := r.db.WithContext(ctx).Debug().Model(&models.Food{})

	if input != nil {
		r.filterToQuery(input, tb)
	}

	var total int64
	foods := make([]*models.Food, 0)
	// if err := tb.Count(&total).Scopes(utils.Paginate(input.Page, input.PageSize)).Preload("Shop").Preload("Categories").Preload("FoodImages").Find(&foods).Error; err != nil {
	if err := tb.Count(&total).Scopes(utils.Paginate(input.Page, input.PageSize)).Preload(clause.Associations).Find(&foods).Error; err != nil {
		return nil, 0, err
	}

	return foods, total, nil
}

func (r *foodRepository) Update(ctx context.Context, updateFields *models.UpdateFields, categories []*models.Category) error {
	tx := r.db.WithContext(ctx).Debug().Begin()
	defer tx.Commit()

	tb := tx.Model(&models.Food{})
	if err := tb.Where(updateFields.Field, updateFields.Value).Updates(updateFields.Data).Error; err != nil {
		tx.Rollback()
		return err
	}

	if len(categories) > 0 {
		var food models.Food
		if err := tb.First(&food).Error; err != nil {
			tx.Rollback()
			return err
		}
		tb := tx.Model(&models.Food{})
		if err := tb.Model(&food).Association("Categories").Replace(categories); err != nil {
			tx.Rollback()
			return err
		}
	}
	return nil
}

func (r *foodRepository) DeleteMany(ctx context.Context, ids string) error {
	return r.db.WithContext(ctx).Debug().Delete(&models.Food{}, fmt.Sprintf("id in (%s)", ids)).Error
}

func (r *foodRepository) DeleteFoodWithCategories(ctx context.Context, ids string) error {
	tx := r.db.WithContext(ctx).Debug().Begin()
	defer tx.Commit()

	foods := make([]*models.Food, 0)
	tb := tx.Model(models.Food{})
	if err := tb.Find(&foods, fmt.Sprintf("id in (%s)", ids)).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, food := range foods {
		tb = tx.Model(models.Food{})
		if err := tb.Model(food).Association("Categories").Clear(); err != nil {
			log.Println(err)
			tx.Rollback()
			return err
		}
	}

	tb = tx.Model(models.Food{})
	if err := tb.Delete(&models.Food{}, fmt.Sprintf("id in (%s)", ids)).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *foodRepository) CreateFoodImages(ctx context.Context, foodImages []*models.FoodImage) error {
	return r.db.WithContext(ctx).Debug().Create(foodImages).Error
}

// utilities start here
func (r *foodRepository) filterToQuery(input *model.FilterFood, tb *gorm.DB) {
	tb.Table("foods as f")
	tb.Joins("LEFT JOIN shops s ON f.shop_id = s.id")
	tb.Joins("LEFT JOIN categories_foods cf ON f.id = cf.food_id")
	tb.Joins("LEFT JOIN categories c ON c.id = cf.category_id")
	tb.Group("f.id")
	if input.ID != nil {
		tb.Where("f.id", input.ID)
	}

	if input.Name != nil {
		tb.Where("f.name LIKE ?", "%"+*input.Name+"%")
	}

	if input.Description != nil {
		tb.Where("f.description LIKE ?", "%"+*input.Description+"%")
	}

	if input.StartTime != nil && input.EndTime != nil {
		tb.Where("s.created_at BETWEEN ? AND ?", input.StartTime, input.EndTime)
	}

	if input.ShopID != nil {
		tb.Where("s.id", input.ShopID)
	}

	if input.ShopName != nil {
		tb.Where("s.name LIKE ?", "%"+*input.ShopName+"%")
	}

	if input.CategoryID != nil {
		tb.Where("c.id", input.CategoryID)
	}

	if input.CategoryName != nil {
		tb.Where("c.name LIKE ?", "%"+*input.CategoryName+"%")
	}
}

// utilities end here
