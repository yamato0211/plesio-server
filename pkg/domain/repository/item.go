package repository

import (
	"github.com/yamato0211/plesio-server/pkg/domain/entity"
)

type ItemRepository interface {
	SelectAll() ([]*entity.Item, error)
}
