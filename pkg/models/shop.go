package models

import (
	"database/sql"
	"fmt"
	"v-shi/conf"

	"gorm.io/gorm"
)

type Shop struct {
	gorm.Model
	Name         string `gorm:"column:name;type:varchar(150);not null" json:"name"`
	LogoFilename string `gorm:"column:logo_filename;type:varchar(255)" json:"logo_filename"`
	ShopOwnerID  uint64
	ShopOwner    ShopOwner
	URL          string `gorm:"-" json:"url"`
}

func (s *Shop) AfterFind(tx *gorm.DB) (err error) {
	s.URL = fmt.Sprintf("http://%v/shop/images/%v", conf.FHostName, s.LogoFilename)
	return
}

type ShopOwner struct {
	gorm.Model
	Name     string `gorm:"column:name;type:varchar(150);not null" json:"name"`
	Email    string `gorm:"column:email;type:varchar(255);not null;unique" json:"email"`
	Password string `gorm:"column:password;type:varchar(255)" json:"password"`
	Shops    []*Shop
}

type ShopLocation struct {
	// Lat    float64 `gorm:"type:decimal(10,8)"`
	// Lng    float64 `gorm:"type:decimal(11,8)"`
	gorm.Model
	Name        string         `gorm:"column:name;type:varchar(100);not null" json:"name"`
	Description sql.NullString `gorm:"column:description" json:"description"`
	Lat         float64        `gorm:"column:lat;not null"`
	Lng         float64        `gorm:"column:lng;not null"`
	ShopID      uint64         `gorm:"column:shop_id;not null" json:"shop_id"`
	Shop        Shop
}
