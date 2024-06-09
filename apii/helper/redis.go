package helper

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var (
	RedisCli *redis.Client
	Ctx      = context.Background()
)

func InitRedis() {
	RedisCli = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := RedisCli.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	log.Println("Connected to Redis")
}
