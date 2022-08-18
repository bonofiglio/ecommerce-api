package main

import (
	"context"
	"ecommerceapi/db"
	"ecommerceapi/server"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func initializeEnvironment() (port, dbUser, dbPassword, dbHost, dbPort, dbName string) {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port = os.Getenv("PORT")
	dbUser = os.Getenv("POSTGRES_USER")
	dbPassword = os.Getenv("POSTGRES_PASSWORD")
	dbName = os.Getenv("POSTGRES_DB")
	dbHost = os.Getenv("DATABASE_HOST")
	dbPort = os.Getenv("DATABASE_PORT")
	passwordPepper := os.Getenv("PASSWORD_PEPPER")

	if port == "" {
		log.Fatalf("PORT must be set in .env file")
	}

	if dbUser == "" {
		log.Fatalf("POSTGRES_USER must be set in .env file")
	}

	if dbPassword == "" {
		log.Fatalf("POSTGRES_PASSWORD must be set in .env file")
	}

	if dbName == "" {
		log.Fatalf("POSTGRES_DB must be set in .env file")
	}

	if dbHost == "" {
		log.Fatalf("DATABASE_HOST must be set in .env file")
	}

	if dbPort == "" {
		log.Fatalf("DATABASE_PORT must be set in .env file")
	}

	if passwordPepper == "" {
		log.Fatalf("PASSWORD_PEPPER must be set in .env file")
	}

	return port, dbUser, dbPassword, dbHost, dbPort, dbName
}

func main() {
	ctx := context.Background()
	port, dbUser, dbPassword, dbHost, dbPort, dbName := initializeEnvironment()
	postgreDb := db.InitializeDatabase(&dbUser, &dbPassword, &dbHost, &dbPort, &dbName, &ctx)

	server.Init(&port, postgreDb)
}
