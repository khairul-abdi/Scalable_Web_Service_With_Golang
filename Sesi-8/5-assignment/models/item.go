package models

type Item struct {
	ItemID      int    `json:"item_id,omitempty"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderId     int    `json:"order_id,omitempty"`
}
