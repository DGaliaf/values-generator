package redis

import (
	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

type ValueStorage struct {
	client *redis.Client
}

func NewValueStorage(client *redis.Client) *ValueStorage {
	return &ValueStorage{client: client}
}

func (s ValueStorage) Create(value []byte) (string, []byte, error) {
	key := uuid.New()

	if err := s.client.Set(key.String(), value, 0).Err(); err != nil {
		return "", []byte{}, err
	}

	return key.String(), value, nil
}

func (s ValueStorage) Get(id string) ([]byte, error) {
	result, err := s.client.Get(id).Bytes()
	if err != nil {
		return []byte{}, err
	}

	return result, nil
}
