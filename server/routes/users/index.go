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
	err := db.NewSelect().Model(&models.User{}).ExcludeColumn("password", "salt").Scan(ctx, &users)

	if err != nil {
		log.Print(users)
	}

	return (*c).JSON(200, users)
}

func post(c *echo.Context, db *bun.DB) error {
	ctx := (*c).Request().Context()

	user := models.User{}
	err := (*c).Bind(&user)

	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	// Unset forbidden fields from the request body in case they were sent in the request
	user.ID = 0
	user.CreatedAt = time.Time{}
	user.UpdatedAt = time.Time{}
	user.Orders = []*models.Order{}
	user.Active = true

	// Check required fields
	if user.Username == "" {
		return echo.NewHTTPError(400, "username is required")
	}
	// TODO: Do email validation
	if user.Email == "" {
		return echo.NewHTTPError(400, "email is required")
	}
	if user.FirstName == "" {
		return echo.NewHTTPError(400, "first_name is required")
	}
	if user.LastName == "" {
		return echo.NewHTTPError(400, "last_name is required")
	}
	if user.Password == "" {
		return echo.NewHTTPError(400, "password is required")
	}

	// Generate a random salt to hash the password with
	user.Salt = lib.GenerateRandomString(64)

	// Hash the password with the salt
	user.Password = lib.HashPassword(user.Password, user.Salt)

	// Insert the user into the database
	_, err = db.NewInsert().Model(&user).Exec(ctx)

	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	// Extract the fields of the safe user from the user struct before returning it
	safeUser := user.SafeUser

	return (*c).JSON(200, safeUser)
}

func InitRoutes(app *echo.Echo, db *bun.DB) {
	(*app).GET("/users", func(c echo.Context) error {
		return get(&c, db)
	})
	(*app).POST("/users", func(c echo.Context) error {
		return post(&c, db)
	})
}
