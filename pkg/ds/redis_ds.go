package ds

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

func LoadRDB() *redis.Client {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	user := os.Getenv("REDIS_USER")
	pass := os.Getenv("REDIS_PASS")

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Username: user,
		Password: pass,
		DB:       0,
	})

	log.Println("Successfully connected to redis")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		panic(err)
	}

	return rdb
}
