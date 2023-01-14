package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Category struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name      string        `bson:"name" json:"name"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at" json:"updated_at"`
	Foods     []*Food       `gorm:"many2many:category_foods"`
}
