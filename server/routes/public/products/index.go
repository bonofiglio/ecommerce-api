package products

import (
	"ecommerceapi/db/models"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

func get(c *echo.Context, db *bun.DB) error {
	ctx := (*c).Request().Context()

	products := make([]map[string]interface{}, 0)

	if err := db.NewSelect().Model(&models.Product{}).Scan(ctx, &products); err != nil {
		log.Print(err)
	}

	return (*c).JSON(200, products)
}

func InitRoutes(group *echo.Group, db *bun.DB) {
	group.GET("/products", func(c echo.Context) error {
		return get(&c, db)
	})
	group.GET("/products/:id", func(c echo.Context) error {
		return idGet(&c, db)
	})
}
