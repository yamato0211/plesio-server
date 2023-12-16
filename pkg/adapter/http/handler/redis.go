package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/yamato0211/plesio-server/pkg/usecase"
)

type RedisHandler struct {
	usecasea usecase.IRedisUsecase
}

func NewRedisHandler(ru usecase.IRedisUsecase) *RedisHandler {
	return &RedisHandler{
		usecasea: ru,
	}
}

func (rh *RedisHandler) Ping() echo.HandlerFunc {
	return func(c echo.Context) error {
		err := rh.usecasea.Ping(c)
		if err != nil {
			return c.JSON(500, err)
		}
		return c.JSON(200, "redis conn ok!!")
	}
}
