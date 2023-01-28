package repository

import (
	"context"
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
