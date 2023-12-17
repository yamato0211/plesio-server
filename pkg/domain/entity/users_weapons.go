package entity

type UsersWeapons struct {
	UserID    string `db:"user_id"`
	WeaponID  string `db:"weapon_id"`
	Count     int    `db:"count"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
