package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name  string  `gorm:"column:name;type:varchar(150);unique" json:"name"`
	Foods []*Food `gorm:"many2many:categories_foods"`
}
