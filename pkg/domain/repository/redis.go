package repository

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisRepository interface {
	Publish(ctx context.Context, channel string, payload interface{})
	Subscribe(ctx context.Context, channel string) <-chan *redis.Message
	Ping(ctx context.Context) error
}
