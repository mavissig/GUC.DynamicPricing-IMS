package redis

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/mavissig/GUC.DynamicPricing-IMS/internal/ims/repository"
	"log"
)

type Redis struct {
	client *redis.Client
}

func New(cfg *repository.Config) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.PGRedis.Address,
		Password: cfg.PGRedis.Password,
		DB:       0,
	})
	return &Redis{
		client: client,
	}
}

func (r *Redis) Set(key string, data []byte) error {
	ctx := context.Background()
	_, err := r.client.Set(ctx, key, data, 0).Result()
	log.Printf("[IMS][REPOSITORY][REDIS][SET]: key - %v | data - %v\n", key, string(data))
	return err
}

func (r *Redis) GetByKey(key string) ([]byte, error) {
	ctx := context.Background()
	result, err := r.client.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil, errors.New("not found")
	} else if err != nil {
		return nil, err
	}

	return result, nil
}
