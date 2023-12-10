package usecase

import (
	"github.com/labstack/echo/v4"
	"github.com/yamato0211/plesio-server/pkg/domain/entity"
	"github.com/yamato0211/plesio-server/pkg/domain/repository"
)

type IUserUsecase interface {
	GetUser(ctx echo.Context, id string) (*entity.User, error)
}

type UserUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(us repository.UserRepository) IUserUsecase {
	return &UserUsecase{
		repo: us,
	}
}

func (uu *UserUsecase) GetUser(ctx echo.Context, id string) (*entity.User, error) {
	user, err := uu.repo.Select(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
