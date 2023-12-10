package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/yamato0211/plesio-server/pkg/domain/entity"
)

type UserRepository interface {
	Select(ctx echo.Context, id string) (*entity.User, error)
}
