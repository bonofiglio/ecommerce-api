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
	err := db.NewSelect().Model(&models.Product{}).Scan(ctx, &products)

	if err != nil {
		log.Print(products)
	}

	return (*c).JSON(200, products)
}

func post(c *echo.Context, db *bun.DB) error {
	ctx := (*c).Request().Context()

	product := models.Product{}
	err := (*c).Bind(&product)

	if err != nil {
		return err
	}

	// Check required fields
	if product.Name == "" {
		return echo.NewHTTPError(400, "name is required")
	}
	if product.Price == 0 {
		return echo.NewHTTPError(400, "price is required")
	}
	if product.Sku == "" {
		return echo.NewHTTPError(400, "sku is required")
	}
	if product.Slug == "" {
		return echo.NewHTTPError(400, "slug is required")
	}

	_, err = db.NewInsert().Model(&product).Exec(ctx)

	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return (*c).JSON(200, product)
}

func InitRoutes(app *echo.Echo, db *bun.DB) {
	(*app).GET("/products", func(c echo.Context) error {
		return get(&c, db)
	})
	(*app).GET("/products", func(c echo.Context) error {
		return post(&c, db)
	})
	(*app).GET("/products/:id", func(c echo.Context) error {
		return idGet(&c, db)
	})
	(*app).PUT("/products", func(c echo.Context) error {
		return idPut(&c, db)
	})
}
