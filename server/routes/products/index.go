package products

import (
	"ecommerceapi/db/models"
	"ecommerceapi/lib"
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

func post(c *echo.Context, db *bun.DB) error {
	ctx := (*c).Request().Context()

	product := models.Product{}

	if err := (*c).Bind(&product); err != nil {
		return lib.CreateNewResponseError(400, err.Error())
	}

	// Validate the request
	if err := (*c).Validate(product); err != nil {
		return err
	}

	if _, err := db.NewInsert().Model(&product).Exec(ctx); err != nil {
		return lib.CreateNewResponseError(500, err.Error())
	}

	return (*c).JSON(200, product)
}

func InitRoutes(app *echo.Echo, db *bun.DB) {
	(*app).GET("/products", func(c echo.Context) error {
		return get(&c, db)
	})
	(*app).POST("/products", func(c echo.Context) error {
		return post(&c, db)
	})
	(*app).GET("/products/:id", func(c echo.Context) error {
		return idGet(&c, db)
	})
	(*app).PUT("/products/:id", func(c echo.Context) error {
		return idPut(&c, db)
	})
}
