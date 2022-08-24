package products

import (
	"ecommerceapi/db/models"
	"ecommerceapi/lib"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

func idGet(c *echo.Context, db *bun.DB) error {
	ctx := (*c).Request().Context()
	id := (*c).Param("id")

	product := models.Product{}

	if err := db.NewSelect().Model(&models.Product{}).Where("id = ?", id).Scan(ctx, &product); err != nil {
		return lib.CreateNewResponseError(500, err.Error())
	}

	return (*c).JSON(200, product)
}

func idPut(c *echo.Context, db *bun.DB) error {
	ctx := (*c).Request().Context()
	id := (*c).Param("id")

	query := db.NewSelect().Model(&models.Product{}).Where("id = ?", id)

	count, err := query.Count(ctx)

	if err != nil {
		return lib.CreateNewResponseError(500, err.Error())
	}

	if count == 0 {
		return lib.CreateNewResponseError(404, "Product not found")
	}

	product := make(map[string]interface{})

	if err := (*c).Bind(&product); err != nil {
		return lib.CreateNewResponseError(400, err.Error())
	}

	if _, err := db.NewUpdate().Model(&product).TableExpr("products").Where("id = ?", id).Exec(ctx); err != nil {
		return lib.CreateNewResponseError(500, err.Error())
	}

	finalProduct := models.Product{}

	if _, err := query.Exec(ctx, &finalProduct); err != nil {
		return lib.CreateNewResponseError(500, err.Error())
	}

	return (*c).JSON(200, finalProduct)
}
