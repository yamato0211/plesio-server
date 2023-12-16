package entity

import "time"

type Item struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Type      string    `db:"type"`
	Amount    int       `db:"amount"`
	Reality   int       `db:"reality"`
	Price     int       `db:"price"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
