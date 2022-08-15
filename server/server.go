package server

import (
	"ecommerceapi/server/routes/products"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func initializeRoutes(app *echo.Echo) {
	app.GET("/products", products.GET)
}

func InitializeServer() *echo.Echo {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get the port from the environment variables
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatalf("PORT must be set in .env file")
	}

	// Create an instance of Echo
	server := echo.New()

	// Initialize the routes
	initializeRoutes(server)

	// Start the server
	server.Logger.Fatal(server.Start(":" + port))

	return server
}
