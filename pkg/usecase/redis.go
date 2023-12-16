package usecase

import (
	"github.com/labstack/echo/v4"
	"github.com/yamato0211/plesio-server/pkg/domain/repository"
)

type IRedisUsecase interface {
	Ping(ctx echo.Context) error
}

type RedisUsecase struct {
	repo repository.RedisRepository
}

func NewRedisUsecase(rr repository.RedisRepository) IRedisUsecase {
	return &RedisUsecase{
		repo: rr,
	}
}

func (ru *RedisUsecase) Ping(ctx echo.Context) error {
	err := ru.repo.Ping(ctx.Request().Context())
	if err != nil {
		return err
	}
	return nil
}
