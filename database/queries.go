package database

import (
	"context"
	"github.com/ikiselewskii/avito-test-task/models"
	"github.com/uptrace/bun"
)

func CreateTables(db *bun.DB) {
	createUserTable(db)
	createReservationsTable(db)
}

func createUserTable(base *bun.DB) {
	_, err := base.NewCreateTable().
		Model((*models.User)(nil)).
		IfNotExists().
		Exec(context.Background())
	if err != nil {
		panic(err)
	}
	return
}

func createReservationsTable(base *bun.DB) {
	_, err := base.NewCreateTable().
		Model((*models.Reservation)(nil)).
		IfNotExists().
		Exec(context.Background())
	if err != nil {
		panic(err)
	}
	return
}
