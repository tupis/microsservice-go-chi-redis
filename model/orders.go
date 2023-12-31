package model

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	ID          uint64      `json:"id"`
	CustomerID  uuid.UUID   `json:"customer_id"`
	LineItems   []LineItems `json:"line_items"`
	CreatedAt   *time.Time  `json:"created_at"`
	ShippedAt   *time.Time  `json:"shipped_at"`
	CompletedAt *time.Time  `json:"completed_at"`
}

type LineItems struct {
	ID       uuid.UUID `json:"id"`
	Quantity uint      `json:"quantity"`
	Price    float64   `json:"price"`
}
