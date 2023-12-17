package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/yamato0211/plesio-server/pkg/domain/entity"
	"github.com/yamato0211/plesio-server/pkg/domain/repository"
)

type weaponRepository struct {
	db *sqlx.DB
}

func NewWeaponRepository(db *sqlx.DB) repository.WeaponRepository {
	return &weaponRepository{
		db: db,
	}
}

func (wr *weaponRepository) SelectAll() ([]*entity.Weapon, error) {
	weapons := []*entity.Weapon{}
	err := wr.db.Select(&weapons, "SELECT * FROM weapons ORDER BY reality DESC and atk DESC")
	if err != nil {
		return nil, err
	}
	return weapons, nil
}
