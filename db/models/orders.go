package models

import (
	"time"

	"github.com/uptrace/bun"
)

type OrderStatus string

const (
	Pending  OrderStatus = "pending"
	Complete OrderStatus = "complete"
	Canceled OrderStatus = "canceled"
)

type Order struct {
	bun.BaseModel `bun:"table:orders"`

	ID        int64       `bun:"id,pk,autoincrement" json:"id"`
	UserID    int64       `bun:"user_id,notnull" json:"user_id"`
	Products  []*Product  `bun:"m2m:order_to_products,join:Order=Product" json:"products"`
	Total     float64     `bun:"total,notnull" json:"total" validate:"required"`
	Status    OrderStatus `bun:"status,notnull" json:"status"`
	CreatedAt time.Time   `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time   `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updated_at"`
	Active    bool        `bun:"active,notnull" json:"active"`
}
