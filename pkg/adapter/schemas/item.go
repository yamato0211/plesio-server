package schemas

import "time"

type Item struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Amount    int       `json:"amount"`
	Reality   int       `json:"reality"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BuyItemRequest struct {
	ItemID string `json:"item_id"`
	Count  int    `json:"count"`
}
