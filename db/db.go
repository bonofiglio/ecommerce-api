package db

import (
	"context"
	"database/sql"
	"ecommerceapi/db/models"
	"ecommerceapi/db/models/m2mrelations"
	"ecommerceapi/lib"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func initializeTables(db *bun.DB, ctx *context.Context) {
	lib.PanicOnDbError(db.NewCreateTable().IfNotExists().Model((*models.Product)(nil)).Exec(*ctx))
	lib.PanicOnDbError(db.NewCreateTable().IfNotExists().Model((*models.User)(nil)).Exec(*ctx))
	lib.PanicOnDbError(db.NewCreateTable().IfNotExists().Model((*models.Order)(nil)).Exec(*ctx))
}

func initializeRelations(db *bun.DB) {
	db.RegisterModel((*m2mrelations.OrderToProduct)(nil))
}

func InitializeDatabase(username, password, host, port, name *string, ctx *context.Context) *bun.DB {
	dsn := "postgres://" + *username + ":" + *password + "@" + *host + ":" + *port + "/" + *name + "?sslmode=disable"

	// Connect to the database
	db := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	bunInstance := bun.NewDB(db, pgdialect.New())

	initializeRelations(bunInstance)
	initializeTables(bunInstance, ctx)

	return bunInstance
}
