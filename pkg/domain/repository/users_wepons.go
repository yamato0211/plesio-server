package repository

import "github.com/yamato0211/plesio-server/pkg/domain/entity"

type UsersWeaponsRepository interface {
	DrawGacha(userID string) (*entity.Weapon, error)
}
