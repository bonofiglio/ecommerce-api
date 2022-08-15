package products

import "github.com/labstack/echo/v4"

func GET(c echo.Context) error {
	products := []string{}

	return c.JSON(200, products)
}
