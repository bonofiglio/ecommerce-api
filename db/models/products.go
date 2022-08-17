package models

import "github.com/uptrace/bun"

type Product struct {
	bun.BaseModel `bun:"table:products"`

	ID          int64   `bun:"id,pk,autoincrement"`
	Slug        string  `bun:"slug,unique,notnull"`
	Sku         string  `bun:"sku,unique,notnull"`
	Name        string  `bun:"name,notnull"`
	Description string  `bun:"description"`
	Price       float64 `bun:"price,notnull"`
	Active      bool    `bun:"active,notnull"`
}
