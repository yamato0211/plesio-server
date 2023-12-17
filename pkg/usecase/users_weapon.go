package usecase

import (
	"github.com/yamato0211/plesio-server/pkg/domain/entity"
	"github.com/yamato0211/plesio-server/pkg/domain/repository"
)

type IUsersWeaponsUseCase interface {
	DrawGacha(userID string) (*entity.Weapon, error)
}

type UsersWeaponsUseCase struct {
	repo repository.UsersWeaponsRepository
}

func NewUsersWeaponsUseCase(repo repository.UsersWeaponsRepository) IUsersWeaponsUseCase {
	return &UsersWeaponsUseCase{
		repo: repo,
	}
}

func (u *UsersWeaponsUseCase) DrawGacha(userID string) (*entity.Weapon, error) {
	weapon, err := u.repo.DrawGacha(userID)
	if err != nil {
		return nil, err
	}
	return weapon, nil
}
