package entity

type Weapon struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	Type        string `db:"type"`
	Description string `db:"description"`
	Reality     int    `db:"reality"`
	Atk         int    `db:"atk"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
}
