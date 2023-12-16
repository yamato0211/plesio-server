package entity

import (
	"time"
)

type User struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	GitID     string    `db:"git_id"`
	Email     string    `db:"email"`
	IsLogined bool      `db:"is_logined"`
	Coin      int       `db:"coin"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
