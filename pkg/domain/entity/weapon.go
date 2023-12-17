package entity

import "time"

type Weapon struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Type        string    `db:"type"`
	Description string    `db:"description"`
	Reality     int       `db:"reality"`
	Atk         int       `db:"atk"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
