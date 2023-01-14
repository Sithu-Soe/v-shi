package ds

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type DataSource struct {
	DB  *gorm.DB
	RDB *redis.Client
	// MDB *mongo.Database
}

func NewDataSource() *DataSource {
	return &DataSource{
		DB:  LoadDB(),
		RDB: LoadRDB(),
		// MDB: LoadMDB(),
	}

}
