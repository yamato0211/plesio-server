package repository

import "github.com/yamato0211/plesio-server/pkg/domain/entity"

type WeaponRepository interface {
	SelectAll() ([]*entity.Weapon, error)
}
