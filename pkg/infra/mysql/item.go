package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/yamato0211/plesio-server/pkg/domain/entity"
	"github.com/yamato0211/plesio-server/pkg/domain/repository"
)

type itemRepository struct {
	db *sqlx.DB
}

func NewItemRepository(db *sqlx.DB) repository.ItemRepository {
	return &itemRepository{
		db: db,
	}
}

func (ir *itemRepository) SelectAll() ([]*entity.Item, error) {
	items := []*entity.Item{}
	err := ir.db.Select(&items, "SELECT * FROM items")
	if err != nil {
		return nil, err
	}
	return items, nil
}
