package repository

type UsersItemsRepository interface {
	Insert(userID string, itemID string, count int) error
}
