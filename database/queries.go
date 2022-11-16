package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/ikiselewskii/avito-test-task/models"
)

func CreateTables() {
	err := createUsersTable()
	if err != nil {
		panic(err)
	}
	err = createReservationsTable()
	if err != nil {
		panic(err)
	}
}

func createUsersTable() error {
	_, err := DB.NewCreateTable().
		Model((*models.Customer)(nil)).
		IfNotExists().
		Exec(context.Background())
	if err != nil {
		log.Println("Failed to create Users Table ", err)
	}
	return err
}

func createReservationsTable() error {
	_, err := DB.NewCreateTable().
		Model((*models.Transaction)(nil)).
		IfNotExists().
		Exec(context.Background())
	if err != nil {
		log.Println("Failed to create ReservationsTable ", err)
	}
	return err
}

func AddMoney(to models.Customer, ctx context.Context) error {
	_, err := DB.NewInsert().
		Model(&to).
		On("CONFLICT (id) DO UPDATE").
		Set("balance = customer.balance + EXCLUDED.balance").
		Exec(ctx)
	if err != nil {
		log.Println(err)
	}
	return err
}

func Reserve(tr models.Transaction, ctx context.Context) error {
	tx, err := DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.Println("Failed to begin transaction ", err)
		tx.Rollback()
		return err
	}

	count, err := tx.NewSelect().
		Model((*models.Customer)(nil)).
		Where("id = ?", tr.FromID).
		Where("balance >= ?", tr.Amount).
		Count(ctx)
	if count == 0 {
		log.Println("Not enough money on balance ", sql.ErrNoRows)
		tx.Rollback()
		return sql.ErrNoRows
	}
	if err != nil {
		log.Println("Something went wrong ", err)
		return err
	}
	_, err = tx.NewUpdate().
		Model((*models.Customer)(nil)).
		Set("balance = balance - ?", tr.Amount).
		Where("id = ?", tr.FromID).
		Exec(ctx)
	if err != nil {
		log.Println("Failed to reduce balance ", err)
		tx.Rollback()
		return err
	}
	tr.Status = "reserved"
	_, err = tx.NewInsert().
		Model(&tr).
		Exec(ctx)
	if err != nil {
		log.Println("Can not reserve ", err)
		tx.Rollback()
		return err
	}
	return err
}
