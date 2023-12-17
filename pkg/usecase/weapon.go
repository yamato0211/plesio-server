package usecase

import (
	"github.com/yamato0211/plesio-server/pkg/domain/entity"
	"github.com/yamato0211/plesio-server/pkg/domain/repository"
)

type IWeaponUseCase interface {
	GetWeapons() ([]*entity.Weapon, error)
}

type WeaponUseCase struct {
	repo repository.WeaponRepository
}

func NewWeaponUseCase(weaponRepo repository.WeaponRepository) IWeaponUseCase {
	return &WeaponUseCase{
		repo: weaponRepo,
	}
}

func (u *WeaponUseCase) GetWeapons() ([]*entity.Weapon, error) {
	weapons, err := u.repo.SelectAll()
	if err != nil {
		return nil, err
	}

	return weapons, nil
}
