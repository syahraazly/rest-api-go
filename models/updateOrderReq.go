package models

import "time"

type UpdateOrderRequest struct {
	CustomerName string     `json:"customerName"`
	OrderedAt    *time.Time `json:"orderedAt"`
	Items        []Item     `json:"items"`
}
