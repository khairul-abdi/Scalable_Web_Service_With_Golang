package models

import (
	"time"
)

type Order struct {
	OrderID      int       `json:"order_id,omitempty"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Items        []Item    `json:"items"`
}
