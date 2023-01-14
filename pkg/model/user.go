package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string  `gorm:"column:name;type:varchar(150);not null" json:"name"`
	Email    *string `gorm:"column:email;type:varchar(255)" json:"email"`
	Phone    *string `gorm:"column:phone;type:char(15)" json:"phone"`
	Password string  `gorm:"column:password;type:varchar(255)" json:"password"`
}
