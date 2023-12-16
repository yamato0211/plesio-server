package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/yamato0211/plesio-server/pkg/domain/repository"
	"github.com/yamato0211/plesio-server/pkg/utils/config"
)

type redisRepository struct {
	rdb *redis.Client
}

func NewRedisRepository(cfg *config.RedisConfig) repository.RedisRepository {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisEndpoint,
		Password: "",
		DB:       0,
	})
	return &redisRepository{
		rdb: rdb,
	}
}

func (rr *redisRepository) Publish(ctx context.Context, channel string, payload interface{}) error {
	return rr.rdb.Publish(ctx, channel, payload).Err()
}

func (rr *redisRepository) Subscribe(ctx context.Context, channel string) <-chan *redis.Message {
	return rr.rdb.Subscribe(ctx, channel).Channel()
}

func (rr *redisRepository) Ping(ctx context.Context) error {
	return rr.rdb.Ping(ctx).Err()
}
