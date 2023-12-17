package entity

import "time"

type UsersWeapons struct {
	UserID    string    `db:"user_id"`
	WeaponID  string    `db:"weapon_id"`
	Count     int       `db:"count"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type UserWeapons struct {
	Weapon
	Count int `db:"count"`
}
