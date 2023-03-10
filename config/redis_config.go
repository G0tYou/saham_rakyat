package config

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

func ConnectRedis() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0, // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
		return client, err
	}

	return client, err
}
