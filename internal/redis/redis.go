package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v9"
)

type RedisCli struct {
	con *redis.Client
}

type NoSQL interface {
	Get(ctx context.Context, key string) (string, error)
	SetNX(ctx context.Context, key string, value any, expiration time.Duration) error
}

func New() *RedisCli {
	con := redis.NewClient(&redis.Options{
		Addr:     "0.0.0.0:6379",
		Password: "",
		DB:       0,
	})
	return &RedisCli{
		con,
	}
}

func (r *RedisCli) Get(ctx context.Context, key string) (string, error) {
	val, err := r.con.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *RedisCli) SetNX(ctx context.Context, key string, value any, expiration time.Duration) error {
	err := r.con.SetNX(ctx, key, value, expiration).Err()
	return err
}
