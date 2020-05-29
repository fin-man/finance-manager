package models

import (
	"github.com/fin-man/finance-manager/redis"
)

type RedisModel struct {
	RedisClient *redis.RedisClient
}

func NewRedisModel() *RedisModel {

	redisClient := redis.NewRedisClient()

	return &RedisModel{
		RedisClient: redisClient,
	}
}
