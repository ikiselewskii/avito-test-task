package database

import (
	"context"
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
	test := models.Customer{ID: 11, Balance: 1488}
	err = AddMoney(test, context.Background())
	if err != nil {
		panic(err)
	}

}

func createUsersTable()  error {
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
		Model((*models.Reservation)(nil)).
		IfNotExists().
		Exec(context.Background())
	if err != nil {
		log.Println("Failed to create ReservationsTable ", err)
	}
	return err
}

func AddMoney(to models.Customer, ctx context.Context)error{
	_, err := DB.NewInsert().
	Model(&to).
	On("CONFLICT (id) DO UPDATE").
	Set("balance = customer.balance + EXCLUDED.balance").
	Exec(ctx)
	if err != nil{
		log.Println(err)
	}
	return err
}