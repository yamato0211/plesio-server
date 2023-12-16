package usecase

import (
	"github.com/labstack/echo/v4"
	"github.com/yamato0211/plesio-server/pkg/domain/entity"
	"github.com/yamato0211/plesio-server/pkg/domain/repository"
)

type IUserUsecase interface {
	GetUser(ctx echo.Context, id string) (*entity.User, error)
	CreateUser(ctx echo.Context, name string, email string, git_id string) error
	LoginBonus(ctx echo.Context, id string, git_id string) (*entity.User, error)
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

func (uu *UserUsecase) CreateUser(ctx echo.Context, name string, email string, git_id string) error {
	err := uu.repo.Insert(ctx, name, email, git_id)
	if err != nil {
		return err
	}
	return nil
}

func (uu *UserUsecase) LoginBonus(ctx echo.Context, id string, git_id string) (*entity.User, error) {
	user, err := uu.repo.LoginBonus(ctx, id, git_id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
