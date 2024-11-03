package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

func (r *Repo) SaveTokenInRedis(key, data string) error {
	ctx := context.Background()
	err := r.redis.Set(ctx, key, data, 2*time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) DeleteTokenRedis(key string) error {
	ctx := context.Background()
	err := r.redis.Del(ctx, key).Err()

	if err != nil {
		if err == redis.Nil {
			return nil
		}

		return err
	}

	return nil
}

func (r *Repo) GetRedisData(key string) (string, error) {
	ctx := context.Background()

	data, err := r.redis.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return "", err
	}

	return data, nil
}

func (r *Repo) SaveAccountFreeze(key, data string) error {
	ctx := context.Background()
	err := r.redis.Set(ctx, key, data, 20*time.Minute).Err()
	if err != nil {
		return err
	}

	return nil
}
