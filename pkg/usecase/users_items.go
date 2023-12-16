package usecase

import "github.com/yamato0211/plesio-server/pkg/domain/repository"

type IUsersItemsUseCase interface {
	BuyItem(userID string, itemID string, count int) error
}

type UsersItemsUseCase struct {
	repo repository.UsersItemsRepository
}

func NewUsersItemsUseCase(repo repository.UsersItemsRepository) IUsersItemsUseCase {
	return &UsersItemsUseCase{
		repo: repo,
	}
}

func (u *UsersItemsUseCase) BuyItem(userID string, itemID string, count int) error {
	err := u.repo.BuyItem(userID, itemID, count)
	if err != nil {
		return err
	}
	return nil
}
