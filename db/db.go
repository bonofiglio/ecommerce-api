package db

import (
	"context"
	"database/sql"
	"ecommerceapi/db/models"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func InitializeDatabase(username, password, host, port, name *string, ctx *context.Context) *bun.DB {
	dsn := "postgres://" + *username + ":" + *password + "@" + *host + ":" + *port + "/" + *name + "?sslmode=disable"

	// Connect to the database
	db := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	bunInstance := bun.NewDB(db, pgdialect.New())

	_, err := bunInstance.NewCreateTable().IfNotExists().Model((*models.Product)(nil)).Exec(*ctx)

	if err != nil {
		panic(err)
	}

	return bunInstance
}
