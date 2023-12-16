package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/yamato0211/plesio-server/pkg/domain/repository"
)

type usersItemsRepository struct {
	db *sqlx.DB
}

func NewUsersItemsRepository(db *sqlx.DB) repository.UsersItemsRepository {
	return &usersItemsRepository{
		db: db,
	}
}

func (uir *usersItemsRepository) Insert(userID string, itemID string, count int) error {
	sql := `
    INSERT INTO users_items (user_id, item_id, count) 
    VALUES (:user_id, :item_id, :count) 
    ON DUPLICATE KEY UPDATE count = count + VALUES(count);`

	params := struct {
		UserID string `db:"user_id"`
		ItemID string `db:"item_id"`
		Count  int    `db:"count"`
	}{
		UserID: userID,
		ItemID: itemID,
		Count:  count,
	}

	_, err := uir.db.NamedExec(sql, params)
	if err != nil {
		return err
	}
	return nil
}
