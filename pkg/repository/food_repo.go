package repository

import (
	"context"
	"fmt"
	"log"
	"v-shi/pkg/models"

	"gorm.io/gorm"
)

type foodRepository struct {
	db *gorm.DB
}

func newFoodRepository(rConfig *RepoConfig) *foodRepository {
	return &foodRepository{
		db: rConfig.DS.DB,
	}
}

//	func (r *foodRepository) Create(ctx context.Context, food *models.Food) error {
//		return r.db.WithContext(ctx).Debug().Create(&food).Error
//	}
func (r *foodRepository) Create(ctx context.Context, food *models.Food, categories []*models.Category) error {
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
