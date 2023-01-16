package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name  string  `bson:"name" json:"name"`
	Foods []*Food `gorm:"many2many:category_foods"`
}
