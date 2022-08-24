package products

import (
	"ecommerceapi/db/models"
	"ecommerceapi/lib"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

func get(c *echo.Context, db *bun.DB) error {
	return (*c).Redirect(http.StatusMovedPermanently, "/api/products")
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

func InitRoutes(group *echo.Group, db *bun.DB) {
	group.GET("/products", func(c echo.Context) error {
		return get(&c, db)
	})
	group.POST("/products", func(c echo.Context) error {
		return post(&c, db)
	})
	group.GET("/products/:id", func(c echo.Context) error {
		return idGet(&c, db)
	})
	group.PUT("/products", func(c echo.Context) error {
		return idPut(&c, db)
	})
}
