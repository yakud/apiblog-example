package redis

import (
	"fmt"

	"github.com/go-redis/redis"
)

func NewConnection(opt *redis.Options) (*redis.Client, error) {
	if opt == nil {
		opt = &redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		}
	}

	client := redis.NewClient(opt)

	pong, err := client.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("redis ping error: %s", err.Error())
	}

	if pong != "PONG" {
		return nil, fmt.Errorf("redis ping error: answer is not PONG")
	}

	return client, err
}
