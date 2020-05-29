package services

import "github.com/fin-man/finance-manager/server/models"

type CollectorService struct {
	Redis *models.RedisModel
}

func NewCollectorService() *CollectorService {
	redis := models.NewRedisModel()
	return &CollectorService{
		Redis: redis,
	}
}

func (c *CollectorService) CreateNewCollector(key, value string) error {
	return c.Redis.RedisClient.Set(key, value)
}

func (c *CollectorService) GetNewCollector(key string) (string, error) {
	return c.Redis.RedisClient.Get(key)
}

func (c *CollectorService) GetAllCollectors() ([]string, error) {
	return c.Redis.RedisClient.GetAllKeys()
}
