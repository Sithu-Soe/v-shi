package models

import (
	"fmt"
	"v-shi/conf"

	"gorm.io/gorm"
)

type Food struct {
	gorm.Model
	Name        string      `gorm:"column:name;type:varchar(150);not null" json:"name"`
	Description string      `gorm:"column:description" json:"description"`
	Categories  []*Category `gorm:"many2many:categories_foods"`
	ShopID      uint64      `gorm:"column:shop_id;not null" json:"shop_id"`
	Shop        Shop
}

func (m *Food) TableName() string {
	return "foods"
}

type FoodImage struct {
	gorm.Model
	Filename string `gorm:"column:filename;type:varchar(150);not null" json:"filename"`
	FoodID   uint64
	Food     Food
	URL      string `gorm:"-" json:"url"`
}

func (m *FoodImage) TableName() string {
	return "food_images"
}

func (s *FoodImage) AfterFind(tx *gorm.DB) (err error) {
	s.URL = fmt.Sprintf("http://%v/food/images/%v", conf.FHostName, s.Filename)
	return
}
