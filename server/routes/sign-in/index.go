package signin

import (
	"ecommerceapi/db/models"
	"ecommerceapi/lib"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

var invalidCredentialsError = lib.CreateNewResponseError(401, "Invalid credentials")

type LoginFormData struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func post(c *echo.Context, db *bun.DB) error {
	ctx := (*c).Request().Context()

	// Get the data from the login form post request
	loginFormData := LoginFormData{}
	if err := (*c).Bind(&loginFormData); err != nil {
		return lib.CreateNewResponseError(400, err.Error())
	}

	query := db.NewSelect().Model(&models.User{}).Where("email = ?", loginFormData.Email)

	// Check that the user exists in the database
	count, err := query.Count(ctx)

	if err != nil {
		return lib.CreateNewResponseError(500, err.Error())
	}

	if count == 0 {
		return invalidCredentialsError
	}

	// Get the user object from the database
	user := models.User{}
	if _, err := query.Exec(ctx, &user); err != nil {
		return lib.CreateNewResponseError(500, err.Error())
	}

	// Compare the hashes to see if the entered password is correct
	hashedPassword := lib.HashPassword(loginFormData.Password, user.Salt)
	if hashedPassword != user.Password {
		return invalidCredentialsError
	}

	// Set the token cookie
	expirationDate := lib.CreateExpirationDate()
	tokenString, err := lib.CreateTokenString(user.ID, expirationDate)

	if err != nil {
		return lib.CreateNewResponseError(500, err.Error())
	}

	tokenCookie := lib.MapTokenToCookie(tokenString, expirationDate)
	(*c).SetCookie(tokenCookie)

	// Extract the fields of the safe user from the user struct before returning it
	safeUser := user.SafeUser

	return (*c).JSON(200, safeUser)
}

func InitRoutes(app *echo.Echo, db *bun.DB) {
	(*app).POST("/sign-in", func(c echo.Context) error {
		return post(&c, db)
	})
}
