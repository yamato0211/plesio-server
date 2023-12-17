package mysql

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/yamato0211/plesio-server/pkg/domain/entity"
	"github.com/yamato0211/plesio-server/pkg/domain/repository"
)

type usersWeaponsRepository struct {
	db *sqlx.DB
}

func NewUsersWeaponsRepository(db *sqlx.DB) repository.UsersWeaponsRepository {
	return &usersWeaponsRepository{
		db: db,
	}
}

const gachaCost = 50

func (ur *usersWeaponsRepository) DrawGacha(userID string) (*entity.Weapon, error) {
	tx, err := ur.db.Beginx()
	if err != nil {
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	weapons := []*entity.Weapon{}
	err = tx.Select(&weapons, "SELECT * FROM weapons ORDER BY reality ASC, atk ASC")
	if err != nil {
		return nil, err
	}

	var user entity.User
	if err = tx.Get(&user, "SELECT * FROM users WHERE id = ?", userID); err != nil {
		return nil, err
	}

	// コインが足りているか確認
	if user.Coin < gachaCost {
		return nil, fmt.Errorf("insufficient coins")
	}

	newCoin := user.Coin - gachaCost
	if _, err = tx.Exec("UPDATE users SET coin = ? WHERE id = ?", newCoin, userID); err != nil {
		return nil, err
	}

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100)

	var selectedWeapon *entity.Weapon
	for _, weapon := range weapons {
		randomNumber -= 5
		if randomNumber < 0 {
			selectedWeapon = weapon
			break
		}
	}

	if selectedWeapon != nil {
		query := `
		INSERT INTO users_weapons (user_id, weapon_id, count) 
		VALUES (?, ?, 1) 
		ON DUPLICATE KEY UPDATE count = count + 1;`

		_, err := tx.Exec(query, userID, selectedWeapon.ID)
		if err != nil {
			return nil, err
		}
		return selectedWeapon, nil
	}

	return nil, fmt.Errorf("failed to draw a weapon")
}
