package m2mrelations

import (
	"ecommerceapi/db/models"
)

type OrderToProduct struct {
	OrderID   int64           `bun:"order_id,pk"`
	Order     *models.Order   `bun:"rel:belongs-to,join:order_id=id"`
	ProductID int64           `bun:"product_id,pk"`
	Product   *models.Product `bun:"rel:belongs-to,join:product_id=id"`
}
