package ds

import (
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

type DataSource struct {
	RDB *redis.Client
	MDB *mongo.Client
}

func NewDataSource() *DataSource {
	return &DataSource{
		RDB: LoadRDB(),
		MDB: LoadMDB(),
	}

}
