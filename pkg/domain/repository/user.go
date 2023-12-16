package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/yamato0211/plesio-server/pkg/domain/entity"
)

type UserRepository interface {
	Select(ctx echo.Context, id string) (*entity.User, error)
	Insert(ctx echo.Context, name string, email string, git_id string) error
	LoginBonus(ctx echo.Context, id string) (*entity.User, error)
}
