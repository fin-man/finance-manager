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

func (c *CollectorService) GetNewCollector(collector string) (string, error) {
	return c.Redis.RedisClient.Get(collector)
}

func (c *CollectorService) GetAllCollectors() ([]string, error) {
	return c.Redis.RedisClient.GetAllKeys()
}

func (c *CollectorService) RemoveCollector(collector string) (int64, error) {
	return c.Redis.RedisClient.Remove(collector)
}
