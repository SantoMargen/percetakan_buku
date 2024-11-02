package redis

import "github.com/go-redis/redis/v8"

type Repo struct {
	redis *redis.Client
}

func New(redis *redis.Client) (*Repo, error) {
	return &Repo{
		redis: redis,
	}, nil
}
