package redis

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient() *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	for i := 1; i <= 10; i++ {
		_, err := client.Ping(context.Background()).Result()

		if err != nil && i != 10 {
			log.Println(err)
		} else if err != nil {
			log.Fatal(err)
		}
		time.Sleep(5 * time.Second)
	}

	return &RedisClient{
		Client: client,
	}
}

func (r *RedisClient) Set(key, value string) error {
	return r.Client.Set(context.Background(), key, value, 30*time.Second).Err()
}

func (r *RedisClient) Get(key string) (string, error) {
	return r.Client.Get(context.Background(), key).Result()
}

func (r *RedisClient) GetAllKeys() ([]string, error) {
	return r.Client.Keys(context.Background(), "*").Result()
}
