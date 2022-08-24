package users

import (
	"ecommerceapi/db/models"
	"ecommerceapi/lib"
	"log"
	"time"

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

func post(c *echo.Context, db *bun.DB) error {
	ctx := (*c).Request().Context()

	user := models.User{}

	if err := (*c).Bind(&user); err != nil {
		return lib.CreateNewResponseError(400, err.Error())
	}

	// Unset forbidden fields from the request body in case they were sent in the request
	user.ID = 0
	user.CreatedAt = time.Time{}
	user.UpdatedAt = time.Time{}
	user.Orders = []*models.Order{}
	user.Active = true

	// Validate the request
	if err := (*c).Validate(user); err != nil {
		return err
	}

	// Generate a random salt to hash the password with
	user.Salt = lib.GenerateRandomString(64)

	// Hash the password with the salt
	user.Password = lib.HashPassword(user.Password, user.Salt)

	// Insert the user into the database
	if _, err := db.NewInsert().Model(&user).Exec(ctx); err != nil {
		return lib.CreateNewResponseError(500, err.Error())
	}

	// Extract the fields of the safe user from the user struct before returning it
	safeUser := user.SafeUser

	return (*c).JSON(200, safeUser)
}

func InitRoutes(group *echo.Group, db *bun.DB) {
	group.GET("/users", func(c echo.Context) error {
		return get(&c, db)
	})
	group.POST("/users", func(c echo.Context) error {
		return post(&c, db)
	})
}
