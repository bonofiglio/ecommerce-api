package server

import (
	"ecommerceapi/server/routes/products"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

func initializeRoutes(app *echo.Echo, db *bun.DB) {
	// products
	app.GET("/products", func(context echo.Context) error {
		return products.IndexGET(&context, db)
	})
	app.POST("/products", func(context echo.Context) error {
		return products.IndexPOST(&context, db)
	})
	app.GET("/products/:id", func(context echo.Context) error {
		return products.IdGET(&context, db)
	})
	app.PUT("/products/:id", func(context echo.Context) error {
		return products.IdPUT(&context, db)
	})
}

func InitializeServer(port *string, db *bun.DB) *echo.Echo {
	// Create an instance of Echo
	server := echo.New()

	// Initialize the routes
	initializeRoutes(server, db)

	// Start the server
	server.Logger.Fatal(server.Start(":" + *port))

	return server
}
