package products

import (
	"ecommerceapi/db/models"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

func IdGET(c *echo.Context, db *bun.DB) error {
	ctx := (*c).Request().Context()
	id := (*c).Param("id")

	product := models.Product{}
	err := db.NewSelect().Model(&models.Product{}).Where("id = ?", id).Scan(ctx, &product)

	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return (*c).JSON(200, product)
}

func IdPUT(c *echo.Context, db *bun.DB) error {
	ctx := (*c).Request().Context()
	id := (*c).Param("id")

	query := db.NewSelect().Model(&models.Product{}).Where("id = ?", id)

	count, err := query.Count(ctx)

	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	if count == 0 {
		return echo.NewHTTPError(404, "Product not found")
	}

	product := make(map[string]interface{})
	err = (*c).Bind(&product)

	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	_, err = db.NewUpdate().Model(&product).TableExpr("products").Where("id = ?", id).Exec(ctx)

	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	finalProduct := models.Product{}

	_, err = query.Exec(ctx, &finalProduct)

	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return (*c).JSON(200, finalProduct)
}
