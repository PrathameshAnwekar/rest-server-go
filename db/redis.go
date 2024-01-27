package db

import (
	"fmt"
	"log"

	"github.com/PrathameshAnwekar/rest-server-go/constants"
	"github.com/go-redis/redis"
)

type Redis struct {
	Client *redis.Client
}

func NewRedisClient() *Redis {
	var newClient Redis
	newClient.Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", constants.RedisHost, constants.RedisPort),
		Password: "",
		DB:       0,
	})

	return &newClient
}

func (client *Redis) Close() {
	if client != nil {
		log.Println("Redis connection closed.")
		client.Close()
	}
}
