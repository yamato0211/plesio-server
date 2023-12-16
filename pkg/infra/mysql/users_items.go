package mysql

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/yamato0211/plesio-server/pkg/domain/entity"
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

func (uir *usersItemsRepository) BuyItem(userID string, itemID string, count int) error {
	// トランザクションを開始
	tx, err := uir.db.Beginx()
	if err != nil {
		return err
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

	// ユーザー情報を取得
	var user entity.User
	if err = tx.Get(&user, "SELECT * FROM users WHERE id = ?", userID); err != nil {
		return err
	}

	// アイテム情報を取得
	var item entity.Item
	if err = tx.Get(&item, "SELECT * FROM items WHERE id = ?", itemID); err != nil {
		return err
	}

	// コインが足りているか確認
	totalPrice := item.Price * count
	if user.Coin < totalPrice {
		return fmt.Errorf("insufficient coins")
	}

	// ユーザーのコインを更新
	newCoin := user.Coin - totalPrice
	if _, err = tx.Exec("UPDATE users SET coin = ? WHERE id = ?", newCoin, userID); err != nil {
		return err
	}

	// UsersItems テーブルにレコードを挿入または更新
	if _, err = tx.Exec(`INSERT INTO users_items (user_id, item_id, count) VALUES (?, ?, ?) 
                         ON DUPLICATE KEY UPDATE count = count + VALUES(count)`,
		userID, itemID, count); err != nil {
		return err
	}

	return nil
}
