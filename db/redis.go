package db

import (
	"log"

	"github.com/go-redis/redis"
)

type Redis struct {
	Client *redis.Client
}

func NewRedisClient() *Redis {
	var newClient Redis
	newClient.Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // replace with your Redis server address
		Password: "",               // replace with your Redis server password
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
