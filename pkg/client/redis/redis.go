package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"go-values-generator/internal/config"
)

func NewClient(cfg *config.Config) (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	if err := client.Ping().Err(); err != nil {
		return nil, err
	}

	return client, nil
}
