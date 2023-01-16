package models

import (
	"gorm.io/gorm"
)

type Shop struct {
	gorm.Model
	Name     string `gorm:"column:name;type:varchar(150);not null" json:"name"`
	LogoPath string `gorm:"column:logo_path;type:varchar(255)" json:"logo_path"`
}

type ShopOwner struct {
	gorm.Model
	Name     string `gorm:"column:name;type:varchar(150);not null" json:"name"`
	Email    string `gorm:"column:email;type:varchar(255);not null" json:"email"`
	Password string `gorm:"column:password;type:varchar(255)" json:"password"`
	Shops    []*Shop
}
