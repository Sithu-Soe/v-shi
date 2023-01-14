package model

import "gorm.io/gorm"

type Food struct {
	gorm.Model
	Name        string      `gorm:"column:name;type:varchar(150);not null" json:"name"`
	Description string      `gorm:"column:description" json:"description"`
	Categorys   []*Category `gorm:"many2many:category_foods"`
}

func (m *Food) TableName() string {
	return "foods"
}

type FoodImage struct {
	gorm.Model
	Path   string `gorm:"column:path;type:varchar(150);not null" json:"path"`
	URL    string `gorm:"column:url;type:varchar(255);not null" json:"url"`
	FoodID uint64
	Food   Food
}

func (m *FoodImage) TableName() string {
	return "food_images"
}
