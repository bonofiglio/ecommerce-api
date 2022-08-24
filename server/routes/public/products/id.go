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
