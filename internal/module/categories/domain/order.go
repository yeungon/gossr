package domain

type Order struct {
	ID     int `json:"id"`
	ItemID int `json:"item_id"`
	Qty    int `json:"qty"`
}
