package entity

import "time"

type UsersItems struct {
	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	ItemID    string    `db:"item_id"`
	Count     int       `db:"count"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
