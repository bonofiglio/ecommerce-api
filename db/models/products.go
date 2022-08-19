package models

import "github.com/uptrace/bun"

type Product struct {
	bun.BaseModel `bun:"table:products"`

	ID          int64    `bun:"id,pk,autoincrement" json:"id"`
	Slug        string   `bun:"slug,unique,notnull" json:"slug" validate:"required"`
	Sku         string   `bun:"sku,unique,notnull" json:"sku" validate:"required"`
	Name        string   `bun:"name,notnull" json:"name" validate:"required"`
	Description string   `bun:"description" json:"description"`
	Price       *float64 `bun:"price,notnull" json:"price" validate:"required"`
	Stock       int      `bun:"stock,notnull" json:"stock"`
	Active      bool     `bun:"active,notnull" json:"active"`
}
