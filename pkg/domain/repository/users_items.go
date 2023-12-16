package repository

type UsersItemsRepository interface {
	BuyItem(userID string, itemID string, count int) error
}
