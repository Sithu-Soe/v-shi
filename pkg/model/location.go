package model

import "gorm.io/gorm"

type Location struct {
	// Lat    float64 `gorm:"type:decimal(10,8)"`
	// Lng    float64 `gorm:"type:decimal(11,8)"`
	gorm.Model
	ShopID uint64  `gorm:"column:shop_id;not null" json:"shop_id"`
	Lat    float64 `gorm:"column:lat;type:point;not null"`
	Lng    float64 `gorm:"column:lng;type:point;not null"`
	Shop   Shop
}
