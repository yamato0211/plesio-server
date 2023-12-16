package usecase

import (
	"github.com/yamato0211/plesio-server/pkg/domain/entity"
	"github.com/yamato0211/plesio-server/pkg/domain/repository"
)

type IItemUsecase interface {
	GetAllItem() ([]*entity.Item, error)
}

type ItemUsecase struct {
	repo repository.ItemRepository
}

func NewItemUsecase(ir repository.ItemRepository) IItemUsecase {
	return &ItemUsecase{
		repo: ir,
	}
}

func (iu *ItemUsecase) GetAllItem() ([]*entity.Item, error) {
	items, err := iu.repo.SelectAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}
