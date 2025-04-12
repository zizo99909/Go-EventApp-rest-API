package caching

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func ConnectRedis() {
	ctx := context.Background()

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "redis-18116.c246.us-east-1-4.ec2.redns.redis-cloud.com:18116",
		Username: "default",
		Password: "QntaHqgZMvYzhW55O0f7kWE96K2evXu1",
		DB:       0,
	})

	_, err := RedisClient.Ping(ctx).Result()

	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	} else {
		fmt.Println("Connected to Redis successfully!")
	}

}
