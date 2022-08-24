package users

import (
	"ecommerceapi/db/models"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

func get(c *echo.Context, db *bun.DB) error {
	ctx := (*c).Request().Context()

	users := []models.SafeUser{}

	if err := db.NewSelect().Model(&models.User{}).ExcludeColumn("password", "salt").Scan(ctx, &users); err != nil {
		log.Print(users)
	}

	return (*c).JSON(200, users)
}

func InitRoutes(app *echo.Echo, db *bun.DB) {
	(*app).GET("/users", func(c echo.Context) error {
		return get(&c, db)
	})
}
