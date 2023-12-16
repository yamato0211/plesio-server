package redis

import (
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/yamato0211/plesio-server/pkg/domain/repository"
	"github.com/yamato0211/plesio-server/pkg/utils/config"
)

type redisRepository struct {
	rdb *redis.Client
}

func NewRedisConnector(cfg *config.RedisConfig) repository.RedisRepository {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisEndpoint,
		Password: "",
		DB:       0,
	})
	return &redisRepository{
		rdb: rdb,
	}
}

func (rr *redisRepository) Ping(ctx echo.Context) error {
	return rr.rdb.Ping(ctx.Request().Context()).Err()
}
