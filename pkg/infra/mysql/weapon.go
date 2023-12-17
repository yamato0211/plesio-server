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
	err := wr.db.Select(&weapons, "SELECT * FROM weapons ORDER BY reality ASC, atk ASC")
	if err != nil {
		return nil, err
	}
	return weapons, nil
}

func (wr *weaponRepository) SelectAllByID(userID string) ([]*entity.UserWeapons, error) {
	weapons := []*entity.UserWeapons{}
	query := `SELECT w.*, uw.count
		FROM users_weapons uw
		INNER JOIN weapons w ON uw.weapon_id = w.id
		WHERE uw.user_id = ?`
	err := wr.db.Select(&weapons, query, userID)

	if err != nil {
		return nil, err
	}
	return weapons, nil
}
