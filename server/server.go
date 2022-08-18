package server

import (
	"ecommerceapi/server/routes/products"
	"ecommerceapi/server/routes/users"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

func initializeRoutes(app *echo.Echo, db *bun.DB) {
	products.InitRoutes(app, db)
	users.InitRoutes(app, db)
}

func Init(port *string, db *bun.DB) *echo.Echo {
	// Create an instance of Echo
	server := echo.New()

	// Initialize the routes
	initializeRoutes(server, db)

	// Start the server
	server.Logger.Fatal(server.Start(":" + *port))

	return server
}
