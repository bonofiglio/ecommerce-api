package models

import "github.com/uptrace/bun"

type Product struct {
	bun.BaseModel `bun:"table:products"`

	ID          int64   `bun:"id,pk,autoincrement" json:"id"`
	Slug        string  `bun:"slug,unique,notnull" json:"slug"`
	Sku         string  `bun:"sku,unique,notnull" json:"sku"`
	Name        string  `bun:"name,notnull" json:"name"`
	Description string  `bun:"description" json:"description"`
	Price       float64 `bun:"price,notnull" json:"price"`
	Active      bool    `bun:"active,notnull" json:"active"`
}
